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

	configloader "github.com/acjohnson/nsqustodian/cmd/configloader"
	nsqadmin "github.com/acjohnson/nsqustodian/cmd/nsqadmin"
	"github.com/spf13/cobra"
)

// pauseTopicCmd represents the pauseTopic command
var pauseTopicCmd = &cobra.Command{
	Use:   "pause-topic",
	Short: "Pause NSQ topic.",
	Long:  `Pause NSQ topic.`,
	Run: func(cmd *cobra.Command, args []string) {
		pauseTopicMain(cmd)
	},
}

func pauseTopic(nsqadminAddr string, topic string, httpHeaders string) error {
	payload := []byte(`{"action":"pause"}`)
	url := fmt.Sprintf("https://%s:443/api/topics/%s", nsqadminAddr, topic)
	method := "POST"

	r, err := nsqadmin.NsqAdminCall(nsqadminAddr, httpHeaders, payload, url, method)
	fmt.Println("Response body:", string(r))

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// If we got here, the topic was successfully paused
	fmt.Printf("Topic '%s' paused successfully.\n", topic)
	return nil
}

func pauseTopicMain(cmd *cobra.Command) {
	topic, _ := cmd.Flags().GetString("topic")

	// Get the current config
	config := configloader.ConfigMap()
	currentContext := config.GetString("current_context")
	contextCfg := config.Sub("contexts")
	subCfg := contextCfg.Sub(currentContext)
	nsqadminAddr := subCfg.GetString("nsq-admin")
	httpHeaders := subCfg.GetString("http-headers")

	err := pauseTopic(nsqadminAddr, topic, httpHeaders)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	pauseTopicCmd.Flags().StringP("topic", "t", "", "Topic to pause")
	pauseTopicCmd.MarkFlagRequired("topic")
	topicsCmd.AddCommand(pauseTopicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pauseTopicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pauseTopicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
