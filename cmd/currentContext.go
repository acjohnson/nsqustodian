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

	configloader "github.com/acjohnson/nsqustodian/cmd/configloader"
	"github.com/spf13/cobra"
)

// currentContextCmd represents the currentContext command
var currentContextCmd = &cobra.Command{
	Use:   "current-context",
	Short: "Displays the currently active context in the config file.",
	Long:  `Displays the currently active context in the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentContextMain(cmd)
	},
}

func currentContextMain(cmd *cobra.Command) {
	// Get the current config
	config := configloader.ConfigMap()
	currentContext := config.GetString("current_context")
	fmt.Printf("Current context is: %s\n", currentContext)
}

func init() {
	configCmd.AddCommand(currentContextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// currentContextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// currentContextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
