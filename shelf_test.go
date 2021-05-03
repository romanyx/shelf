package shelf

import (
	"fmt"
	"strings"
	"testing"
)

func TestTake(t *testing.T) {
	type dependency struct {
		check string
	}

	dep := dependency{"dep"}

	Put("dep", &dep)

	got := Take("dep")

	if &dep != got {
		t.Errorf("should return correct dep from shelf")
	}

	defer func() {
		if r := recover(); r != nil {
			if fmt.Sprint(r) != "dependency unknown does not exists" {
				t.Errorf("should return correct panic")
			}
		}
	}()

	Take("unknown")
}

func TestTakeNil(t *testing.T) {
	Put("dep", nil)

	defer func() {
		if r := recover(); r != nil {
			if !strings.Contains(fmt.Sprint(r), "dependency dep holds nil value") {
				t.Errorf("should return correct panic")
			}
		}
	}()

	Take("dep")
}
