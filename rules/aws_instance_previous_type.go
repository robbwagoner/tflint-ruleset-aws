package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-aws/project"
)

// AwsInstancePreviousTypeRule checks whether the resource uses previous generation instance type
type AwsInstancePreviousTypeRule struct {
	tflint.DefaultRule

	resourceType          string
	attributeName         string
	previousInstanceTypes map[string]bool
}

// NewAwsInstancePreviousTypeRule returns new rule with default attributes
func NewAwsInstancePreviousTypeRule() *AwsInstancePreviousTypeRule {
	return &AwsInstancePreviousTypeRule{
		resourceType:  "aws_instance",
		attributeName: "instance_type",
		previousInstanceTypes: map[string]bool{
			"t1.micro":    true,
			"m1.small":    true,
			"m1.medium":   true,
			"m1.large":    true,
			"m1.xlarge":   true,
			"m3.medium":   true,
			"m3.large":    true,
			"m3.xlarge":   true,
			"m3.2xlarge":  true,
			"c1.medium":   true,
			"c1.xlarge":   true,
			"cc2.8xlarge": true,
			"cg1.4xlarge": true,
			"c3.large":    true,
			"c3.xlarge":   true,
			"c3.2xlarge":  true,
			"c3.4xlarge":  true,
			"c3.8xlarge":  true,
			"g2.2xlarge":  true,
			"g2.8xlarge":  true,
			"m2.xlarge":   true,
			"m2.2xlarge":  true,
			"m2.4xlarge":  true,
			"cr1.8xlarge": true,
			"r3.large":    true,
			"r3.xlarge":   true,
			"r3.2xlarge":  true,
			"r3.4xlarge":  true,
			"r3.8xlarge":  true,
			"i2.xlarge":   true,
			"i2.2xlarge":  true,
			"i2.4xlarge":  true,
			"i2.8xlarge":  true,
			"hi1.4xlarge": true,
			"hs1.8xlarge": true,
		},
	}
}

// Name returns the rule name
func (r *AwsInstancePreviousTypeRule) Name() string {
	return "aws_instance_previous_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsInstancePreviousTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsInstancePreviousTypeRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *AwsInstancePreviousTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether the resource's `instance_type` is included in the list of previous generation instance type
func (r *AwsInstancePreviousTypeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var instanceType string
		err := runner.EvaluateExpr(attribute.Expr, &instanceType, nil)

		err = runner.EnsureNoError(err, func() error {
			if r.previousInstanceTypes[instanceType] {
				runner.EmitIssue(
					r,
					fmt.Sprintf("\"%s\" is previous generation instance type.", instanceType),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
