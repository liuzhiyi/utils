package str

import (
	"fmt"
	"testing"
)

type item struct {
	Id    []byte
	Value string
}

type key struct {
	K string
}

func Test_Serialize(t *testing.T) {
	k, _ := Serialize(key{"test"})
	i, err := Serialize(item{(k), "forvet"})
	fmt.Println(string(i), err)
	var dei item
	Deserialize(i, &dei)
	fmt.Println(dei)
}
