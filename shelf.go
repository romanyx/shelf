package shelf

import (
	"fmt"
	"reflect"
	"sync"
)

var (
	shelf sync.Map
)

type dep[T any] struct {
	s *stack
	d T
}

// Put allows to put dependency on shelf with specific name
// or using type as a key.
func Put[T any](d T, name ...string) {
	key := getKey[T](name...)
	shelf.Store(key, dep[T]{
		d: d,
		s: callers(),
	})
}

// Take allows to take dependency from shelf by
// type or specific name.
func Take[T any](name ...string) interface{} {
	key := getKey[T](name...)
	d, ok := shelf.Load(key)
	if !ok {
		panic(fmt.Sprintf("dependency %s does not exists", key))
	}

	dep := d.(dep[T]) // nolint: errcheck

	if isNil(dep.d) {
		m := fmt.Sprintf("dependency %s holds nil value\n", key)
		for _, frame := range dep.s.StackTrace() {
			m += fmt.Sprintf(
				"%s:%d\n",
				frame.File,
				frame.Line,
			)
		}
		panic(m)
	}

	return dep.d
}

func getKey[T any](name ...string) string {
	if len(name) > 1 {
		panic("only one key allowed")
	}

	var key string
	if len(name) == 1 {
		key = name[0]
	}
	if key == "" {
		v := new(T)
		key = fmt.Sprintf("%s", reflect.TypeOf(*v))
	}

	return key
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr,
		reflect.Map,
		reflect.Array,
		reflect.Chan,
		reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
