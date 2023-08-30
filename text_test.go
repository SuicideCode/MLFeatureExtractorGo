
package featuremill

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExtractText(t *testing.T) {

	text := "testing some shit"

	expected := []string{
		"164984899:1",
		"650894796:1",