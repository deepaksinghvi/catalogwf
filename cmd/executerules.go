package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// executerulesCmd represents the executerules command
var executerulesCmd = &cobra.Command{
	Use:   "executerules",
	Short: "Execute Rules",
	Long: `Execute Rules Command.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("executerules called")
	},
}

func init() {
	rootCmd.AddCommand(executerulesCmd)
}
