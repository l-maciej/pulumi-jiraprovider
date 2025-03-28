package main

import (
	"sdk/go/jiraprovider"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := jiraprovider.NewJiraGroup(ctx, "pulumites2t", &jiraprovider.JiraGroupArgs{
			CreatedBy: pulumi.String("pulumitest"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
