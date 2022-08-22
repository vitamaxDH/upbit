package util

import (
	"net/url"
)

func AddQueryIfNotEmpty(queries *url.Values, queryStr, query string) {
	if query != "" {
		queries.Add(queryStr, query)
	}
}
