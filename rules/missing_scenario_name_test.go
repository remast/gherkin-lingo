package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	gherkin "github.com/cucumber/gherkin-go"
	lingo "github.com/remast/gherkin-lingo"
)

func TestScenarioWithMissingScenarioName(t *testing.T) {
	const filepath = "/tmp/file"

	rule := &MissingScenarioName{}
	var document *gherkin.GherkinDocument
	var issues []*lingo.Issue
	var issue *lingo.Issue

	t.Run("ScenarioWithoutMissingScenarioName", func(t *testing.T) {
		scenarioWithoutMissingScenarioName :=
			`Feature: Feature 1 
		
		Scenario: Scenario 1

		And a step
		When a step
		Then a step
		`

		document, _ = gherkin.ParseGherkinDocument(strings.NewReader(scenarioWithoutMissingScenarioName))
		issues = lingo.ApplyRule(rule, filepath, document)

		assert.Equal(t, 0, len(issues))
	})

	t.Run("ScenarioWithMissingScenarioName", func(t *testing.T) {
		scenarioWithMissingScenarioName :=
			`Feature: Feature 1 
		
		Scenario: 

		And a step
		When a step
		Then a step
		`

		document, _ = gherkin.ParseGherkinDocument(strings.NewReader(scenarioWithMissingScenarioName))
		issues = lingo.ApplyRule(rule, filepath, document)

		assert.Equal(t, 1, len(issues))
		issue = issues[0]

		assert.Equal(t, lingo.Error, issue.Severity)
		assert.Equal(t, 3, issue.Line)
	})

	t.Run("ScenarioOutlineWithMissingScenarioName", func(t *testing.T) {
		scenarioOutlineWithMissingScenarioName :=
			`Feature: Feature 1 
		
		Scenario Outline: 

		And a step
		When a step
		Then a step
		`

		document, _ = gherkin.ParseGherkinDocument(strings.NewReader(scenarioOutlineWithMissingScenarioName))
		issues = lingo.ApplyRule(rule, filepath, document)

		assert.Equal(t, 1, len(issues))
		issue = issues[0]

		assert.Equal(t, lingo.Error, issue.Severity)
		assert.Equal(t, 3, issue.Line)
	})
}
