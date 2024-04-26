---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_workspacesweb_identity_provider Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::WorkSpacesWeb::IdentityProvider Resource Type
---

# awscc_workspacesweb_identity_provider (Resource)

Definition of AWS::WorkSpacesWeb::IdentityProvider Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identity_provider_details` (Map of String)
- `identity_provider_name` (String)
- `identity_provider_type` (String)

### Optional

- `portal_arn` (String)

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `identity_provider_arn` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_workspacesweb_identity_provider.example <resource ID>
```