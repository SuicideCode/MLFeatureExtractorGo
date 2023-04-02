
package featuremill

import (
	"fmt"
	"math"

	"github.com/spaolacci/murmur3"
)

// ExtractNumericalMaxMin returns a min/max scalled vector to a deterministic feature ID
func ExtractNumericalMaxMin(field string, num float32, min float32, max float32) string {
	featureID := int32(murmur3.Sum32([]byte(uniqueHashPrefixStr + field)))
	return fmt.Sprintf("%d:%f", featureID, (num-min)/(max-min))
}

// ExtractNumerical returns a min/max scalled vector to a deterministic feature ID
func ExtractNumerical(field string, num float32, min float32, max float32) string {