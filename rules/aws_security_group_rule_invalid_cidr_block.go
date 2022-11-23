package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSecurityGroupRuleInvalidCidrBlockRule checks whether ...
type AwsSecurityGroupRuleInvalidCidrBlockRule struct {
	tflint.DefaultRule

	resourceType string
	cidrBlocks   []string
}

// NewAwsSecurityGroupRuleInvalidCidrBlockRule returns a new rule
func NewAwsSecurityGroupRuleInvalidCidrBlockRule() *AwsSecurityGroupRuleInvalidCidrBlockRule {
	return &AwsSecurityGroupRuleInvalidCidrBlockRule{}
}

// Name returns the rule name
func (r *AwsSecurityGroupRuleInvalidCidrBlockRule) Name() string {
	return "aws_security_group_rule_invalid_cidr_block"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSecurityGroupRuleInvalidCidrBlockRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSecurityGroupRuleInvalidCidrBlockRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSecurityGroupRuleInvalidCidrBlockRule) Link() string {
	return "TODO"
}

// Check checks whether ...
func (r *AwsSecurityGroupRuleInvalidCidrBlockRule) Check(runner tflint.Runner) error {
	// This rule is an example to get a top-level resource attribute.
	resources, err := runner.GetResourceContent("aws_security_group_rule", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "cidr_blocks"},
			{Name: "ipv6_cidr_blocks"},
			{Name: "from_port"},
			{Name: "to_port"},
			{Name: "type"},
		},
	}, nil)
	if err != nil {
		return err
	}

	// Put a log that can be output with `TFLINT_LOG=debug`
	logger.Debug(fmt.Sprintf("Get %d security groups", len(resources.Blocks)))

	for _, resource := range resources.Blocks {
		typeAttribute, exists := resource.Body.Attributes["type"]
		if !exists {
			continue
		}

		var sgType string
		err := runner.EvaluateExpr(typeAttribute.Expr, &sgType, nil)

		logger.Debug(fmt.Sprintf("Found %s type of Security Group Rule", sgType))

		if sgType != "ingress" {
			continue
		}

		fromPortAttribute, exists := resource.Body.Attributes["from_port"]
		if !exists {
			continue
		}

		toPortAttribute, exists := resource.Body.Attributes["to_port"]
		if !exists {
			continue
		}

		// TODO what means that each field doesnt exist?
		// Case 1:
		// From 0, to 0.

		var fromPort int
		err = runner.EvaluateExpr(fromPortAttribute.Expr, &fromPort, nil)

		var toPort int
		err = runner.EvaluateExpr(toPortAttribute.Expr, &toPort, nil)

		logger.Debug(fmt.Sprintf("Security Group Rule port range is %d to %d", fromPort, toPort))

		// No need to check the CIDR blocks if the ports range does not contain 22 or 3389.
		if !doesPortRangeContainsRemoteAccess(fromPort, toPort) {
			continue
		}

		ipv6CidrBlocksAttribute, exists := resource.Body.Attributes["ipv6_cidr_blocks"]
		if exists {
			var ipv6CidrBlocks []string
			err = runner.EvaluateExpr(ipv6CidrBlocksAttribute.Expr, &ipv6CidrBlocks, nil)

			//err = runner.EnsureNoError(err, func() error {
			if doesCidrBlocksContainAll(ipv6CidrBlocks) {
				return runner.EmitIssue(
					r,
					fmt.Sprintf("cidr_blocks are %v", ipv6CidrBlocks),
					ipv6CidrBlocksAttribute.Expr.Range(),
				)
			} //)

		}

		cidrBlocksAttribute, exists := resource.Body.Attributes["ipv6_cidr_blocks"]
		if exists {
			var cidrBlocks []string
			err = runner.EvaluateExpr(cidrBlocksAttribute.Expr, &cidrBlocks, nil)

			//err = runner.EnsureNoError(err, func() error {
			if doesCidrBlocksContainAll(cidrBlocks) {

				return runner.EmitIssue(
					r,
					fmt.Sprintf("cidr_blocks are %v", cidrBlocks),
					cidrBlocksAttribute.Expr.Range(),
				)
			}
			//})
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func doesPortRangeContainsRemoteAccess(fromPort int, toPort int) bool {
	remoteAccessPorts := []int{22, 3389}

	for _, port := range remoteAccessPorts {
		isIncludedInRange := fromPort <= port && port <= toPort
		logger.Debug(fmt.Sprintf("%v for %d isIncludedInRange", isIncludedInRange, port))
		if isIncludedInRange {
			return true
		}
	}

	return false
}

func doesCidrBlocksContainAll(cidrBlocks []string) bool {
	return true

}
