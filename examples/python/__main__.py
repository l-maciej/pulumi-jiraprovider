import pulumi
import pulumi_jiraprovider as jiraprovider

pulumites2_t = jiraprovider.JiraGroup("pulumites2t", jira_group_name="pulumitest")
