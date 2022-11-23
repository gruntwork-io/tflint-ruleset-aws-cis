package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
	"testing"
)

func TestDoesPortRangeContainsRemoteAccess(t *testing.T) {


	tests := []struct {
				Name     string
				Content  string
				Expected helper.Issues
			}{
				{
					Name: "issue found",
					Content: `
		resource "aws_instance" "web" {
		   instance_type = "t2.m
	doesPortRangeContainsRemoteAccess
}
