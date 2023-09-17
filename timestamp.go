
package featuremill

// feature index determinestic by murmur32(uniqueHashPrefixStr+<user field>+<interval>)

import (
	"fmt"

	"github.com/araddon/dateparse"
	"github.com/spaolacci/murmur3"
)

// ExtractTimestamp returns a slice of 3 scaled seasonality vectors: minute of hour, hour of day, day of week
// each with a deterministic feature ID
func ExtractTimestamp(field, timestamp string) ([]string, error) {
	var out []string

	dt, err := dateparse.ParseAny(timestamp)
	if err != nil {
		return out, err
	}