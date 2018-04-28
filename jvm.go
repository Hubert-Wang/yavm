package main

import (
	"fmt"
	"strings"
	"yavm/classpath"
)

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)

	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)

	classData, _, err := cp.ReadClass(className)

	if err != nil {
		fmt.Printf("Could not find or load main class %s\n ", cmd.class)
	}

	fmt.Printf("class data:%v\n", classData)
}

