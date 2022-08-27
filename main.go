package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"os/exec"
)

func read_pipeline() map[interface{}]interface{} {
	yfile, err := os.ReadFile("pipeline")

	if err != nil {
		log.Fatal(err)
	}

	data := make(map[interface{}]interface{})
	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	return data
}

func main() {
	//cmd := "build -t test:latest ."
	//args := strings.Fields(cmd)
	//args := []string{"build", "-t", "test:latest", "."}

	data := read_pipeline()

	for k, v := range data {
		fmt.Printf("%s -> %s\n", k, v)
	}

	command := "version"

	out, err := exec.Command("/usr/bin/docker", command).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}
