import * as pulumi from "@pulumi/pulumi";
import * as jiraprovider from "@pulumi/jiraprovider";

const myRandomResource = new jiraprovider.Random("myRandomResource", {length: 11});
const myRandomComponent = new jiraprovider.RandomComponent("myRandomComponent", {length: 14});
export const output = {
    value: myRandomResource.result,
};
