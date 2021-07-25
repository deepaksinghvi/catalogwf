package cmd

import (
	"fmt"
	_ "github.com/deepaksinghvi/catalogwf/docs"
	"github.com/deepaksinghvi/catalogwf/pkg/api"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// catalogserviceCmd represents the catalogservice command
var catalogserviceCmd = &cobra.Command{
	Use:   "catalogservice",
	Short: "Catalog Service",
	Long: `Catalog Web-Service.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("catalogservice API Service Start-up")
		router := gin.Default()
		initAPI(router)
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		router.Run(fmt.Sprintf(":%d", 8080))
	},
}

func initAPI(router *gin.Engine) {
	apis := router.Group("/api/v1")
	catalog := apis.Group("/catalog")
	{
		catalog.POST("", api.ProcessCatalog)
		catalog.POST("/workflow",api.InitiateCatalogWorkflow)
	}
	health := apis.Group("/health")
	{
		health.GET("", api.HealthCheck)
	}
}



func init() {
	rootCmd.AddCommand(catalogserviceCmd)
}
