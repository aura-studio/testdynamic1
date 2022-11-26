package testdynamic1

import (
	"fmt"

	"github.com/aura-studio/dynamic"
	"github.com/aura-studio/dynamic/test/testmod"
)

type TestDynamic1 struct {
}

func (*TestDynamic1) Init() {
}

func (*TestDynamic1) Invoke(string, string) string {
	fmt.Println(testmod.Display())
	return ""
}

func (*TestDynamic1) Close() {
}

var Tunnel dynamic.Tunnel = &TestDynamic1{}
