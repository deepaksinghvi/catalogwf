package cmd

import (
	"github.com/deepaksinghvi/catalogwf/pkg/service"
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
		service.ApprovalFlow(args[0])
	},
}

func init() {
	rootCmd.AddCommand(approvalflowCmd)
}
