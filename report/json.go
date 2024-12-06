package report

import (
	"bytes"
	"encoding/json"
	"io"
)

func writeJson(findings []Finding, w io.WriteCloser) error {
	if len(findings) == 0 {
		findings = []Finding{}
	}
	for i := range findings {
		// Remove `Line` from JSON output
		findings[i].Line = ""
	}
	return writeJsonExtra(findings, w)
}

func writeJsonExtra(findings []Finding, w io.WriteCloser) error {
	if len(findings) == 0 {
		findings = []Finding{}
	}
	defer w.Close()

	// Convert findings to bytes with CRLF line endings
	data, err := json.MarshalIndent(findings, "", " ")
	if err != nil {
		return err
	}

	// Replace LF with CRLF
	data = bytes.ReplaceAll(data, []byte{'\n'}, []byte{'\r', '\n'})

	// Write with trailing CRLF
	if _, err := w.Write(data); err != nil {
		return err
	}
	_, err = w.Write([]byte{'\r', '\n'})
	return err
}
