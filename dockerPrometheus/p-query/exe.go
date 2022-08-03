package main

import (
	"fmt"
	"log"
	"os/exec"
	"os"
	"io/ioutil"
)

func main() {
	log.Print("Enter exe.go")	
	log.Print(os.Getwd())
	files, errr := ioutil.ReadDir("./")
    if errr != nil {
        log.Fatal(errr)
    }

    for _, f := range files {
            log.Print(f.Name())
    }

	ts_cmd := exec.Command("tsc", "instant_query.ts")
	err := ts_cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		log.Print(err.Error())
		return
	}

	js_cmd := exec.Command("node", "instant_query.js")

	stdout, err := js_cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		log.Print(err.Error())
		return
	}
	fmt.Print("Here is the output")
	fmt.Print(string(stdout))
}
