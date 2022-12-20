package urlsparsing

import (
	"net/url"
	"strings"
)

type ParsedURL struct {
	Scheme string
	Host   string
	Path   string
	Query  map[string]string
}

func Parse(rawUrl string) ParsedURL {
	result, _ := url.Parse(rawUrl)

	parsedURL := ParsedURL{
		Scheme: result.Scheme,
		Host:   result.Host,
		Path:   result.Path,
		Query:  map[string]string{},
	}

	for key, value := range result.Query() {
		parsedURL.Query[key] = value[0]
	}

	return parsedURL
}

func Construct(parsedURL ParsedURL) string {
	result := &url.URL{
		Scheme: parsedURL.Scheme,
		Host:   parsedURL.Host,
		Path:   parsedURL.Path,
	}

	rawQuery := ""
	for key, value := range parsedURL.Query {
		rawQuery = rawQuery + key + "=" + value + "&"
	}

	result.RawQuery = strings.TrimRight(rawQuery, "&")

	return result.String()
}
