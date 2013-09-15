package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func run(cmd *exec.Cmd, outputname string) {
	fo, err := os.Create("/tmp/" + outputname + ".out")
	ferr, err := os.Create("/tmp/" + outputname + ".err")
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
		if err := ferr.Close(); err != nil {
			panic(err)
		}
	}()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	io.Copy(fo, stdout)
	io.Copy(fo, stderr)
	err = cmd.Wait()
}

func sleep(c chan<- string, fname string) {
	cmd := exec.Command("echo", "hello")
	run(cmd, fname)
	c <- "I'm done"
}

func main() {
	c := make(chan string)
	for i := 0; i < 10; i++ {
		go sleep(c, fmt.Sprintf("%d", i))
	}
	for i := 0; i < 10; i++ {
		result := <-c
		log.Printf("Received %v", result)
	}
}
