package config

import "os"

func GetHttpCookiePath() string {
	return os.Getenv("HTTP_COOKIE_PATH")
}

func GetHTTPCookiePathWithDefault(priorityDefault string) string {
	path := GetHttpCookiePath()
	if path == "" {
		return priorityDefault
	}
	return path
}