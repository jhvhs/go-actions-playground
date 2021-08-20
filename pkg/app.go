package pkg

import (
	"fmt"
	"io"
)

type Application struct {
	Name   string
	Output io.Writer
}

func (a Application) Run() {
	a.Output.Write([]byte(fmt.Sprintf("Hello, %s! Nice to see you.", a.Name)))
}
