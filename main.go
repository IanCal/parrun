package main

import (
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
)

func create_runner(command_line, outputname string) func() {
	return func() {
		cmd := exec.Command("sh", "-c", command_line)
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

func create_workflow(conf Configuration) (int, chan bool) {
	finaliser := make(chan bool)
	// Create jobs
	jobs := map[string]*Job{}
	for name, _ := range conf.Jobdescs {
		jobs[name] = NewJob()
	}
	// Set dependencies
	for name, jobConf := range conf.Jobdescs {
		for _, dependency := range jobConf.Dependencies {
			jobs[name].AddDependency(jobs[dependency])
		}
		jobs[name].AddListener(finaliser)
	}
	// Set commands
	for name, jobConf := range conf.Jobdescs {
		jobs[name].SetProcess(create_runner(jobConf.Command, name))
	}
	return len(jobs), finaliser
}

func main() {
	var configFile = flag.String("config", "run.toml", "The location of the configuration file")
	flag.Parse()
	conf := LoadConfig(*configFile)
	totalJobs, finaliser := create_workflow(conf)
	for i := 0; i < totalJobs; i++ {
		<-finaliser
		log.Printf("Job %d/%d finished", i+1, totalJobs)
	}
}
