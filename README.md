# Gherkin Linter written in Go
[![Build Status](https://travis-ci.org/remast/gherkin-lingo.svg?branch=master)](https://travis-ci.org/remast/gherkin-lingo)

Linter for Gherkin files written in Go. Output is written in a standard format:

    <file>:<line>:[<column>]: <message> (<linter>)

eg.

    stutter.go:9::warning: unused global variable unusedGlobal (varcheck)
    stutter.go:12:6:warning: exported type MyStruct should have comment or be unexported (golint)


## Checks

missing-feature-name Feature name is empty or missing.
missing-scenario-name Scenario name is empty or missing.
given-in-first-line Given must occurr in first line and in first line only.
outline-with-too-few-examples Scenario Outline with no more than one example

### Ideas

    no_repeating_keywords
    no_empty_features
    avoid outline for single example
    avoid period
    avoid scripting
    be declarative
    background does more than setup
    background requires scenario
    bad scenario name
    file name differs feature name
    invalid file name
    invalid step flow
    missing example name
    missing feature description
    missing test action
    missing verification
    required tag starts with - disabled by default
    same tag for all scenarios
    tag used multiple times
    too clumsy
    too long step
    too many different tags
    too many steps
    too many tags
    unique scenario names
    unknown variable
    use background
    use outline
no-tags-on-backgrounds * 	Disallows tags on Background
one-feature-per-file * 	Disallows multiple Feature definitions in the same file
up-to-one-background-per-file * 	Disallows multiple Background definition in the same file
indentation 	Allows the user to specify indentation rules
name-length 	Allows restricting length of Feature/Scenario/Step names
new-line-at-eof 	Disallows/enforces new line at EOF
no-dupe-feature-names 	Disallows duplicate Feature names
no-dupe-scenario-names 	Disallows duplicate Scenario names
no-duplicate-tags 	Disallows duplicate tags on the same Feature or Scenario
no-empty-file 	Disallows empty feature files
no-files-without-scenarios 	Disallows files with no scenarios
no-homogenous-tags 	Disallows tags present on every Scenario in a Feature, rather than on the Feature itself
no-multiple-empty-lines 	Disallows multiple empty lines
no-partially-commented-tag-lines 	Disallows partially commented tag lines
no-restricted-tags 	Disallow use of particular @tags
no-scenario-outlines-without-examples 	Disallows scenario outlines without examples
no-superfluous-tags 	Disallows tags present on a Feature and a Scenario in that Feature
no-trailing-spaces 	Disallows trailing spaces
use-and 	Disallows repeated step names requiring use of And instead


## Credits

This is heavily inspired and copied from gometalinter.