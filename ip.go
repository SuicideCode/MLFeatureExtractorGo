
package featuremill

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"

	"github.com/spaolacci/murmur3"
)

// ExtractIP returns a vector that is a scaled integer representation of IPv4 and IPv6 IPs
// with a deterministic feature ID
func ExtractIP(field, addr string) (string, error) {