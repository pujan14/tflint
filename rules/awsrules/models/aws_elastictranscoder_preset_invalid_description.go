// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsElastictranscoderPresetInvalidDescriptionRule checks the pattern is valid
type AwsElastictranscoderPresetInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsElastictranscoderPresetInvalidDescriptionRule returns new rule with default attributes
func NewAwsElastictranscoderPresetInvalidDescriptionRule() *AwsElastictranscoderPresetInvalidDescriptionRule {
	return &AwsElastictranscoderPresetInvalidDescriptionRule{
		resourceType:  "aws_elastictranscoder_preset",
		attributeName: "description",
		max:           255,
	}
}

// Name returns the rule name
func (r *AwsElastictranscoderPresetInvalidDescriptionRule) Name() string {
	return "aws_elastictranscoder_preset_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElastictranscoderPresetInvalidDescriptionRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsElastictranscoderPresetInvalidDescriptionRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsElastictranscoderPresetInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElastictranscoderPresetInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}