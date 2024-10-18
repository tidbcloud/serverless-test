package imp

import (
	"flag"
	"os"
	"testing"
)

var IsOpenAPI string

func initFlag() {
	flag.StringVar(&IsOpenAPI, "IsOpenAPI", "N", "")
	flag.Parse()
	println(IsOpenAPI)
}

func setup() {
	println("setup")
}

func TestCreateExport(t *testing.T) {
	println("CreateExport")
}

func TestCreateExport1(t *testing.T) {
	println("CreateExport1")
}

func TestMain(m *testing.M) {
	initFlag()
	setup()
	code := m.Run()
	os.Exit(code)
}
