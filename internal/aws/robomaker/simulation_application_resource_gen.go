// Code generated by generators/resource/main.go; DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_robomaker_simulation_application", simulationApplicationResourceType)
}

// simulationApplicationResourceType returns the Terraform awscc_robomaker_simulation_application resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::RoboMaker::SimulationApplication resource type.
func simulationApplicationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"current_revision_id": {
			// Property: CurrentRevisionId
			// CloudFormation resource type schema:
			// {
			//   "description": "The current revision id.",
			//   "type": "string"
			// }
			Description: "The current revision id.",
			Type:        types.StringType,
			Optional:    true,
		},
		"environment": {
			// Property: Environment
			// CloudFormation resource type schema:
			// {
			//   "description": "The URI of the Docker image for the robot application.",
			//   "type": "string"
			// }
			Description: "The URI of the Docker image for the robot application.",
			Type:        types.StringType,
			Optional:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the simulation application.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the simulation application.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"rendering_engine": {
			// Property: RenderingEngine
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Information about a rendering engine.",
			//   "properties": {
			//     "Name": {
			//       "description": "The name of the rendering engine.",
			//       "enum": [
			//         "OGRE"
			//       ],
			//       "type": "string"
			//     },
			//     "Version": {
			//       "description": "The version of the rendering engine.",
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Name",
			//     "Version"
			//   ],
			//   "type": "object"
			// }
			Description: "Information about a rendering engine.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Description: "The name of the rendering engine.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"OGRE",
							}),
						},
					},
					"version": {
						// Property: Version
						Description: "The version of the rendering engine.",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Optional: true,
		},
		"robot_software_suite": {
			// Property: RobotSoftwareSuite
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Information about a robot software suite (ROS distribution).",
			//   "properties": {
			//     "Name": {
			//       "description": "The name of the robot software suite (ROS distribution).",
			//       "enum": [
			//         "ROS",
			//         "ROS2"
			//       ],
			//       "type": "string"
			//     },
			//     "Version": {
			//       "description": "The version of the robot software suite (ROS distribution).",
			//       "enum": [
			//         "Kinetic",
			//         "Melodic",
			//         "Dashing",
			//         "Foxy"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Name",
			//     "Version"
			//   ],
			//   "type": "object"
			// }
			Description: "Information about a robot software suite (ROS distribution).",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Description: "The name of the robot software suite (ROS distribution).",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"ROS",
								"ROS2",
							}),
						},
					},
					"version": {
						// Property: Version
						Description: "The version of the robot software suite (ROS distribution).",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"Kinetic",
								"Melodic",
								"Dashing",
								"Foxy",
							}),
						},
					},
				},
			),
			Required: true,
		},
		"simulation_software_suite": {
			// Property: SimulationSoftwareSuite
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Information about a simulation software suite.",
			//   "properties": {
			//     "Name": {
			//       "description": "The name of the simulation software suite.",
			//       "enum": [
			//         "Gazebo",
			//         "RosbagPlay"
			//       ],
			//       "type": "string"
			//     },
			//     "Version": {
			//       "description": "The version of the simulation software suite.",
			//       "enum": [
			//         "7",
			//         "9",
			//         "11",
			//         "Kinetic",
			//         "Melodic",
			//         "Dashing",
			//         "Foxy"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Name",
			//     "Version"
			//   ],
			//   "type": "object"
			// }
			Description: "Information about a simulation software suite.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Description: "The name of the simulation software suite.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"Gazebo",
								"RosbagPlay",
							}),
						},
					},
					"version": {
						// Property: Version
						Description: "The version of the simulation software suite.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"7",
								"9",
								"11",
								"Kinetic",
								"Melodic",
								"Dashing",
								"Foxy",
							}),
						},
					},
				},
			),
			Required: true,
		},
		"sources": {
			// Property: Sources
			// CloudFormation resource type schema:
			// {
			//   "description": "The sources of the simulation application.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Information about a source configuration.",
			//     "properties": {
			//       "Architecture": {
			//         "description": "The target processor architecture for the application.",
			//         "enum": [
			//           "X86_64",
			//           "ARM64",
			//           "ARMHF"
			//         ],
			//         "type": "string"
			//       },
			//       "S3Bucket": {
			//         "description": "The Amazon S3 bucket name.",
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "S3Key": {
			//         "description": "The s3 object key.",
			//         "maxLength": 1024,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "S3Bucket",
			//       "S3Key",
			//       "Architecture"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The sources of the simulation application.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"architecture": {
						// Property: Architecture
						Description: "The target processor architecture for the application.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"X86_64",
								"ARM64",
								"ARMHF",
							}),
						},
					},
					"s3_bucket": {
						// Property: S3Bucket
						Description: "The Amazon S3 bucket name.",
						Type:        types.StringType,
						Required:    true,
					},
					"s3_key": {
						// Property: S3Key
						Description: "The s3 object key.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 1024),
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
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A key-value pair to associate with a resource.",
			//   "patternProperties": {
			//     "": {
			//       "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A key-value pair to associate with a resource.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "An example resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::RoboMaker::SimulationApplication").WithTerraformTypeName("awscc_robomaker_simulation_application")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"architecture":              "Architecture",
		"arn":                       "Arn",
		"current_revision_id":       "CurrentRevisionId",
		"environment":               "Environment",
		"name":                      "Name",
		"rendering_engine":          "RenderingEngine",
		"robot_software_suite":      "RobotSoftwareSuite",
		"s3_bucket":                 "S3Bucket",
		"s3_key":                    "S3Key",
		"simulation_software_suite": "SimulationSoftwareSuite",
		"sources":                   "Sources",
		"tags":                      "Tags",
		"version":                   "Version",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
