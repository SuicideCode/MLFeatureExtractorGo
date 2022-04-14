package featuremill

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExtractDate(t *testing.T) {

	text := "2018-03-05T03:12:14"

	expected := []string{
		"-627773727:0.166667",
		"210877497:0.