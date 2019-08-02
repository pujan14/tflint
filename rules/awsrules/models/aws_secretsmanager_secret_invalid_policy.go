// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsSecretsmanagerSecretInvalidPolicyRule checks the pattern is valid
type AwsSecretsmanagerSecretInvalidPolicyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSecretsmanagerSecretInvalidPolicyRule returns new rule with default attributes
func NewAwsSecretsmanagerSecretInvalidPolicyRule() *AwsSecretsmanagerSecretInvalidPolicyRule {
	return &AwsSecretsmanagerSecretInvalidPolicyRule{
		resourceType:  "aws_secretsmanager_secret",
		attributeName: "policy",
		max:           20480,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSecretsmanagerSecretInvalidPolicyRule) Name() string {
	return "aws_secretsmanager_secret_invalid_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSecretsmanagerSecretInvalidPolicyRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsSecretsmanagerSecretInvalidPolicyRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsSecretsmanagerSecretInvalidPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSecretsmanagerSecretInvalidPolicyRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"policy must be 20480 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"policy must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
