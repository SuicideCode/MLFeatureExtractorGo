
package featuremill

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExtractCategorical(t *testing.T) {

	expected := "-1661759407:1"

	got := ExtractCategorical("username or something", "sodunn")
