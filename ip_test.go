
package featuremill

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExtractIP(t *testing.T) {

	text := "127.0.0.1"

	expected := "1799088460:2130706433"
