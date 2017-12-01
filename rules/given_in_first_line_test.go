package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	lingo "github.com/remast/gherkin-lingo"

	gherkin "github.com/cucumber/gherkin-go"
)

func TestGivenInFirstLine(t *testing.T) {
	const filepath = "/tmp/file"

	rule := &GivenInFirstLineRule{}
	var document *gherkin.GherkinDocument
	var issues []*lingo.Issue
	var issue *lingo.Issue

	t.Run("FeatureWithGivenInFirstLine", func(t *testing.T) {
		featureWithGivenInFirstLine :=
			`Feature: Feature 1 
		
		Scenario: Scenario 1

		Given a step
		When a step
		Then a step
		`

		document, _ = gherkin.ParseGherkinDocument(strings.NewReader(featureWithGivenInFirstLine))
		issues = lingo.ApplyRule(rule, filepath, document)

		assert.Equal(t, 0, len(issues))
	})

	t.Run("FeatureWithoutGivenInFirstLine", func(t *testing.T) {
		featureWithNoGivenInFirstLine :=
			`Feature: Test 1
		
		Scenario: Test 1

		And a step
		When a step
		Then a step
		Given another step
		`

		document, _ = gherkin.ParseGherkinDocument(strings.NewReader(featureWithNoGivenInFirstLine))
		issues = lingo.ApplyRule(rule, filepath, document)

		assert.Equal(t, 2, len(issues))
		issue = issues[0]

		assert.Equal(t, lingo.Error, issue.Severity)
		assert.Equal(t, 5, issue.Line)

		issue = issues[1]
		assert.Equal(t, lingo.Error, issue.Severity)
		assert.Equal(t, 8, issue.Line)
	})
}
