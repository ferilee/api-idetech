package middleware

import (
	"net"
	"net/http"
	"strings"
)

func ResolveTenantSlug(r *http.Request) string {
	if headerSlug := strings.TrimSpace(r.Header.Get("X-Tenant-Slug")); headerSlug != "" {
		return sanitizeSlug(headerSlug)
	}

	host := r.Host
	if host == "" {
		return ""
	}

	if strings.Contains(host, ":") {
		var err error
		host, _, err = net.SplitHostPort(host)
		if err != nil {
			host = strings.Split(host, ":")[0]
		}
	}

	parts := strings.Split(host, ".")
	if len(parts) >= 3 {
		return sanitizeSlug(parts[0])
	}

	if len(parts) == 2 && parts[1] == "localhost" {
		return sanitizeSlug(parts[0])
	}

	return ""
}

func sanitizeSlug(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	value = strings.Trim(value, ".")
	return value
}
