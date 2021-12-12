
package featuremill

// feature index determinestic by murmur32(uniqueHashPrefixStr+<user field>+<interval>)

import (
	"fmt"

	"github.com/araddon/dateparse"
	"github.com/spaolacci/murmur3"
)

// ExtractDate returns a slice of 2 scaled seasonality vectors: day of week, and month of year
// each with a deterministic feature id
func ExtractDate(field, date string) ([]string, error) {
	var out []string

	dt, err := dateparse.ParseAny(date)
	if err != nil {
		return out, err
	}
