// Copyright 2016-2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/pulumi/pulumi-go-provider/infer"
)

// Each resource has a controlling struct.
// Resource behavior is determined by implementing methods on the controlling struct.
// The `Create` method is mandatory, but other methods are optional.
// - Check: Remap inputs before they are typed.
// - Diff: Change how instances of a resource are compared.
// - Update: Mutate a resource in place.
// - Read: Get the state of a resource from the backing provider.
// - Delete: Custom logic when the resource is deleted.
// - Annotate: Describe fields and set defaults for a resource.
// - WireDependencies: Control how outputs and secrets flows through values.
type JiraGroup struct{}

type JiraGroupArgs struct {
	//CreatedBy string `pulumi:"createdBy"`
	JiraGroupName string `pulumi:"JiraGroupName"`
	//GroupPermission   string   `pulumi:"groupPermission"`
	//AssignedToProject []string `pulumi:"assignedProjectKey"`
}

type JiraGroupState struct {
	JiraGroupArgs
	//Slug   string
	Result string `pulumi:"result"`
}

type group_request_struct struct { //Format for Jira
	Name string `json:"name"`
}

func (JiraGroup) Create(ctx context.Context, name string, input JiraGroupArgs, preview bool) (string, JiraGroupState, error) {
	state := JiraGroupState{JiraGroupArgs: input}
	if preview {
		return name, state, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	dataOut, _ := json.Marshal(group_request_struct{Name: input.JiraGroupName}) //Defined per requset since they have diffrenet structure
	postHandler(dataOut, cfg.Token, cfg.JURL, "/rest/api/2/group")
	return name, state, nil
}

func (JiraGroup) Delete(ctx context.Context, id string, state JiraGroupState) error {
	cfg := infer.GetConfig[Config](ctx)
	endpoint := "/rest/api/2/group?groupname=" + url.QueryEscape(state.JiraGroupName)
	status := deleteHandler(cfg.Token, cfg.JURL, endpoint)
	return status
}
