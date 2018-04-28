package main

import "testing"

func TestStartJvm(t *testing.T)  {
	cmd := &Cmd{}
	cmd.class = "java.lang.Byte"
	cmd.cpOption = "."
	startJVM(cmd)
}
