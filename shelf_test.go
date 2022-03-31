package shelf

import (
	"fmt"
	"strings"
	"testing"
)

type dependency struct {
	check string
}

func TestTake(t *testing.T) {
	tt := []struct {
		name  string
		dep   dependency
		key   []string
		panic bool
	}{
		{
			name: "without key",
			dep:  dependency{"dep"},
		},
		{
			name: "with key",
			dep:  dependency{"dep"},
			key:  []string{"key"},
		},
		{
			name:  "uknown key",
			dep:   dependency{"dep"},
			panic: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panic {
				defer func() {
					if r := recover(); r != nil {
						if fmt.Sprint(r) != fmt.Sprintf("dependency %s does not exists", tc.key) {
							t.Errorf("should return correct panic")
						}
					}
				}()

				Put[*dependency](&tc.dep, tc.key...)
				if &tc.dep != Take[*dependency](tc.key...) {
					t.Errorf("should return correct dep from shelf")
				}
			}
		})
	}
}

func TestTakeNil(t *testing.T) {
	Put[*dependency](nil)

	defer func() {
		if r := recover(); r != nil {
			if !strings.Contains(fmt.Sprint(r), "dependency *shelf.dependency holds nil value") {
				t.Errorf("should return correct panic")
			}
		}
	}()

	Take[*dependency]()
}

func TestTakeKey(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if !strings.Contains(fmt.Sprint(r), "only one key allowed") {
				t.Errorf("should return correct panic")
			}
		}
	}()

	Put[*dependency](nil, "1", "2")
}
