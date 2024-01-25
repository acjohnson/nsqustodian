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
	"github.com/spf13/viper"
)

// createContextCmd represents the createContext command
var createContextCmd = &cobra.Command{
	Use:   "create-context",
	Short: "Create a new NSQustodian context",
	Long:  `Create a new named context in the NSQustodian config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		contextName, _ := cmd.Flags().GetString("name")
		nsqLookupds, _ := cmd.Flags().GetString("nsq-lookupds")

		config := viper.GetViper()

		if _, ok := config.Get("contexts").(map[string]interface{}); !ok {
			config.Set("contexts", map[string]interface{}{})
		}

		contexts := config.Get("contexts").(map[string]interface{})

		if _, ok := contexts[contextName]; ok {
			fmt.Printf("Context '%s' already exists.\n", contextName)
			return
		}

		contexts[contextName] = map[string]string{
			"nsq-lookupds": nsqLookupds,
		}

		err := config.SafeWriteConfig()
		if err != nil {
			fmt.Printf("Failed to write config file: %s\n", err)
			return
		}
		fmt.Printf("Context '%s' created successfully.\n", contextName)
	},
}

func init() {
	createContextCmd.Flags().StringP("name", "n", "", "Name of the context to create")
	createContextCmd.MarkFlagRequired("name")
	createContextCmd.Flags().StringP("nsq-lookupds", "l", "", "URIs for nsq-lookup (can be multiple URIs, comma-separated)")
	createContextCmd.MarkFlagRequired("nsq-lookupds")
	configCmd.AddCommand(createContextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createContextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createContextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
