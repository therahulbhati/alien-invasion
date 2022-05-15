package cmd

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Execute(t *testing.T) {
	actual := captureOutput(func() { Execute("../test/testData/sample.txt", 42, 5, 10) })

	expected := `Alien 0 placed in city Bar
Alien 1 placed in city Bee
Alien 2 placed in city Foo
Alien 3 placed in city Bar
Alien 4 placed in city Foo
Iteration 0
Bar has been destroyed by aliens 0, 3!
Foo has been destroyed by aliens 2, 4!
Iteration 1
Alien 1 trapped in city: Bee

Baz
Bee
Qu-ux

`
	assert.Equal(t, expected, actual)

}

func captureOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	return string(out)
}
