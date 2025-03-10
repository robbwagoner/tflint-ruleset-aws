// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodestarconnectionsConnectionInvalidProviderTypeRule checks the pattern is valid
type AwsCodestarconnectionsConnectionInvalidProviderTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCodestarconnectionsConnectionInvalidProviderTypeRule returns new rule with default attributes
func NewAwsCodestarconnectionsConnectionInvalidProviderTypeRule() *AwsCodestarconnectionsConnectionInvalidProviderTypeRule {
	return &AwsCodestarconnectionsConnectionInvalidProviderTypeRule{
		resourceType:  "aws_codestarconnections_connection",
		attributeName: "provider_type",
		enum: []string{
			"Bitbucket",
			"GitHub",
			"GitHubEnterpriseServer",
		},
	}
}

// Name returns the rule name
func (r *AwsCodestarconnectionsConnectionInvalidProviderTypeRule) Name() string {
	return "aws_codestarconnections_connection_invalid_provider_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodestarconnectionsConnectionInvalidProviderTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodestarconnectionsConnectionInvalidProviderTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodestarconnectionsConnectionInvalidProviderTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodestarconnectionsConnectionInvalidProviderTypeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as provider_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
