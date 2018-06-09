package main
import (
  "fmt"
  "github.com/user/stringutil"
)

// go tool -> standard way to fetch, build, and install
// Go packages and commands
// go tool builds source packages and installs resulting binaries
// to the pkg and bin directories

// workpsace -> where programmer keeps all their Go code
// Will contain multiple repositories
// Each repo will have one or more packages
// Each package consists of one or more Go source files

// src -- contains Go source files
// pkg -- contains package objects
// bin -- contains executable commands

// echo $(go env GOPATH) --> /Users/sarahheacock/go

// good to create pathname that will not collide with future imports
// build and install program with go tool
// --> `go install github.com/user/hello` (can run command anywhere in system)
// --> `cd $GOPATH/src/github.com/user/hello` && `go install`
// creates hello command producing executable binary
// creates new file called `hello` in `bin`

// build --> just compile file and move to destination
// install --> move executable file to $GOPATH/bin and cache all non-main packages
// which is imported to $GOPATH/pkg
// helps avoid recompiling unnecessarily

// run program by typing full path `$GOPATH/bin/hello`
// OR `hello` since we have added $GOPATH/bin to PATH

func main(){
  //fmt.Printf("Hello, world.\n")
  fmt.Printf(stringutil.Reverse("!oG ,olleH"))
}

// `go install` --> installs package or binary but also whatever dependencies
