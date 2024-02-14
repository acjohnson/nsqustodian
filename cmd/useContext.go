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

	config_loader "github.com/acjohnson/nsqustodian/cmd/config_loader"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// useContextCmd represents the useContext command
var useContextCmd = &cobra.Command{
	Use:   "use-context",
	Short: "Sets the active context for managing NSQ clusters in the config file.",
	Long:  `Sets the active context for managing NSQ clusters in the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("use-context called")
		contextName, _ := cmd.Flags().GetString("name")
		// Check if the context exists
		exists, err := contextExists(contextName)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		if !exists {
			fmt.Fprintln(os.Stderr, "Error:", contextName+" is not a valid context")
			fmt.Println("To create a new context, run 'nsqustodian create-context'")
			os.Exit(1)
		}
		// Set the top-level "context" key in the config file to the name of the context
		viper.Set("current_context", contextName)
		err = viper.WriteConfig()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		fmt.Printf("Switched to context %s\n", contextName)
	},
}

func contextExists(name string) (bool, error) {
	// Get the current config
	config := config_loader.ConfigMap()

	// Check if the "contexts" key exists
	contexts, ok := config.Get("contexts").(map[string]interface{})
	if !ok {
		return false, fmt.Errorf("config file does not contain a 'contexts' key")
	}

	// Check if the named context exists
	_, ok = contexts[name]
	return ok, nil
}

func init() {
	useContextCmd.Flags().StringP("name", "n", "", "Name of the context to use")
	configCmd.AddCommand(useContextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useContextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useContextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
