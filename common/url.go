package common

import (
	"fmt"
	"strings"
)

func ResolveURL(url string) string {
	if strings.Contains(url, "http://") {
		return url
	}
	return fmt.Sprintf("%s%s", "http://", url)
}
