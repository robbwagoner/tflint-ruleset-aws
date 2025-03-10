package rules

import (
	"encoding/json"
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-aws/project"
)

type AwsIAMPolicySidInvalidCharactersPolicyStatement struct {
	Sid string `json:"Sid"`
}
type AwsIAMPolicySidInvalidCharactersPolicy struct {
	Statement []AwsIAMPolicySidInvalidCharactersPolicyStatement `json:"Statement"`
}
type AwsIAMPolicySidInvalidCharactersPolicyWithSingleStatement struct {
	Statement AwsIAMPolicySidInvalidCharactersPolicyStatement `json:"Statement"`
}

// AwsIAMPolicySidInvalidCharactersRule checks for invalid characters in SID
type AwsIAMPolicySidInvalidCharactersRule struct {
	resourceType    string
	attributeName   string
	validCharacters *regexp.Regexp
}

// NewAwsIAMPolicySidInvalidCharactersRule returns new rule with default attributes
func NewAwsIAMPolicySidInvalidCharactersRule() *AwsIAMPolicySidInvalidCharactersRule {
	return &AwsIAMPolicySidInvalidCharactersRule{
		resourceType:    "aws_iam_policy",
		attributeName:   "policy",
		validCharacters: regexp.MustCompile(`^[a-zA-Z0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMPolicySidInvalidCharactersRule) Name() string {
	return "aws_iam_policy_sid_invalid_characters"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMPolicySidInvalidCharactersRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMPolicySidInvalidCharactersRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMPolicySidInvalidCharactersRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the unmarshaled policy and loops through statements checking for invalid statement ids
func (r *AwsIAMPolicySidInvalidCharactersRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			var statements []AwsIAMPolicySidInvalidCharactersPolicyStatement

			policy := AwsIAMPolicySidInvalidCharactersPolicy{}
			if err := json.Unmarshal([]byte(val), &policy); err != nil {
				// If the Statement clause includes only one value, you can omit the brackets, so try unmarshal to the struct accordingly.
				// https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_grammar.html
				policy := AwsIAMPolicySidInvalidCharactersPolicyWithSingleStatement{}
				if err := json.Unmarshal([]byte(val), &policy); err != nil {
					return err
				}
				statements = []AwsIAMPolicySidInvalidCharactersPolicyStatement{policy.Statement}
			} else {
				statements = policy.Statement
			}

			for _, statement := range statements {
				if statement.Sid == "" {
					continue
				}

				if !r.validCharacters.MatchString(statement.Sid) {
					runner.EmitIssueOnExpr(
						r,
						fmt.Sprintf("The policy's sid (\"%s\") does not match \"%s\".", statement.Sid, r.validCharacters.String()),
						attribute.Expr,
					)
				}
			}
			return nil
		})
	})
}
