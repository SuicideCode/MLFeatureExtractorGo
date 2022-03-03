package featuremill

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExtractDate(t *testing.T) {

	text := "2018-03-0