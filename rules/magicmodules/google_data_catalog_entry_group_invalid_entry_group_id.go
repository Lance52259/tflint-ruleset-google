// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package magicmodules

import (
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleDataCatalogEntryGroupInvalidEntryGroupIdRule checks the pattern is valid
type GoogleDataCatalogEntryGroupInvalidEntryGroupIdRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleDataCatalogEntryGroupInvalidEntryGroupIdRule returns new rule with default attributes
func NewGoogleDataCatalogEntryGroupInvalidEntryGroupIdRule() *GoogleDataCatalogEntryGroupInvalidEntryGroupIdRule {
	return &GoogleDataCatalogEntryGroupInvalidEntryGroupIdRule{
		resourceType:  "google_data_catalog_entry_group",
		attributeName: "entry_group_id",
	}
}

// Name returns the rule name
func (r *GoogleDataCatalogEntryGroupInvalidEntryGroupIdRule) Name() string {
	return "google_data_catalog_entry_group_invalid_entry_group_id"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleDataCatalogEntryGroupInvalidEntryGroupIdRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleDataCatalogEntryGroupInvalidEntryGroupIdRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleDataCatalogEntryGroupInvalidEntryGroupIdRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleDataCatalogEntryGroupInvalidEntryGroupIdRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		validateFunc := validateRegexp(`^[A-z_][A-z0-9_]{0,63}$`)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}