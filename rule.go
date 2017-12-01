package lingo

import (
	gherkin "github.com/cucumber/gherkin-go"
)

var rules []Rule

func Register(rule Rule) {
	rules = append(rules, rule)
}

func Apply(filepath string, document *gherkin.GherkinDocument) []*Issue {
	issues := make([]*Issue, 0)
	for _, rule := range rules {
		ruleIssues := ApplyRule(rule, filepath, document)
		issues = append(issues, ruleIssues...)
	}
	return issues
}

func ApplyRule(rule Rule, filepath string, document *gherkin.GherkinDocument) []*Issue {
	issues := make([]*Issue, 0)

	documentIssues := rule.CheckDocument(document)
	issues = append(issues, documentIssues...)

	featureIssues := rule.CheckFeature(document.Feature)
	issues = append(issues, featureIssues...)

	for _, child := range document.Feature.Children {
		scenario, ok := child.(*gherkin.Scenario)
		if ok {
			scenarioIssues := rule.CheckScenario(scenario)
			issues = append(issues, scenarioIssues...)

			stepsIssues := rule.CheckSteps(scenario.ScenarioDefinition.Steps)
			issues = append(issues, stepsIssues...)

			continue
		}

		scenarioOutline, ok := child.(*gherkin.ScenarioOutline)
		if ok {
			scenarioOutlineIssues := rule.CheckScenarioOutline(scenarioOutline)
			issues = append(issues, scenarioOutlineIssues...)

			stepsIssues := rule.CheckSteps(scenarioOutline.ScenarioDefinition.Steps)
			issues = append(issues, stepsIssues...)

			continue
		}

	}

	for _, issue := range issues {
		issue.Linter = rule.String()
		issue.Path = filepath
	}

	return issues
}

type Rule interface {
	CheckDocument(gherkin *gherkin.GherkinDocument) []*Issue
	CheckFeature(feature *gherkin.Feature) []*Issue
	CheckScenario(scenario *gherkin.Scenario) []*Issue
	CheckScenarioOutline(scenarioOutline *gherkin.ScenarioOutline) []*Issue
	CheckSteps(steps []*gherkin.Step) []*Issue
	String() string
}

type EmptyRule struct {
	Rule
	Name string
}

func (r *EmptyRule) String() string {
	return r.Name
}

func (r *EmptyRule) CheckDocument(document *gherkin.GherkinDocument) []*Issue {
	return nil
}

func (r *EmptyRule) CheckFeature(feature *gherkin.Feature) []*Issue {
	return nil
}

func (r *EmptyRule) CheckScenario(scenario *gherkin.Scenario) []*Issue {
	return nil
}

func (r *EmptyRule) CheckScenarioOutline(scenarioOutline *gherkin.ScenarioOutline) []*Issue {
	return nil
}

func (r *EmptyRule) CheckSteps(steps []*gherkin.Step) []*Issue {
	return nil
}
