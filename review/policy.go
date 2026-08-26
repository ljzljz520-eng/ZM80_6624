package review

import (
	"errors"
	"fmt"
)

func PolicyName(region string) string {
	switch region {
	case "EU":
		return "export-control"
	case "CN":
		return "heritage-clearance"
	default:
		return "local-review"
	}
}
func WrapReviewError(err error, stage string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", stage, err)
}
func IsRuleFailure(err error) bool { var re RuleError; return errors.As(err, &re) }
