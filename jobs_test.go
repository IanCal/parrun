package main

import (
	"testing"
)

func TestJobWithNoDependenciesRuns(t *testing.T) {
	success := make(chan bool)
	job_run := false
	job := NewJob()
	job.AddListener(success)
	job.SetProcess(func() {
		job_run = true
	})
	if !(<-success) {
		t.Errorf("Job should have returned true")
	}
	if !job_run {
		t.Errorf("Job should actually have been run")
	}

}

func TestOutputWithSingleDependency(t *testing.T) {
	success := make(chan bool)
	job1_run := false
	job2_run := false
	job1 := NewJob()
	job2 := NewJob()
	job2.AddDependency(job1)
	job2.AddListener(success)
	job1.SetProcess(func() {
		job1_run = true
	})
	job2.SetProcess(func() {
		job2_run = true
	})
	if !(<-success) {
		t.Errorf("Job should have returned true")
	}
	if !(job1_run && job2_run) {
		t.Errorf("Job should actually have been run")
	}
}

func TestChainedJobs(t *testing.T) {
	success := make(chan bool)
	job1 := NewJob()
	job2 := NewJob()
	job3 := NewJob()
	job2.AddDependency(job1)
	job3.AddDependency(job2)
	job3.AddDependency(job1)
        job3.AddListener(success)
	job1.SetProcess(func() {})
	job2.SetProcess(func() {})
	job3.SetProcess(func() {})
	if !(<-success) {
		t.Errorf("Job should have returned true")
	}
}
