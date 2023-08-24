// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package mediatailor

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_mediatailor_vod_source", vodSourceDataSource)
}

// vodSourceDataSource returns the Terraform awscc_mediatailor_vod_source data source.
// This Terraform data source corresponds to the CloudFormation AWS::MediaTailor::VodSource resource.
func vodSourceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe ARN of the VOD source.\u003c/p\u003e",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The ARN of the VOD source.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: HttpPackageConfigurations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eA list of HTTP package configuration parameters for this VOD source.\u003c/p\u003e",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "\u003cp\u003eThe HTTP package configuration properties for the requested VOD source.\u003c/p\u003e",
		//	    "properties": {
		//	      "Path": {
		//	        "description": "\u003cp\u003eThe relative path to the URL for this VOD source. This is combined with \u003ccode\u003eSourceLocation::HttpConfiguration::BaseUrl\u003c/code\u003e to form a valid URL.\u003c/p\u003e",
		//	        "type": "string"
		//	      },
		//	      "SourceGroup": {
		//	        "description": "\u003cp\u003eThe name of the source group. This has to match one of the \u003ccode\u003eChannel::Outputs::SourceGroup\u003c/code\u003e.\u003c/p\u003e",
		//	        "type": "string"
		//	      },
		//	      "Type": {
		//	        "enum": [
		//	          "DASH",
		//	          "HLS"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Path",
		//	      "SourceGroup",
		//	      "Type"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"http_package_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Path
					"path": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "<p>The relative path to the URL for this VOD source. This is combined with <code>SourceLocation::HttpConfiguration::BaseUrl</code> to form a valid URL.</p>",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: SourceGroup
					"source_group": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "<p>The name of the source group. This has to match one of the <code>Channel::Outputs::SourceGroup</code>.</p>",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Type
					"type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "<p>A list of HTTP package configuration parameters for this VOD source.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SourceLocationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"source_location_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags to assign to the VOD source.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags to assign to the VOD source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VodSourceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"vod_source_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::MediaTailor::VodSource",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaTailor::VodSource").WithTerraformTypeName("awscc_mediatailor_vod_source")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"http_package_configurations": "HttpPackageConfigurations",
		"key":                         "Key",
		"path":                        "Path",
		"source_group":                "SourceGroup",
		"source_location_name":        "SourceLocationName",
		"tags":                        "Tags",
		"type":                        "Type",
		"value":                       "Value",
		"vod_source_name":             "VodSourceName",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}