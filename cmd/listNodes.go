/*
Copyright © 2024 Aaron Johnson <acjohnson@pcdomain.com>

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	configloader "github.com/acjohnson/nsqustodian/cmd/configloader"
	nsqadmin "github.com/acjohnson/nsqustodian/cmd/nsqadmin"
	"github.com/spf13/cobra"
)

// listNodesCmd represents the listNodes command
var listNodesCmd = &cobra.Command{
	Use:   "list",
	Short: "List your NSQ nodes",
	Long:  `List the nodes in your NSQ cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		listNodesMain(cmd)
	},
}

func listNodesMain(cmd *cobra.Command) {
	// Get the current config
	config := configloader.ConfigMap()
	currentContext := config.GetString("current_context")
	contextCfg := config.Sub("contexts")
	subCfg := contextCfg.Sub(currentContext)
	nsqadminAddr := subCfg.GetString("nsq-admin")
	httpHeaders := subCfg.GetString("http-headers")

	resp, err := nsqadmin.GetNsqNodes(nsqadminAddr, httpHeaders)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', 0)

	fmt.Fprintln(w, "HOSTNAME\tADDRESS\tTCP PORT\tVERSION")
	for _, node := range resp.Nodes {
		fmt.Fprintf(w, "%s\t%s\t%d\t%s\n", node.Hostname, node.BroadcastAddress, node.TcpPort, node.Version)
	}

	w.Flush()
}

func init() {
	nodesCmd.AddCommand(listNodesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listNodesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listNodesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
