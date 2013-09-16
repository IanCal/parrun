package main

import (
	"testing"
)

func TestJobWithNoDependenciesRuns(t *testing.T) {
	success := make(chan bool)
	job_run := false
	job := NewJob()
	job.AddListener(success)
	job.SetProcess(func() bool {
		job_run = true
		return true
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
	job1.SetProcess(func() bool {
		job1_run = true
		return true
	})
	job2.SetProcess(func() bool {
		job2_run = true
		return true
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
	job1.SetProcess(func() bool { return true })
	job2.SetProcess(func() bool { return true })
	job3.SetProcess(func() bool { return true })
	if !(<-success) {
		t.Errorf("Job should have returned true")
	}
}

func TestJobsAfterFailingJobDontRun(t *testing.T) {
	success := make(chan bool)
	job1run := false
	job2run := false
	job3run := false
	job1 := NewJob()
	job2 := NewJob()
	job3 := NewJob()
	job2.AddDependency(job1)
	job3.AddDependency(job2)
	job3.AddListener(success)
	job1.SetProcess(func() bool {
		job1run = true
		return true
	})
	job2.SetProcess(func() bool {
		job2run = true
		return false
	})
	job3.SetProcess(func() bool {
		job3run = true
		return true
	})
	if <-success {
		t.Errorf("Job should have returned false")
	}
	if !job1run {
		t.Errorf("Job 1 should have run")
	}
	if !job2run {
		t.Errorf("Job 2 should have run")
	}
	if job3run {
		t.Errorf("Job 3 should not have run")
	}
}
