
package featuremill

import (
	"errors"
	"fmt"

	"github.com/spaolacci/murmur3"
)

// ExtractBoolean returns a 0/1 vector for the deterministic feature ID
func ExtractBoolean(field, boolean string) (string, error) {

	// deterministic collision resistant feature id
	featureID := int32(murmur3.Sum32([]byte(uniqueHashPrefixStr + field)))

	if boolean == "false" ||
		boolean == "False" ||