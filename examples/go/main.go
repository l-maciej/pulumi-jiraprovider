package main

import (
	"example.com/pulumi-jiraprovider/sdk/go/jiraprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myRandomResource, err := jiraprovider.NewRandom(ctx, "myRandomResource", &jiraprovider.RandomArgs{
			Length: pulumi.Int(11),
		})
		if err != nil {
			return err
		}
		_, err = jiraprovider.NewRandomComponent(ctx, "myRandomComponent", &jiraprovider.RandomComponentArgs{
			Length: pulumi.Int(14),
		})
		if err != nil {
			return err
		}
		ctx.Export("output", pulumi.StringMap{
			"value": myRandomResource.Result,
		})
		return nil
	})
}
