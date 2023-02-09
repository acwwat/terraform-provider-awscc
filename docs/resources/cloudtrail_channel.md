---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cloudtrail_channel Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  A channel receives events from a specific source (such as an on-premises storage solution or application, or a partner event data source), and delivers the events to one or more event data stores. You use channels to ingest events into CloudTrail from sources outside AWS.
---

# awscc_cloudtrail_channel (Resource)

A channel receives events from a specific source (such as an on-premises storage solution or application, or a partner event data source), and delivers the events to one or more event data stores. You use channels to ingest events into CloudTrail from sources outside AWS.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `destinations` (Attributes Set) One or more resources to which events arriving through a channel are logged and stored. (see [below for nested schema](#nestedatt--destinations))
- `name` (String) The name of the channel.
- `source` (String) The ARN of an on-premises storage solution or application, or a partner event source.
- `tags` (Attributes List) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `channel_arn` (String) The Amazon Resource Name (ARN) of a channel.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--destinations"></a>
### Nested Schema for `destinations`

Required:

- `location` (String) The ARN of a resource that receives events from a channel.
- `type` (String) The type of destination for events arriving from a channel.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_cloudtrail_channel.example <resource ID>
```