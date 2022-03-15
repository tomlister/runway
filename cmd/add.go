package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tomlister/runway/api"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [author/plugin]",
	Short: "Adds a plugin to your plugin list",
	Long:  `Adds a plugin to your plugin list`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		entry := args[0]
		if !strings.Contains(entry, "/") {
			fmt.Printf("Oops: %s is an invalid entry.\nMake sure it matches the author/plugin syntax.\ne.g IntellectualSites/FastAsyncWorldEdit\n", entry)
			return
		}
		split := strings.Split(entry, "/")
		author := split[0]
		plugin := split[1]
		_, found, err := api.GetProject(author, plugin)
		if err != nil {
			fmt.Printf("Oops: An error occurred while fetching the project data.\n")
			return
		}
		if !found {
			fmt.Printf("Oops: Couldn't find the plugin: %s.\nMake sure that the author and plugin name are correct.", entry)
			return
		}
		fmt.Printf("found")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
