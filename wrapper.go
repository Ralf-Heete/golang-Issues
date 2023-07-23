package main

import (
	"fmt"
)

type value interface{}
type dict map[string]value
type list []value

func (recv *dict) SetItem(k value, v value) bool {
	key, ok := k.(string)
	if ok {
		(*recv)[key] = v
		return true
	}
	return false
}

func (recv *list) SetItem(k value, v value) bool {
	i, ok := k.(int)
	if !ok || len(*recv) <= i {
		return false
	}

	(*recv)[i] = v
	return true
}

// Wrapper function that calls SetItem and returns the result.
func SetItemWrapper(d value, k value, v value) bool {
	switch t := d.(type) {
	case *dict:
		return t.SetItem(k, v)
	case *list:
		return t.SetItem(k, v)
	default:
		return false
	}
}

func main() {
	d := dict{}
	l := list{"1", "2"}
	v := value(&d)

	fmt.Printf("dict:%v\n", SetItemWrapper(v, "foo", "bar"))
	fmt.Printf("d:'%v'\n", d)

	v = value(&l)
	fmt.Printf("list:%v\n", SetItemWrapper(v, 1, "bar"))
	fmt.Printf("l:'%v'\n", l)
}
