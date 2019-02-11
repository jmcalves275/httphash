package common

import (
	"fmt"
	"strings"
)

// Function to resolve the url
//
// Example:
// google.com turns into http://google.com
func ResolveURL(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	}
	return fmt.Sprintf("%s%s", "http://", url)
}
