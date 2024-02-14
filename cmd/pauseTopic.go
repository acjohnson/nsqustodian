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
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	config_loader "github.com/acjohnson/nsqustodian/cmd/config_loader"
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
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	headerStrings := strings.Split(httpHeaders, ",")

	for _, headerString := range headerStrings {
		// Split the header string into key and value
		keyValue := strings.SplitN(headerString, ":", 2)
		if len(keyValue) == 2 {
			// Set the header in the request header
			req.Header.Set(keyValue[0], keyValue[1])
		}
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request and check the response status code
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	fmt.Println("Response body: ", string(body))

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to pause topic: %s", string(bodyBytes))
	}

	// If we got here, the topic was successfully paused
	fmt.Printf("Topic '%s' paused successfully.\n", topic)
	return nil
}

func pauseTopicMain(cmd *cobra.Command) {
	topic, _ := cmd.Flags().GetString("topic")

	// Get the current config
	config := config_loader.ConfigMap()
	currentContext := config.Get("current_context").(string)
	contextCfg := config.Sub("contexts")
	subCfg := contextCfg.Sub(currentContext)
	nsqadminAddr := subCfg.Get("nsq-admin").(string)
	httpHeaders := subCfg.Get("http-headers").(string)

	err := pauseTopic(nsqadminAddr, topic, httpHeaders)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	pauseTopicCmd.Flags().StringP("topic", "n", "", "Topic to pause")
	topicsCmd.AddCommand(pauseTopicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pauseTopicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pauseTopicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
