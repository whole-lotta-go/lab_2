package lab2

import (
	"bytes"
	"strings"
	"testing"
)

func TestCompute(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantOutput string
		wantError  bool
	}{
		{"successful evaluation", "2 3 +", "5\n", false},
		{"syntax error in input", "2 +", "", true},
		{"empty input", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			writer := &bytes.Buffer{}

			handler := &ComputeHandler{
				Reader: reader,
				Writer: writer,
			}

			err := handler.Compute()

			if (err != nil) != tt.wantError {
				t.Errorf("Compute(\"%s\") error = %v, wantError = %v", tt.input, err, tt.wantError)
				return
			}

			if gotOutput := writer.String(); gotOutput != tt.wantOutput {
				t.Errorf("Compute(\"%s\") output = %q, wantOutput = %q", tt.input, gotOutput, tt.wantOutput)
			}
		})
	}
}
