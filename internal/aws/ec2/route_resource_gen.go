// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_route", routeResource)
}

// routeResource returns the Terraform awscc_ec2_route resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::Route resource.
func routeResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CarrierGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the carrier gateway.\n You can only use this option when the VPC contains a subnet which is associated with a Wavelength Zone.",
		//	  "type": "string"
		//	}
		"carrier_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the carrier gateway.\n You can only use this option when the VPC contains a subnet which is associated with a Wavelength Zone.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CidrBlock
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CoreNetworkArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the core network.",
		//	  "type": "string"
		//	}
		"core_network_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the core network.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DestinationCidrBlock
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IPv4 CIDR address block used for the destination match. Routing decisions are based on the most specific match. We modify the specified CIDR block to its canonical form; for example, if you specify ``100.68.0.18/18``, we modify it to ``100.68.0.0/18``.",
		//	  "type": "string"
		//	}
		"destination_cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IPv4 CIDR address block used for the destination match. Routing decisions are based on the most specific match. We modify the specified CIDR block to its canonical form; for example, if you specify ``100.68.0.18/18``, we modify it to ``100.68.0.0/18``.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DestinationIpv6CidrBlock
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IPv6 CIDR block used for the destination match. Routing decisions are based on the most specific match.",
		//	  "type": "string"
		//	}
		"destination_ipv_6_cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IPv6 CIDR block used for the destination match. Routing decisions are based on the most specific match.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DestinationPrefixListId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of a prefix list used for the destination match.",
		//	  "type": "string"
		//	}
		"destination_prefix_list_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of a prefix list used for the destination match.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EgressOnlyInternetGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "[IPv6 traffic only] The ID of an egress-only internet gateway.",
		//	  "type": "string"
		//	}
		"egress_only_internet_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "[IPv6 traffic only] The ID of an egress-only internet gateway.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of an internet gateway or virtual private gateway attached to your VPC.",
		//	  "type": "string"
		//	}
		"gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of an internet gateway or virtual private gateway attached to your VPC.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InstanceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of a NAT instance in your VPC. The operation fails if you specify an instance ID unless exactly one network interface is attached.",
		//	  "type": "string"
		//	}
		"instance_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of a NAT instance in your VPC. The operation fails if you specify an instance ID unless exactly one network interface is attached.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LocalGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the local gateway.",
		//	  "type": "string"
		//	}
		"local_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the local gateway.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NatGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "[IPv4 traffic only] The ID of a NAT gateway.",
		//	  "type": "string"
		//	}
		"nat_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "[IPv4 traffic only] The ID of a NAT gateway.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NetworkInterfaceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of a network interface.",
		//	  "type": "string"
		//	}
		"network_interface_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of a network interface.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RouteTableId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the route table for the route.",
		//	  "type": "string"
		//	}
		"route_table_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the route table for the route.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of a transit gateway.",
		//	  "type": "string"
		//	}
		"transit_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of a transit gateway.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcEndpointId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of a VPC endpoint. Supported for Gateway Load Balancer endpoints only.",
		//	  "type": "string"
		//	}
		"vpc_endpoint_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of a VPC endpoint. Supported for Gateway Load Balancer endpoints only.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcPeeringConnectionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of a VPC peering connection.",
		//	  "type": "string"
		//	}
		"vpc_peering_connection_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of a VPC peering connection.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Specifies a route in a route table. For more information, see [Routes](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html#route-table-routes) in the *Amazon VPC User Guide*.\n You must specify either a destination CIDR block or prefix list ID. You must also specify exactly one of the resources as the target.\n If you create a route that references a transit gateway in the same template where you create the transit gateway, you must declare a dependency on the transit gateway attachment. The route table cannot use the transit gateway until it has successfully attached to the VPC. Add a [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) in the ``AWS::EC2::Route`` resource to explicitly declare a dependency on the ``AWS::EC2::TransitGatewayAttachment`` resource.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::Route").WithTerraformTypeName("awscc_ec2_route")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"carrier_gateway_id":              "CarrierGatewayId",
		"cidr_block":                      "CidrBlock",
		"core_network_arn":                "CoreNetworkArn",
		"destination_cidr_block":          "DestinationCidrBlock",
		"destination_ipv_6_cidr_block":    "DestinationIpv6CidrBlock",
		"destination_prefix_list_id":      "DestinationPrefixListId",
		"egress_only_internet_gateway_id": "EgressOnlyInternetGatewayId",
		"gateway_id":                      "GatewayId",
		"instance_id":                     "InstanceId",
		"local_gateway_id":                "LocalGatewayId",
		"nat_gateway_id":                  "NatGatewayId",
		"network_interface_id":            "NetworkInterfaceId",
		"route_table_id":                  "RouteTableId",
		"transit_gateway_id":              "TransitGatewayId",
		"vpc_endpoint_id":                 "VpcEndpointId",
		"vpc_peering_connection_id":       "VpcPeeringConnectionId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
