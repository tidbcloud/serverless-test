package imp

import (
	"os"
	"testing"
)

func setup() {
	println("setup")
}

func TestImport(t *testing.T) {
	println("TestImport")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
