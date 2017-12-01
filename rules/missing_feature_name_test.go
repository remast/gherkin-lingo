package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	gherkin "github.com/cucumber/gherkin-go"
	lingo "github.com/remast/gherkin-lingo"
)

func TestFeatureWithMissingFeatureName(t *testing.T) {
	const filepath = "/tmp/file"

	rule := &MissingFeatureName{}
	var document *gherkin.GherkinDocument
	var issues []*lingo.Issue
	var issue *lingo.Issue

	t.Run("ScenarioWithoutMissingFeatureName", func(t *testing.T) {
		featureWithoutFeatureName :=
			`Feature: Feature 1 
		
		Scenario: Scenario 1

		Given a step
		When a step
		Then a step
		`

		document, _ = gherkin.ParseGherkinDocument(strings.NewReader(featureWithoutFeatureName))
		issues = lingo.ApplyRule(rule, filepath, document)

		assert.Equal(t, 0, len(issues))
	})

	t.Run("FeatureWithMissingFeatureName", func(t *testing.T) {
		featureWithMissingFeatureName :=
			`Feature: 
		
		Scenario: Test 1

		And a step
		When a step
		Then a step
		`

		document, _ = gherkin.ParseGherkinDocument(strings.NewReader(featureWithMissingFeatureName))
		issues = lingo.ApplyRule(rule, filepath, document)

		assert.Equal(t, 1, len(issues))
		issue = issues[0]

		assert.Equal(t, lingo.Error, issue.Severity)
		assert.Equal(t, 1, issue.Col)
	})
}
