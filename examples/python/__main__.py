import pulumi
import pulumi_jiraprovider as jiraprovider

my_random_resource = jiraprovider.Random("myRandomResource", length=11)
my_random_component = jiraprovider.RandomComponent("myRandomComponent", length=14)
pulumi.export("output", {
    "value": my_random_resource.result,
})
