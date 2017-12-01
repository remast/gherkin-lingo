package rules

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	gherkin "github.com/cucumber/gherkin-go"
	lingo "github.com/remast/gherkin-lingo"
)

func TestScenarioWithOutlineWithTooFewExamples(t *testing.T) {
	const filepath = "/tmp/file"

	rule := &OutlineWithTooFewExamples{}
	var document *gherkin.GherkinDocument
	var issues []*lingo.Issue
	var issue *lingo.Issue

	t.Run("ScenarioOutlineWithTooFewExamples", func(t *testing.T) {
		scenarioWithTooFewExamples :=
			`Feature: Feature 1 
		
		Scenario Outline: ScenarioOutlineWithTooFewExamples

		Given a step with <Col1>
		When a step with <Col2>
		Then a step shows <Result>

		Examples:
		 | Col1 | Col2 | Result |
		 | 1    | 2   | 3      |
		 `

		document, err := gherkin.ParseGherkinDocument(strings.NewReader(scenarioWithTooFewExamples))
		fmt.Print(err)
		issues = lingo.ApplyRule(rule, filepath, document)

		assert.Equal(t, 1, len(issues))
		issue = issues[0]

		assert.Equal(t, lingo.Error, issue.Severity)
		assert.Equal(t, 3, issue.Line)
	})

	t.Run("ScenarioOutlineWithEnoughExamples", func(t *testing.T) {
		scenarioOutlineWithEnoughExamples :=
			`Feature: Feature 1 
		
		Scenario Outline:  ScenarioOutlineWithEnoughExamples

		Given a step with <Col1>
		When a step with <Col2>
		Then a step shows <Result>

		Examples:
		 | Col1 | Col2 | Result |
 		 | 1    | 2    | 3      |
		 | 2    | 3    | 5      |
		 `

		document, _ = gherkin.ParseGherkinDocument(strings.NewReader(scenarioOutlineWithEnoughExamples))
		issues = lingo.ApplyRule(rule, filepath, document)

		assert.Equal(t, 0, len(issues))
	})
}
