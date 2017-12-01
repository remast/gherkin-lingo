package rules

import (
	"strings"

	gherkin "github.com/cucumber/gherkin-go"
	lingo "github.com/remast/gherkin-lingo"
)

func init() {
	rule := &MissingScenarioName{}
	rule.Name = "missing-scenario-name"

	lingo.Register(rule)
}

type MissingScenarioName struct {
	lingo.EmptyRule
}

func (r *MissingScenarioName) CheckScenario(scenario *gherkin.Scenario) []*lingo.Issue {
	return checkScenarioName(scenario.Name, scenario.Location)
}

func (r *MissingScenarioName) CheckScenarioOutline(scenarioOutline *gherkin.ScenarioOutline) []*lingo.Issue {
	return checkScenarioName(scenarioOutline.Name, scenarioOutline.Location)
}

func checkScenarioName(name string, location *gherkin.Location) []*lingo.Issue {
	if strings.TrimSpace(name) == "" {
		issue := &lingo.Issue{
			Col:      location.Column,
			Line:     location.Line,
			Message:  "Scenario name missing.",
			Severity: lingo.Error,
		}
		return []*lingo.Issue{issue}
	}
	return nil
}
