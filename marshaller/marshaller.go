package marshaller

import (
	"context"
	"encoding/json"
	"reflect"
	"time"
)

func Marshal[T struct{}](ctx context.Context, s *T, fieldTag string, seperator string) string {
	key := ""

	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	for i := 0; i < v.Elem().NumField(); i++ {

		if t.Elem().Field(i).Tag == "" {
			continue
		}

		tag, ok := t.Elem().Field(i).Tag.Lookup(fieldTag)
		if !ok {
			continue
		}
		if v.Elem().Field(i).IsNil() {
			continue
		}

		value := v.Elem().Field(i).Interface()

		if t, ok := value.(time.Time); ok {
			value = t.Unix()
		}
		if t, ok := value.(*time.Time); ok {
			value = t.Unix()
		}

		if key != "" {
			key += seperator
		}

		valueMarshalled, _ := json.Marshal(value)
		key += tag + seperator + string(valueMarshalled)
	}
	return key
}
