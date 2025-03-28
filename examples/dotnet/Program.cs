using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Jiraprovider = Pulumi.Jiraprovider;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new Jiraprovider.Random("myRandomResource", new()
    {
        Length = 11,
    });

    var myRandomComponent = new Jiraprovider.RandomComponent("myRandomComponent", new()
    {
        Length = 14,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myRandomResource.Result },
        },
    };
});

