package unexportedglobal

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	t.Parallel()

	analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), Analyzer, "./...")
}
