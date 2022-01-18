// Code generated by generators/resource/main.go; DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_redshift_event_subscription", eventSubscriptionResourceType)
}

// eventSubscriptionResourceType returns the Terraform awscc_redshift_event_subscription resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Redshift::EventSubscription resource type.
func eventSubscriptionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"cust_subscription_id": {
			// Property: CustSubscriptionId
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Amazon Redshift event notification subscription.",
			//   "type": "string"
			// }
			Description: "The name of the Amazon Redshift event notification subscription.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"customer_aws_id": {
			// Property: CustomerAwsId
			// CloudFormation resource type schema:
			// {
			//   "description": "The AWS account associated with the Amazon Redshift event notification subscription.",
			//   "type": "string"
			// }
			Description: "The AWS account associated with the Amazon Redshift event notification subscription.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"enabled": {
			// Property: Enabled
			// CloudFormation resource type schema:
			// {
			//   "description": "A boolean value; set to true to activate the subscription, and set to false to create the subscription but not activate it.",
			//   "type": "boolean"
			// }
			Description: "A boolean value; set to true to activate the subscription, and set to false to create the subscription but not activate it.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"event_categories": {
			// Property: EventCategories
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the Amazon Redshift event categories to be published by the event notification subscription.",
			//   "insertionOrder": false,
			//   "items": {
			//     "enum": [
			//       "configuration",
			//       "management",
			//       "monitoring",
			//       "security",
			//       "pending"
			//     ],
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Specifies the Amazon Redshift event categories to be published by the event notification subscription.",
			Type:        types.SetType{ElemType: types.StringType},
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayForEach(validate.StringInSlice([]string{
					"configuration",
					"management",
					"monitoring",
					"security",
					"pending",
				})),
			},
		},
		"event_categories_list": {
			// Property: EventCategoriesList
			// CloudFormation resource type schema:
			// {
			//   "description": "The list of Amazon Redshift event categories specified in the event notification subscription.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The list of Amazon Redshift event categories specified in the event notification subscription.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"severity": {
			// Property: Severity
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the Amazon Redshift event severity to be published by the event notification subscription.",
			//   "enum": [
			//     "ERROR",
			//     "INFO"
			//   ],
			//   "type": "string"
			// }
			Description: "Specifies the Amazon Redshift event severity to be published by the event notification subscription.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ERROR",
					"INFO",
				}),
			},
		},
		"sns_topic_arn": {
			// Property: SnsTopicArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the Amazon SNS topic used to transmit the event notifications.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the Amazon SNS topic used to transmit the event notifications.",
			Type:        types.StringType,
			Optional:    true,
		},
		"source_ids": {
			// Property: SourceIds
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of one or more identifiers of Amazon Redshift source objects.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "A list of one or more identifiers of Amazon Redshift source objects.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"source_ids_list": {
			// Property: SourceIdsList
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of the sources that publish events to the Amazon Redshift event notification subscription.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "A list of the sources that publish events to the Amazon Redshift event notification subscription.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				tfsdk.UseStateForUnknown(),
			},
		},
		"source_type": {
			// Property: SourceType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of source that will be generating the events.",
			//   "enum": [
			//     "cluster",
			//     "cluster-parameter-group",
			//     "cluster-security-group",
			//     "cluster-snapshot",
			//     "scheduled-action"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of source that will be generating the events.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"cluster",
					"cluster-parameter-group",
					"cluster-security-group",
					"cluster-snapshot",
					"scheduled-action",
				}),
			},
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of the Amazon Redshift event notification subscription.",
			//   "enum": [
			//     "active",
			//     "no-permission",
			//     "topic-not-exist"
			//   ],
			//   "type": "string"
			// }
			Description: "The status of the Amazon Redshift event notification subscription.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"subscription_creation_time": {
			// Property: SubscriptionCreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The date and time the Amazon Redshift event notification subscription was created.",
			//   "type": "string"
			// }
			Description: "The date and time the Amazon Redshift event notification subscription was created.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"subscription_name": {
			// Property: SubscriptionName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Amazon Redshift event notification subscription",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the Amazon Redshift event notification subscription",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "The `AWS::Redshift::EventSubscription` resource creates an Amazon Redshift Event Subscription.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Redshift::EventSubscription").WithTerraformTypeName("awscc_redshift_event_subscription")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cust_subscription_id":       "CustSubscriptionId",
		"customer_aws_id":            "CustomerAwsId",
		"enabled":                    "Enabled",
		"event_categories":           "EventCategories",
		"event_categories_list":      "EventCategoriesList",
		"key":                        "Key",
		"severity":                   "Severity",
		"sns_topic_arn":              "SnsTopicArn",
		"source_ids":                 "SourceIds",
		"source_ids_list":            "SourceIdsList",
		"source_type":                "SourceType",
		"status":                     "Status",
		"subscription_creation_time": "SubscriptionCreationTime",
		"subscription_name":          "SubscriptionName",
		"tags":                       "Tags",
		"value":                      "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}