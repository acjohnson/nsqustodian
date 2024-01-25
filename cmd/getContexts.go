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

// getContextsCmd represents the getContexts command
var getContextsCmd = &cobra.Command{
	Use:   "get-contexts",
	Short: "Lists all named contexts in the config file.",
	Long: `Lists all named contexts in the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get-contexts called")
	},
}

func init() {
	configCmd.AddCommand(getContextsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getContextsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getContextsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
