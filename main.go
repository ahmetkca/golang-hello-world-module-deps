package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ahmetkca/golang-hello-world-modules/hello"
	"github.com/ahmetkca/golang-hello-world-modules/world"

	"github.com/ahmetkca/golang-hello-world-modules/greeting"
)

func displayGreetings(w io.Writer) {
	fmt.Fprintln(w, hello.Greet())
	fmt.Fprintln(w, world.Greet())
	fmt.Fprintln(w, greeting.SayHi())
	fmt.Fprintln(w, greeting.SayHelloWorld())
}

func main() {
	displayGreetings(os.Stdout)
}
