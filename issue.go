package lingo

import (
	"fmt"
	"strings"
)

// DefaultIssueFormat used to print an issue
const DefaultIssueFormat = "{{.Path}}:{{.Line}}:{{if .Col}}{{.Col}}{{end}}:{{.Severity}}: {{.Message}} ({{.Linter}})"

// Severity of linter message
type Severity string

// Linter message severity levels.
const (
	Error   Severity = "error"
	Warning Severity = "warning"
)

type Issue struct {
	Linter   string   `json:"linter"`
	Severity Severity `json:"severity"`
	Path     string   `json:"path"`
	Line     int      `json:"line"`
	Col      int      `json:"col"`
	Message  string   `json:"message"`
}

func (i *Issue) String() string {
	col := ""
	if i.Col != 0 {
		col = fmt.Sprintf("%d", i.Col)
	}
	return fmt.Sprintf("%s:%d:%s:%s: %s (%s)", strings.TrimSpace(i.Path), i.Line, col, i.Severity, strings.TrimSpace(i.Message), i.Linter)
}

type sortedIssues struct {
	issues []*Issue
	order  []string
}

func (s *sortedIssues) Len() int      { return len(s.issues) }
func (s *sortedIssues) Swap(i, j int) { s.issues[i], s.issues[j] = s.issues[j], s.issues[i] }

func (s *sortedIssues) Less(i, j int) bool {
	l, r := s.issues[i], s.issues[j]
	return CompareIssue(*l, *r, s.order)
}

// CompareIssue two Issues and return true if left should sort before right
// nolint: gocyclo
func CompareIssue(l, r Issue, order []string) bool {
	for _, key := range order {
		switch {
		case key == "path" && l.Path != r.Path:
			return l.Path < r.Path
		case key == "line" && l.Line != r.Line:
			return l.Line < r.Line
		case key == "column" && l.Col != r.Col:
			return l.Col < r.Col
		case key == "severity" && l.Severity != r.Severity:
			return l.Severity < r.Severity
		case key == "message" && l.Message != r.Message:
			return l.Message < r.Message
		case key == "linter" && l.Linter != r.Linter:
			return l.Linter < r.Linter
		}
	}
	return true
}
