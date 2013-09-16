package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func create_runner(command_line, outputname string) func() bool {
	return func() bool {
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
		io.Copy(ferr, stderr)
		if err := cmd.Wait(); err != nil {
			log.Fatalf("Error running command '%s' for job '%s', please see the error logs", command_line, outputname)
			return false
		}
		return true
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

func print_dotfile(conf Configuration, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(file)
	defer func() {
		w.Flush()
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	fmt.Fprintln(w, "digraph dependencies {")
	for name, jobConf := range conf.Jobdescs {
		for _, dependency := range jobConf.Dependencies {
			fmt.Fprintln(w, "    "+dependency+" -> "+name+";")
		}
	}
	fmt.Fprintln(w, "}")
}

func main() {
	var configFile = flag.String("config", "run.toml", "The location of the configuration file")
	var dryrun = flag.Bool("dryrun", false, "Don't actually run the jobs")
	var dotfile = flag.String("dot", "diagram.dot", "Print out a dotfile")
	flag.Parse()
	conf := LoadConfig(*configFile)
	print_dotfile(conf, *dotfile)
	if !(*dryrun) {
		totalJobs, finaliser := create_workflow(conf)
		for i := 0; i < totalJobs; i++ {
			<-finaliser
			log.Printf("Job %d/%d finished", i+1, totalJobs)
		}
	}
}
