package rules

import (
	"strings"

	gherkin "github.com/cucumber/gherkin-go"
	lingo "github.com/remast/gherkin-lingo"
)

func init() {
	rule := &GivenInFirstLineRule{}
	rule.Name = "given-in-first-line"

	lingo.Register(rule)
}

type GivenInFirstLineRule struct {
	lingo.EmptyRule
}

func (r *GivenInFirstLineRule) CheckSteps(steps []*gherkin.Step) []*lingo.Issue {
	issues := make([]*lingo.Issue, 0)
	for i, step := range steps {
		keyword := strings.ToLower(strings.TrimSpace(step.Keyword))
		if i == 0 && keyword != "given" {
			issue := &lingo.Issue{
				Col:      step.Location.Column,
				Line:     step.Location.Line,
				Message:  "First step should start with keyword Given.",
				Severity: lingo.Error,
			}
			issues = append(issues, issue)
		} else if i != 0 && keyword == "given" {
			issue := &lingo.Issue{
				Col:      step.Location.Column,
				Line:     step.Location.Line,
				Message:  "Given should only be used in first line.",
				Severity: lingo.Error,
			}
			issues = append(issues, issue)

		}
	}
	return issues
}
