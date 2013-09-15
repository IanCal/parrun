package main

import (
	"testing"
        "fmt"
)

func TestOutputWhenSingleDependencyFinished(t *testing.T) {
	dependency := make(chan bool)
	job_run := false
	job := NewJob()
	job.AddDependency(dependency)
	job.SetProcess(func() {
		job_run = true
	})
	if job_run {
		t.Errorf("Job should not run before dependency finishes")
	}
	dependency <- true
	success := <-job.output
	if !success {
		t.Errorf("Job should have returned true")
	}
	if !job_run {
		t.Errorf("Job should actually have been run")
	}

}

func TestOutputWithMultipleDependencies(t *testing.T) {
	dependency1 := make(chan bool)
	dependency2 := make(chan bool)
	dependency3 := make(chan bool)
	job_run := false
	job := NewJob()
	job.AddDependency(dependency1)
	job.AddDependency(dependency2)
	job.AddDependency(dependency3)
	job.SetProcess(func() {
		job_run = true
	})
	if job_run {
		t.Errorf("Job should not run before dependency finishes")
	}
	dependency1 <- true
	dependency2 <- true
	dependency3 <- true
	success := <-job.output
	if !success {
		t.Errorf("Job should have returned true")
	}
	if !job_run {
		t.Errorf("Job should actually have been run")
	}
}

func TestChainedJobs(t *testing.T) {
	dependency1 := make(chan bool)
	dependency2 := make(chan bool)
	dependency3 := make(chan bool)
	job1 := NewJob()
	job1.SetProcess(func() {
            fmt.Printf("Job1\n")
	})
	job2 := NewJob()
	job2.SetProcess(func() {
            fmt.Printf("Job2\n")
	})
	job3 := NewJob()
	job3.SetProcess(func() {
            fmt.Printf("Job3\n")
	})
	job1.AddDependency(dependency1)
	job2.AddDependency(dependency2)
        job2.AddDependency(job1.output)
        job3.AddDependency(job2.output)
        job3.AddDependency(dependency3)
        dependency1 <- true
        dependency2 <- true
        dependency3 <- true
	success := <-job3.output
	if !success {
		t.Errorf("Job should have returned true")
	}
}
