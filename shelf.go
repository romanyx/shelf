package shelf

import (
	"fmt"
	"reflect"
	"sync"
)

var (
	shelf sync.Map
)

type dep struct {
	s *stack
	d interface{}
}

// Put allows to put dependency on shelf
func Put(name string, d interface{}) {
	shelf.Store(name, dep{
		d: d,
		s: callers(),
	})
}

// Take allows to take dependency from shelf.
func Take(name string) interface{} {
	d, ok := shelf.Load(name)
	if !ok {
		panic(fmt.Sprintf("dependency %s does not exists", name))
	}

	dep := d.(dep) // nolint: errcheck

	if isNil(dep.d) {
		m := fmt.Sprintf("dependency %s holds nil value\n", name)
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
