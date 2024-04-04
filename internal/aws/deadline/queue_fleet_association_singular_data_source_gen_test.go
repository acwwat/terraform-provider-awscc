// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package deadline_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSDeadlineQueueFleetAssociationDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Deadline::QueueFleetAssociation", "awscc_deadline_queue_fleet_association", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSDeadlineQueueFleetAssociationDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Deadline::QueueFleetAssociation", "awscc_deadline_queue_fleet_association", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
