package report

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var simpleFinding = Finding{
	Description: "",
	RuleID:      "test-rule",
	Match:       "line containing secret",
	Line:        "whole line containing secret",
	Secret:      "a secret",
	StartLine:   1,
	EndLine:     2,
	StartColumn: 1,
	EndColumn:   2,
	Message:     "opps",
	File:        "auth.py",
	SymlinkFile: "",
	Commit:      "0000000000000000",
	Author:      "John Doe",
	Email:       "johndoe@gmail.com",
	Date:        "10-19-2003",
	Tags:        []string{},
}

func TestWriteJSON(t *testing.T) {
	tests := []struct {
		findings       []Finding
		testReportName string
		expected       string
		wantEmpty      bool
	}{
		{
			testReportName: "simple",
			expected:       filepath.Join(expectPath, "report", "json_simple.json"),
			findings: []Finding{
				simpleFinding,
			}},
		{

			testReportName: "empty",
			expected:       filepath.Join(expectPath, "report", "empty.json"),
			findings:       []Finding{}},
		{
			testReportName: "with_quotes",
			expected:       filepath.Join(expectPath, "report", "json_with_quotes.json"),
			findings: []Finding{
				{
					Description: "",
					RuleID:      "test-rule",
					Match:       `line containing "quoted" secret`,
					Line:        "whole line containing secret",
					Secret:      `a "quoted" secret`,
					StartLine:   1,
					EndLine:     2,
					StartColumn: 1,
					EndColumn:   2,
					Message:     `oops with "quotes"`,
					File:        "auth.py",
					SymlinkFile: "",
					Commit:      "0000000000000000",
					Author:      "John Doe",
					Email:       "johndoe@gmail.com",
					Date:        "10-19-2003",
					Tags:        []string{},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testReportName, func(t *testing.T) {
			tmpfile, err := os.Create(filepath.Join(t.TempDir(), test.testReportName+".json"))
			require.NoError(t, err)
			err = writeJson(test.findings, tmpfile)
			require.NoError(t, err)
			assert.FileExists(t, tmpfile.Name())
			got, err := os.ReadFile(tmpfile.Name())
			require.NoError(t, err)
			if test.wantEmpty {
				assert.Empty(t, got)
				return
			}
			want, err := os.ReadFile(test.expected)
			require.NoError(t, err)
			// require.JSONEq(t, string(want), string(got))
			// fmt.Println(string(got) == string(want))
			// fmt.Println(string(got))
			// fmt.Println(string(want))
			assert.Equal(t, want, got)
		})
	}
}

func TestWriteJSONExtra(t *testing.T) {
	findings := []Finding{
		simpleFinding,
	}
	expected := filepath.Join(expectPath, "report", "json_extra_simple.json")

	tmpfile, err := os.Create(filepath.Join(t.TempDir(), "simple_extra.json"))
	require.NoError(t, err)

	err = writeJsonExtra(findings, tmpfile)
	require.NoError(t, err)
	assert.FileExists(t, tmpfile.Name())

	got, err := os.ReadFile(tmpfile.Name())
	require.NoError(t, err)
	want, err := os.ReadFile(expected)
	require.NoError(t, err)
	assert.Equal(t, want, got)
}
