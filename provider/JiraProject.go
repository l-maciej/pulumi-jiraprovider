package provider

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type JiraProject struct{}

type JiraProjectArgs struct {
	JiraProjectKey        string `pulumi:"JiraProjectKey"`
	JiraProjectDesc       string `pulumi:"JiraProjectDesc"`
	JiraProjectLead       string `pulumi:"JiraProjectLead"`
	JiraProjectName       string `pulumi:"JiraProjectName"`
	JiraProjectAssType    string `pulumi:"JiraProjectAssType"`
	JiraProjectType       string `pulumi:"JiraProjectType"`
	JiraProjectAvatrID    int64  `pulumi:"JiraProjectAvatrID"`
	JiraProjectCategoryID int64  `pulumi:"JiraProjectCategoryID"`
	//JiraProjectIssScheme  int64  `pulumi:"JiraProjectIssScheme"`
	//JiraProjectNotScheme  int64  `pulumi:"JiraProjectNotScheme"`
	//JiraProjectPermScheme int64  `pulumi:"JiraProjectPermScheme"`

}
type JiraProjectState struct {
	JiraProjectArgs
	ReturnCode int    `pulumi:"returnCode"`
	Result     string `pulumi:"result"`
}

type project_request_struct struct {
	Key          string `json:"key"`
	Description  string `json:"description"`
	Lead         string `json:"lead"`
	Name         string `json:"name"`
	AssigneeType string `json:"assigneeType"`
	ProjectType  string `json:"projectTypeKey"`
	AvatarId     int64  `json:"avatarId"`
	CategoryId   int64  `json:"categoryId"`
	//IssueSecurityScheme int64  `json:"issueSecurityScheme"`
	//NotificationScheme  int64  `json:"notificationScheme"`
	//PermissionScheme    int64  `json:"permissionScheme"`
}

// Create - Delete - Annotate- update
// https://github.com/pulumi/pulumi-go-provider/blob/a019455cf196d45f706cace3d9742ece9b90c33a/examples/dna-store/main.go#L189

func (JiraProject) Create(ctx context.Context, name string, input JiraProjectArgs, preview bool) (string, JiraProjectState, error) {
	state := JiraProjectState{JiraProjectArgs: input, ReturnCode: 420}
	if preview {
		return name, state, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	dataOut, _ := json.Marshal(project_request_struct{
		Key:          input.JiraProjectKey,
		Description:  input.JiraProjectDesc,
		Lead:         input.JiraProjectLead,
		Name:         input.JiraProjectName,
		AssigneeType: input.JiraProjectAssType,
		ProjectType:  input.JiraProjectType,
		AvatarId:     input.JiraProjectAvatrID,
		CategoryId:   input.JiraProjectCategoryID,
		//IssueSecurityScheme: input.JiraProjectIssScheme,
		//NotificationScheme:  input.JiraProjectNotScheme,
		//PermissionScheme:    input.JiraProjectPermScheme,
	})
	outStatus := postHandler(dataOut, cfg.Token, cfg.JURL, "/rest/api/2/project")
	state.appendProjectstate(outStatus)
	return name, state, nil
}

func (JiraProject) Delete(ctx context.Context, id string, state JiraProjectState) error {
	cfg := infer.GetConfig[Config](ctx)
	endpoint := "/rest/api/2/project/" + url.QueryEscape(state.JiraProjectKey)
	status := deleteHandler(cfg.Token, cfg.JURL, endpoint)
	return status
}
