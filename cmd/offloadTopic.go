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
	//"bytes"
	"encoding/json"
	"fmt"
	//"log"
	"os"

	configloader "github.com/acjohnson/nsqustodian/cmd/configloader"
	nsqadmin "github.com/acjohnson/nsqustodian/cmd/nsqadmin"
	//"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/session"
	//"github.com/aws/aws-sdk-go/service/s3"
	//"github.com/nsqio/go-nsq"
	"github.com/spf13/cobra"
)

// offloadTopicCmd represents the offloadTopic command
var offloadTopicCmd = &cobra.Command{
	Use:   "offload-topic",
	Short: "Offload messages from an NSQ topic to an S3 bucket",
	Long:  `Offload messages from an NSQ topic to an S3 bucket in JSON format.`,
	Run: func(cmd *cobra.Command, args []string) {
		offloadTopicMain(cmd)
	},
}

type NSQNode struct {
	Hostname         string `json:"hostname"`
	BroadcastAddress string `json:"broadcast_address"`
	TcpPort          int    `json:"tcp_port"`
	Version          string `json:"version"`
}

type NSQNodeResponse struct {
	Nodes []NSQNode `json:"nodes"`
}

func getNSQNodes(nsqadminAddr string, httpHeaders string) (string, error) {
	payload := []byte(`{}`)
	url := fmt.Sprintf("https://%s:443/api/nodes", nsqadminAddr)
	method := "GET"

	r, err := nsqadmin.NsqAdminCall(nsqadminAddr, httpHeaders, payload, url, method)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return r, nil
}

func offloadTopicMain(cmd *cobra.Command) {
	topic, _ := cmd.Flags().GetString("topic")
	s3BucketName, _ := cmd.Flags().GetString("s3-bucket-name")
	s3BucketKey, _ := cmd.Flags().GetString("s3-bucket-key")
	fmt.Printf("topic: %s\ns3-bucket-name: %s\ns3-bucket-key: %s\n", topic, s3BucketName, s3BucketKey)

	// Get the current config
	config := configloader.ConfigMap()
	currentContext := config.GetString("current_context")
	contextCfg := config.Sub("contexts")
	subCfg := contextCfg.Sub(currentContext)
	nsqadminAddr := subCfg.GetString("nsq-admin")
	httpHeaders := subCfg.GetString("http-headers")

	r, err := getNSQNodes(nsqadminAddr, httpHeaders)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var resp NSQNodeResponse
	json.Unmarshal([]byte(r), &resp)

	for _, node := range resp.Nodes {
		fmt.Printf("Hostname: %s, Broadcast Address: %s, TCP Port: %d, Version: %s\n",
			node.Hostname, node.BroadcastAddress, node.TcpPort, node.Version)
	}
}

func init() {
	offloadTopicCmd.Flags().StringP("topic", "t", "", "NSQ topic to offload messages from (required)")
	offloadTopicCmd.Flags().StringP("s3-bucket-name", "b", "", "S3 bucket name to write messages to (required)")
	offloadTopicCmd.Flags().StringP("s3-bucket-key", "k", "", "S3 bucket key (folder) to write messages to (optional)")
	offloadTopicCmd.MarkFlagRequired("topic")
	offloadTopicCmd.MarkFlagRequired("s3-bucket-name")
	topicsCmd.AddCommand(offloadTopicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// offloadTopicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// offloadTopicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
