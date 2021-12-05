
package featuremill

// feature index determinestic by murmur32(uniqueHashPrefixStr+<user field>+<interval>)

import (
	"fmt"
