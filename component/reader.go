package component

import ()

type Reader map[string]interface{}

func (obj Reader) ReadString(key string) string {
	val := obj.Read(key)

	if s, ok := val.(string); ok {
		return s
	} else {
		return ""
	}
}

func (obj Reader) Read(key string) interface{} {
	if obj.Has(key) {
		return obj[key]
	}
	return nil
}

func (obj Reader) Has(key string) bool {
	_, ok := obj[key]
	return ok
}
