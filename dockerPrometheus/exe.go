package main

import (
	"fmt"
	"os/exec"
)

func main() {

	ts_cmd := exec.Command("tsc", "instant_query.ts")
	err := ts_cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	js_cmd := exec.Command("node", "instant_query.js")

	stdout, err := js_cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
