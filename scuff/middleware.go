package scuff

import "log"

func (scuff *scuff) Use(data map[string]interface{}) {
	mergeMaps(scuff.AsMap, data)
}

func mergeMaps(into map[string]interface{},from map[string]interface{}) {
	for key, v := range from {
		switch mergeMe := v.(type) {
		case map[string]interface{}:
			if intoChild, ok := into[key].(map[string]interface{}); ok {
				mergeMaps(intoChild, mergeMe)
			}
			log.Panic("incompatible types to merge at key", key)
		}
	}
}