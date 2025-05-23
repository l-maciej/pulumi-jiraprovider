# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .jira_group import *
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_jiraprovider.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_jiraprovider.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "jiraprovider",
  "mod": "index",
  "fqn": "pulumi_jiraprovider",
  "classes": {
   "jiraprovider:index:JiraGroup": "JiraGroup"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "jiraprovider",
  "token": "pulumi:providers:jiraprovider",
  "fqn": "pulumi_jiraprovider",
  "class": "Provider"
 }
]
"""
)
