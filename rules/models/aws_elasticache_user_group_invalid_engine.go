// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsElastiCacheUserGroupInvalidEngineRule checks the pattern is valid
type AwsElastiCacheUserGroupInvalidEngineRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsElastiCacheUserGroupInvalidEngineRule returns new rule with default attributes
func NewAwsElastiCacheUserGroupInvalidEngineRule() *AwsElastiCacheUserGroupInvalidEngineRule {
	return &AwsElastiCacheUserGroupInvalidEngineRule{
		resourceType:  "aws_elasticache_user_group",
		attributeName: "engine",
		pattern:       regexp.MustCompile(`^[a-zA-Z]*$`),
	}
}

// Name returns the rule name
func (r *AwsElastiCacheUserGroupInvalidEngineRule) Name() string {
	return "aws_elasticache_user_group_invalid_engine"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElastiCacheUserGroupInvalidEngineRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElastiCacheUserGroupInvalidEngineRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElastiCacheUserGroupInvalidEngineRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElastiCacheUserGroupInvalidEngineRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z]*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
