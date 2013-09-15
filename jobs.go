package main

type Job struct {
    dependencies [](chan bool)
    output (chan bool)
}

func NewJob() *Job {
    job := Job{
        dependencies: make([](chan bool), 0, 100),
        output: make(chan bool),
    }
    return &job
}

func (job *Job) AddDependency(dependency chan bool) {
    job.dependencies = append(job.dependencies, dependency)
}

func (job *Job) SetProcess(function func()) {
    go func() {
        for i := 0; i < len(job.dependencies); i++ {
            <-job.dependencies[i]
        }
        function()
        job.output <- true
    }()
}

