// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package jiraprovider

import (
	"context"
	"reflect"

	"errors"
	"github.com/l-maciej/pulumi-jiraprovider/sdk/go/jiraprovider/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type JiraGroup struct {
	pulumi.CustomResourceState

	JiraGroupName pulumi.StringOutput `pulumi:"JiraGroupName"`
	Result        pulumi.StringOutput `pulumi:"result"`
}

// NewJiraGroup registers a new resource with the given unique name, arguments, and options.
func NewJiraGroup(ctx *pulumi.Context,
	name string, args *JiraGroupArgs, opts ...pulumi.ResourceOption) (*JiraGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JiraGroupName == nil {
		return nil, errors.New("invalid value for required argument 'JiraGroupName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource JiraGroup
	err := ctx.RegisterResource("jiraprovider:index:JiraGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJiraGroup gets an existing JiraGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJiraGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JiraGroupState, opts ...pulumi.ResourceOption) (*JiraGroup, error) {
	var resource JiraGroup
	err := ctx.ReadResource("jiraprovider:index:JiraGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JiraGroup resources.
type jiraGroupState struct {
}

type JiraGroupState struct {
}

func (JiraGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*jiraGroupState)(nil)).Elem()
}

type jiraGroupArgs struct {
	JiraGroupName string `pulumi:"JiraGroupName"`
}

// The set of arguments for constructing a JiraGroup resource.
type JiraGroupArgs struct {
	JiraGroupName pulumi.StringInput
}

func (JiraGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jiraGroupArgs)(nil)).Elem()
}

type JiraGroupInput interface {
	pulumi.Input

	ToJiraGroupOutput() JiraGroupOutput
	ToJiraGroupOutputWithContext(ctx context.Context) JiraGroupOutput
}

func (*JiraGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**JiraGroup)(nil)).Elem()
}

func (i *JiraGroup) ToJiraGroupOutput() JiraGroupOutput {
	return i.ToJiraGroupOutputWithContext(context.Background())
}

func (i *JiraGroup) ToJiraGroupOutputWithContext(ctx context.Context) JiraGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JiraGroupOutput)
}

type JiraGroupOutput struct{ *pulumi.OutputState }

func (JiraGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JiraGroup)(nil)).Elem()
}

func (o JiraGroupOutput) ToJiraGroupOutput() JiraGroupOutput {
	return o
}

func (o JiraGroupOutput) ToJiraGroupOutputWithContext(ctx context.Context) JiraGroupOutput {
	return o
}

func (o JiraGroupOutput) JiraGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *JiraGroup) pulumi.StringOutput { return v.JiraGroupName }).(pulumi.StringOutput)
}

func (o JiraGroupOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v *JiraGroup) pulumi.StringOutput { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JiraGroupInput)(nil)).Elem(), &JiraGroup{})
	pulumi.RegisterOutputType(JiraGroupOutput{})
}
