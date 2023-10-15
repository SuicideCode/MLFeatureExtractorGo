
package featuremill

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExtractTimestamp(t *testing.T) {

	text := "2018-03-05T03:12:14"