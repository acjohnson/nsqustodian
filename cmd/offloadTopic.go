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

	"github.com/spf13/cobra"
)

// offloadTopicCmd represents the offloadTopic command
var offloadTopicCmd = &cobra.Command{
	Use:   "offload-topic",
	Short: "Offload messages from an NSQ topic to an S3 bucket",
	Long:  `Offload messages from an NSQ topic to an S3 bucket in JSON format.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("offloadTopic called")
	},
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
