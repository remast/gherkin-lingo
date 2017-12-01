package rules

import (
	gherkin "github.com/cucumber/gherkin-go"
	lingo "github.com/remast/gherkin-lingo"
)

func init() {
	rule := &OutlineWithTooFewExamples{}
	rule.Name = "outline-with-too-few-examples"

	lingo.Register(rule)
}

type OutlineWithTooFewExamples struct {
	lingo.EmptyRule
}

func (r *OutlineWithTooFewExamples) CheckScenarioOutline(scenarioOutline *gherkin.ScenarioOutline) []*lingo.Issue {
	issues := make([]*lingo.Issue, 0)
	for _, examples := range scenarioOutline.Examples {
		if len(examples.TableBody) <= 1 {
			issue := &lingo.Issue{
				Col:      scenarioOutline.Location.Column,
				Line:     scenarioOutline.Location.Line,
				Message:  "Scenario Outline with too few examples (maybe use a Scenario instead).",
				Severity: lingo.Error,
			}
			issues = append(issues, issue)
		}
	}
	return issues
}
