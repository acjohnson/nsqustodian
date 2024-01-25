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
		nsqAdmin, _ := cmd.Flags().GetString("nsq-admin")
		httpHeaders, _ := cmd.Flags().GetString("http-headers")

		config := viper.GetViper()

		// Read the existing config file
		err := config.ReadInConfig()
		if err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// Config file not found, so create a new one
				config.Set("contexts", map[string]interface{}{
					contextName: map[string]interface{}{
						"nsq-lookupds": nsqLookupds,
						"nsq-admin":    nsqAdmin,
						"http-headers": httpHeaders,
					},
				})
			} else {
				fmt.Printf("Failed to read config file: %s\n", err)
				return
			}
		} else {
			// Merge the new context into the existing config
			contexts := config.Get("contexts").(map[string]interface{})
			if _, ok := contexts[contextName]; !ok {
				contexts[contextName] = map[string]interface{}{
					"nsq-lookupds": nsqLookupds,
					"nsq-admin":    nsqAdmin,
					"http-headers": httpHeaders,
				}
			}
			context := contexts[contextName].(map[string]interface{})
			if nsqLookupds != "" {
				context["nsq-lookupds"] = nsqLookupds
			}
			if nsqAdmin != "" {
				context["nsq-admin"] = nsqAdmin
			}
			if httpHeaders != "" {
				context["http-headers"] = httpHeaders
			}
		}

		// Write the updated config file back to disk
		if err := config.WriteConfig(); err != nil {
			fmt.Printf("%s\n", err)
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// SafeWriteConfig() creates the config
				err = config.SafeWriteConfig()
				if err != nil {
					fmt.Printf("Failed to write config file: %s\n", err)
					return
				}
				// WriteConfig() writes the new merged config...
				err = config.WriteConfig()
				if err != nil {
					fmt.Printf("Failed to write config file: %s\n", err)
					return
				}

				fmt.Printf("Context '%s' created successfully.\n", contextName)
			}
		} else {
			fmt.Printf("Context '%s' created successfully.\n", contextName)

		}
	},
}

func init() {
	createContextCmd.Flags().StringP("name", "n", "", "Name of the context to create")
	createContextCmd.Flags().StringP("nsq-lookupds", "l", "", "URIs for nsq-lookup (can be multiple URIs, comma-separated)")
	createContextCmd.Flags().StringP("nsq-admin", "a", "", "nsq-admin URI")
	createContextCmd.Flags().StringP("http-headers", "e", "", "http headers to add to nsq-admin calls (comma-separated)")
	createContextCmd.MarkFlagRequired("name")
	createContextCmd.MarkFlagRequired("nsq-admin")
	configCmd.AddCommand(createContextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createContextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createContextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
