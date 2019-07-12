package convert

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInterface2Map(t *testing.T) {
	m := make(map[string]string, 0)
	m["aa"] = "11"
	m["bb"] = "bb"
	fmt.Println(m)

	var z interface{}
	z = m
	fmt.Println(z)

	m1, _ := Interface2Map(z)
	fmt.Println(m1)

	v1 := reflect.ValueOf(m1)
	fmt.Println("v1:", v1.Kind())

}

func TestBytes2String(t *testing.T) {
	byt := []byte{0x01, 0x02, 0x03}
	s := Bytes2String(byt)
	fmt.Println(s)
}
