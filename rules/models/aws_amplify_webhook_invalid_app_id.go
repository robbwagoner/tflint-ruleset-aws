// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAmplifyWebhookInvalidAppIDRule checks the pattern is valid
type AwsAmplifyWebhookInvalidAppIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsAmplifyWebhookInvalidAppIDRule returns new rule with default attributes
func NewAwsAmplifyWebhookInvalidAppIDRule() *AwsAmplifyWebhookInvalidAppIDRule {
	return &AwsAmplifyWebhookInvalidAppIDRule{
		resourceType:  "aws_amplify_webhook",
		attributeName: "app_id",
		max:           20,
		min:           1,
		pattern:       regexp.MustCompile(`^d[a-z0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsAmplifyWebhookInvalidAppIDRule) Name() string {
	return "aws_amplify_webhook_invalid_app_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAmplifyWebhookInvalidAppIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAmplifyWebhookInvalidAppIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAmplifyWebhookInvalidAppIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAmplifyWebhookInvalidAppIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"app_id must be 20 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"app_id must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^d[a-z0-9]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
