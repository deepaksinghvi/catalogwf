package cmd

import (
	"fmt"
	"github.com/deepaksinghvi/catalogwf/pkg/service"

	"github.com/spf13/cobra"
)

// buildhierarchyCmd represents the buildhierarchy command
var buildhierarchyCmd = &cobra.Command{
	Use:   "buildhierarchy",
	Short: "Build Hierarchy",
	Long: `Build Hierarchy Command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("buildhierarchy called")
		service.BuildHierarchy(args[0])
	},
}

func init() {
	rootCmd.AddCommand(buildhierarchyCmd)
}
