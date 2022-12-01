package testdynamic1

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/aura-studio/dynamic"
	"github.com/aura-studio/dynamic/test/testmod"
)

type TestDynamic1 struct {
}

func (*TestDynamic1) Init() {
}

func (*TestDynamic1) Invoke(string, string) string {
	fmt.Println(testmod.Display())
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	return string(data)
}

func (*TestDynamic1) Close() {
}

var Tunnel dynamic.Tunnel = &TestDynamic1{}
