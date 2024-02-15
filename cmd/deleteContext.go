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
	"log"

	config_loader "github.com/acjohnson/nsqustodian/cmd/config_loader"
	"github.com/spf13/cobra"
)

// deleteContextCmd represents the deleteContext command
var deleteContextCmd = &cobra.Command{
	Use:   "delete-context",
	Short: "Deletes a named context from the config file.",
	Long:  `Deletes a named context from the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteContextMain(cmd)
	},
}

func deleteContextMain(cmd *cobra.Command) {
	// Get the current config
	config := config_loader.ConfigMap()

	contextName, _ := cmd.Flags().GetString("name")
	//configFile := config.ConfigFileUsed()

	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading configuration: %v", err)
	}

	// Get the contexts map from the configuration
	contexts := config.Get("contexts").(map[string]interface{})

	// Check if the current context is being deleted
	currentContext := config.GetString("current_context")
	if currentContext == contextName {
		for key := range contexts {
			if key != contextName {
				config.Set("current_context", key)
				break
			}
		}
	}

	// Delete the context from the map
	delete(contexts, contextName)

	// Write the updated configuration to the file
	err = config.WriteConfig()
	if err != nil {
		log.Fatalf("Error writing configuration: %v", err)
	}

	fmt.Printf("Context '%s' deleted successfully.\n", contextName)

}

func init() {
	deleteContextCmd.Flags().StringP("name", "n", "", "Name of the context to delete")
	deleteContextCmd.MarkFlagRequired("name")
	configCmd.AddCommand(deleteContextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteContextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteContextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
