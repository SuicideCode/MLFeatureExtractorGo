
package featuremill

import (
	"fmt"
	"strings"

	"github.com/spaolacci/murmur3"
)

// ExtractText returns a slice of "featureID:1" strings for each token in the string
func ExtractText(text, delim string) []string {