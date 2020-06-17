package main

import "strings"

type JsonMap map[string]interface{}

var empty = JsonMap{}

func (j JsonMap) UnpackMap(xpath string) JsonMap {
	m := j
	tokens := strings.Split(xpath, ".")
	for _, token := range tokens {
		m = m.asMap(token)
	}
	return m
}

func (j JsonMap) asMap(key string) JsonMap {
	val := j[key]
	if m, ok := val.(map[string]interface{}); ok {
		return m
	}
	return empty
}

func (j JsonMap) AsString(key string, or string) string {
	val := j[key]
	if m, ok := val.(string); ok {
		return m
	}
	return or
}
