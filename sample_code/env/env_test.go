package env

import "testing"

// assume ProcessEnvVars is a function that processes environment variables
// and returns a struct with an OutputFormat field
func TestEnvVarProcess(t *testing.T) {
	t.Setenv("OUTPUT_FORMAT", "JSON")
	cfg := ProcessEnvVars()
	if cfg.OutputFormat != "JSON" {
		t.Error("OutputFormat not set correctly")
	}
	// value of OUTPUT_FORMAT is reset when the test function exits
}
