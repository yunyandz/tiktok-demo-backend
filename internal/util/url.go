package util

import (
	"net/url"
	"strings"
)

func GetRawUrl(url *url.URL) string {
	return strings.Join([]string{url.Scheme, "://", url.Host, url.Path}, "")
}
