
package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/dustin-decker/featuremill"
)

func main() {

	var features []string

	ip, _ := featuremill.ExtractIP("src_ip", "127.0.0.1")
	features = append(features, ip)

	user := featuremill.ExtractCategorical("username", "emily")
	features = append(features, user)

	desc := featuremill.ExtractText("user terminated the session", " ")