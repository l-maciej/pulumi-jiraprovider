// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"example.com/pulumi-jiraprovider/sdk/go/jiraprovider/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

func GetJURL(ctx *pulumi.Context) string {
	return config.Get(ctx, "jiraprovider:jURL")
}
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "jiraprovider:token")
}
