package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func create_runner(cmd *exec.Cmd, outputname string) func() {
	return func() {
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
}

func sleeper(time int, name string) func() {
	cmd := exec.Command("sleep", strconv.Itoa(time))
	return create_runner(cmd, name)
}

func main() {
	totalJobs := 3
	c := make(chan bool)
	j1 := NewJob()
	j2 := NewJob()
	j3 := NewJob()
	j2.AddDependency(j1)
	j3.AddDependency(j1)
	j1.AddListener(c)
	j2.AddListener(c)
	j3.AddListener(c)
	j1.SetProcess(sleeper(5, "uno"))
	j2.SetProcess(sleeper(5, "dos"))
	j3.SetProcess(sleeper(5, "tres"))
	for i := 0; i < totalJobs; i++ {
		<-c
		log.Printf("Job %d/%d finished", i+1, totalJobs)
	}
}
