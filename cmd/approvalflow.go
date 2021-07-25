package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// approvalflowCmd represents the approvalflow command
var approvalflowCmd = &cobra.Command{
	Use:   "approvalflow",
	Short: "Approval Flow",
	Long: `Approval Flow Command.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("approvalflow called")
	},
}

func init() {
	rootCmd.AddCommand(approvalflowCmd)
}
