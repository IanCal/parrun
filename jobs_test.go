package main

import (
	"testing"
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
