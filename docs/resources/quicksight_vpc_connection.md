---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_quicksight_vpc_connection Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of the AWS::QuickSight::VPCConnection Resource Type.
---

# awscc_quicksight_vpc_connection (Resource)

Definition of the AWS::QuickSight::VPCConnection Resource Type.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `availability_status` (String)
- `aws_account_id` (String)
- `dns_resolvers` (List of String)
- `name` (String)
- `role_arn` (String)
- `security_group_ids` (List of String)
- `subnet_ids` (List of String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `vpc_connection_id` (String)

### Read-Only

- `arn` (String)
- `created_time` (String)
- `id` (String) Uniquely identifies the resource.
- `last_updated_time` (String)
- `network_interfaces` (Attributes List) (see [below for nested schema](#nestedatt--network_interfaces))
- `status` (String)
- `vpc_id` (String)

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)


<a id="nestedatt--network_interfaces"></a>
### Nested Schema for `network_interfaces`

Read-Only:

- `availability_zone` (String)
- `error_message` (String)
- `network_interface_id` (String)
- `status` (String)
- `subnet_id` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_quicksight_vpc_connection.example <resource ID>
```