
package featuremill

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExtractBoolean(t *testing.T) {

	text := "no"
	expected := "-1235294128:0"
	got, _ := ExtractBoolean("IsBad", text)