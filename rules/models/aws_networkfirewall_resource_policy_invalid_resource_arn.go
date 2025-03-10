// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsNetworkfirewallResourcePolicyInvalidResourceArnRule checks the pattern is valid
type AwsNetworkfirewallResourcePolicyInvalidResourceArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsNetworkfirewallResourcePolicyInvalidResourceArnRule returns new rule with default attributes
func NewAwsNetworkfirewallResourcePolicyInvalidResourceArnRule() *AwsNetworkfirewallResourcePolicyInvalidResourceArnRule {
	return &AwsNetworkfirewallResourcePolicyInvalidResourceArnRule{
		resourceType:  "aws_networkfirewall_resource_policy",
		attributeName: "resource_arn",
		max:           256,
		min:           1,
		pattern:       regexp.MustCompile(`^arn:aws.*`),
	}
}

// Name returns the rule name
func (r *AwsNetworkfirewallResourcePolicyInvalidResourceArnRule) Name() string {
	return "aws_networkfirewall_resource_policy_invalid_resource_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsNetworkfirewallResourcePolicyInvalidResourceArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsNetworkfirewallResourcePolicyInvalidResourceArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsNetworkfirewallResourcePolicyInvalidResourceArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsNetworkfirewallResourcePolicyInvalidResourceArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"resource_arn must be 256 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"resource_arn must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:aws.*`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
