package api

import (
	"github.com/deepaksinghvi/catalogwf/pkg/dto"
	"github.com/deepaksinghvi/catalogwf/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)


// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}

// ProcessCatalog godoc
// @Summary Process Catalog
// @ID process-catalog
// @Tags Catalog
// @Accept json
// @Produce json
// @Param catalogName body dto.CatalogRequest true "Process Catalog"
// @Success 200 {string} string
// @Router /catalog [post]
func ProcessCatalog(ctx *gin.Context) {
	var catalogRequest dto.CatalogRequest
	err := ctx.BindJSON(&catalogRequest)
	if err != nil {
		ctx.JSON(401, err)
		return
	}
	go processCatalog(catalogRequest.CatalogName)
	ctx.JSON(http.StatusCreated, "Catalog Processing Initiated")
	return
}

func processCatalog(catalogName string) {
	service.ParseCatalog(catalogName)
	service.ExecuteRules(catalogName)
	service.ApprovalFlow(catalogName)
}

// InitiateCatalogWorkflow godoc
// @Summary Initiate Catalog Workflow
// @ID initiate-catalog-workflow
// @Tags Catalog
// @Accept json
// @Produce json
// @Param workflowRequest body dto.WorkFlowRequest true "Initiate Catalog Workflow"
// @Success 200 {string} string
// @Router /catalog/workflow [post]
func InitiateCatalogWorkflow(ctx *gin.Context) {
	var workflowRequest dto.WorkFlowRequest
	err := ctx.BindJSON(&workflowRequest)
	if err != nil {
		ctx.JSON(401, err)
		return
	}
	ctx.JSON(http.StatusCreated, service.InitiateCatalogWorkflow(workflowRequest))
}