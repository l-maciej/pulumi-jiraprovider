package provider

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type jiraProject struct{}

type jiraProjectArgs struct {
	jiraProjectKey        string `pulumi:"jiraProjectKey"`
	jiraProjectDesc       string `pulumi:"jiraProjectDesc"`
	jiraProjectLead       string `pulumi:"jiraProjectLead"`
	jiraProjectName       string `pulumi:"jiraProjectName"`
	jiraProjectAssType    string `pulumi:"jiraProjectAssType"`
	jiraProjectAvatrID    int64  `pulumi:"jiraProjectAvatrID"`
	jiraProjectCategoryID int64  `pulumi:"jiraProjectCategoryID"`
	jiraProjectIssScheme  int64  `pulumi:"jiraProjectIssScheme"`
	jiraProjectNotScheme  int64  `pulumi:"jiraProjectNotScheme"`
	jiraProjectPermScheme int64  `pulumi:"jiraProjectPermScheme"`
}
type jiraProjectState struct {
	jiraProjectArgs
	Result string `pulumi:"result"`
}

type request_struct struct {
	Key                 string
	Description         string
	Lead                string
	Name                string
	AssigneeType        string
	AvatarId            int64
	CategoryId          int64
	IssueSecurityScheme int64
	NotificationScheme  int64
	PermissionScheme    int64
}

// Create - Delete - Annotate- update
// https://github.com/pulumi/pulumi-go-provider/blob/a019455cf196d45f706cace3d9742ece9b90c33a/examples/dna-store/main.go#L189

func (jiraProject) Create(ctx context.Context, name string, input jiraProjectArgs, preview bool) (string, jiraProjectState, error) {
	state := jiraProjectState{jiraProjectArgs: input}
	if preview {
		return name, state, nil
	}
	cfg := infer.GetConfig[config](ctx)
	outStruct := request_struct{
		Key:                 input.jiraProjectKey,
		Description:         input.jiraProjectDesc,
		Lead:                input.jiraProjectLead,
		Name:                input.jiraProjectName,
		AssigneeType:        input.jiraProjectAssType,    // abstracted to global config andd with annotator
		AvatarId:            input.jiraProjectAvatrID,    // abstracted to global config andd with annotator
		CategoryId:          input.jiraProjectCategoryID, // abstracted to global config andd with annotator
		IssueSecurityScheme: input.jiraProjectIssScheme,  // abstracted to global config andd with annotator
		NotificationScheme:  input.jiraProjectNotScheme,  // abstracted to global config andd with annotator
		PermissionScheme:    input.jiraProjectPermScheme, // abstracted to global config andd with annotator
	}
	dataOut, _ := json.Marshal(outStruct)
	postHandler(dataOut, cfg.Token, cfg.JURL, "/rest/api/2/project")
	return name, state, nil
}

func (jiraProject) Delete(ctx context.Context, id string, state jiraProjectState) error {
	cfg := infer.GetConfig[config](ctx)
	endpoint := "/rest/api/2/group?groupname=" + url.QueryEscape(state.jiraProjectKey)
	status := deleteHandler(cfg.Token, cfg.JURL, endpoint)
	return status
}

func (jp *jiraProjectArgs) Annotate(ctx context.Context, a infer.Annotator) {
	cfg := infer.GetConfig[config](ctx)
	a.SetDefault(&jp.jiraProjectAvatrID, cfg.jiraProjecDefultAvatarId)
}
