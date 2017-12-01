package rules

import (
	"strings"

	gherkin "github.com/cucumber/gherkin-go"
	lingo "github.com/remast/gherkin-lingo"
)

func init() {
	rule := &MissingFeatureName{}
	rule.Name = "missing-feature-name"

	lingo.Register(rule)
}

type MissingFeatureName struct {
	lingo.EmptyRule
}

func (r *MissingFeatureName) CheckFeature(feature *gherkin.Feature) []*lingo.Issue {
	if strings.TrimSpace(feature.Name) == "" {
		issue := &lingo.Issue{
			Col:      feature.Location.Column,
			Line:     feature.Location.Line,
			Message:  "Feature name missing.",
			Severity: lingo.Error,
		}
		return []*lingo.Issue{issue}
	}
	return nil
}
