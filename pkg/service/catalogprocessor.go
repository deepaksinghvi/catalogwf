package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/deepaksinghvi/catalogwf/pkg/dto"
	"github.com/deepaksinghvi/catalogwf/pkg/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var min = 1
var max = 5

func ParseCatalog(catalogName string) {
	log.Info(fmt.Sprintf("Parsing started of catalog name :%s", catalogName))
	time.Sleep(time.Duration(rand.Intn(max-min) + min) * time.Second)
	log.Info(fmt.Sprintf("Parsing completed  of catalog name :%s", catalogName))
}

func ExecuteRules(catalogName string) {
	log.Info(fmt.Sprintf("RuleExecution started of catalog name :%s", catalogName))
	time.Sleep(time.Duration(rand.Intn(max-min) + min) * time.Second)
	log.Info(fmt.Sprintf("RuleExecution completed  of catalog name :%s", catalogName))
}

func ApprovalFlow(catalogName string) {
	log.Info(fmt.Sprintf("ApprovalFlow started of catalog name :%s", catalogName))
	time.Sleep(time.Duration(rand.Intn(max-min) + min) * time.Second)
	log.Info(fmt.Sprintf("ApprovalFlow completed of catalog name :%s", catalogName))
}

func InitiateCatalogWorkflow(request dto.WorkFlowRequest) string {
	url := fmt.Sprintf(`%s/api/v1/workflows/argo/submit`,
		request.BaseUrl,
	)
	template, err := json.Marshal(request.WorkFlowTemplate)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	body := bytes.NewBuffer(template)
	client := &http.Client{}
	response, err  := utils.PostRequest(url,body,"application/json", *client)
	if nil != err {
		log.Error("Error submitting workflow")
		return "Error submitting workflow"
	}
	responseData, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	log.Info("Workflow Initiated: " + string(responseData))
	return "Workflow Initiated Successfully"
}