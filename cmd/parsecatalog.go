
package cmd

import (
	"github.com/deepaksinghvi/catalogwf/pkg/service"
	"github.com/spf13/cobra"
)

// parsecatalogCmd represents the parsecatalog command
var parsecatalogCmd = &cobra.Command{
	Use:   "parsecatalog",
	Short: "Parse Catalog",
	Long: `Parse Catalog Command.`,
	Run: func(cmd *cobra.Command, args []string) {
		service.ParseCatalog(args[0])
		//service.ParseCatalog("Catalog_"+ string(rand.Int()))
	},
}

func init() {
	rootCmd.AddCommand(parsecatalogCmd)
}
