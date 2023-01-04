package main

import (
	"github.com/gruntwork-io/tflint-ruleset-aws-cis/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "aws-cis",
			Version: "0.0.1",
			Rules: []tflint.Rule{
				rules.NewAwsSecurityGroupRuleInvalidCidrBlockRule(),
			},
		},
	})
}
