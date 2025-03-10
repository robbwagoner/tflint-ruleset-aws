// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServiceDiscoveryInstanceInvalidInstanceIDRule checks the pattern is valid
type AwsServiceDiscoveryInstanceInvalidInstanceIDRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsServiceDiscoveryInstanceInvalidInstanceIDRule returns new rule with default attributes
func NewAwsServiceDiscoveryInstanceInvalidInstanceIDRule() *AwsServiceDiscoveryInstanceInvalidInstanceIDRule {
	return &AwsServiceDiscoveryInstanceInvalidInstanceIDRule{
		resourceType:  "aws_service_discovery_instance",
		attributeName: "instance_id",
		max:           64,
		pattern:       regexp.MustCompile(`^[0-9a-zA-Z_/:.@-]+$`),
	}
}

// Name returns the rule name
func (r *AwsServiceDiscoveryInstanceInvalidInstanceIDRule) Name() string {
	return "aws_service_discovery_instance_invalid_instance_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServiceDiscoveryInstanceInvalidInstanceIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServiceDiscoveryInstanceInvalidInstanceIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServiceDiscoveryInstanceInvalidInstanceIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServiceDiscoveryInstanceInvalidInstanceIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"instance_id must be 64 characters or less",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[0-9a-zA-Z_/:.@-]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
