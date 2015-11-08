package component
import (
    "reflect"
    "gogy/err"
)

type Reader struct {
    Data map[string]interface{}
}

func (obj *Reader) ReadString(key string) (string, error) {
    val := obj.Read(key)

    if (val == nil) {
        return "", &err.MandatoryArgumentError{key}
    }

    t := reflect.TypeOf(val).String()

    if (t != "string") {
        return "", &err.NotStringError{key};
    } else {
        return val.(string), nil;
    }
}

func (obj *Reader) Read(key string) interface{} {
    if obj.Has(key) {
        return obj.Data[key]
    }
    return nil;
}

func (obj *Reader) Has(key string) bool {
    _, ok := obj.Data[key]
    return ok
}