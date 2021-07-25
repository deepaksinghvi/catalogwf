package dto


type CatalogRequest struct {
	CatalogName      string `json:"catalogname"`
}


type WorkFlowRequest struct {
	BaseUrl	  string `json:"baseurl"`
	WorkFlowTemplate WorkFlowTemplateRequest `json:"workflowtemplate"`
}

type WorkFlowTemplateRequest struct {
	Namespace     string `json:"namespace"`
	ResourceKind  string `json:"resourceKind"`
	ResourceName  string `json:"resourceName"`
	SubmitOptions struct {
		EntryPoint string   `json:"entryPoint"`
		Parameters []string `json:"parameters"`
	} `json:"submitOptions"`
}