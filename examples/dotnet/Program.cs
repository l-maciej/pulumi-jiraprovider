using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Jiraprovider = Pulumi.Jiraprovider;

return await Deployment.RunAsync(() => 
{
    var pulumites2T = new Jiraprovider.JiraGroup("pulumites2t", new()
    {
        JiraGroupName = "pulumitest",
    });

});

