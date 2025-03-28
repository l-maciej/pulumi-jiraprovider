import * as pulumi from "@pulumi/pulumi";
import * as jiraprovider from "@pulumi/jiraprovider";

const pulumites2T = new jiraprovider.JiraGroup("pulumites2t", {JiraGroupName: "pulumitest"});
