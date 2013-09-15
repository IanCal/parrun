package main

type Job struct {
	dependencies [](chan bool)
	children     [](chan bool)
}

func NewJob() *Job {
	job := Job{
		dependencies: make([](chan bool), 0, 100),
		children:     make([](chan bool), 0, 100),
	}
	return &job
}

func (job *Job) AddDependency(dependency *Job) {
	c := make(chan bool)
	job.dependencies = append(job.dependencies, c)
	dependency.children = append(dependency.children, c)
}

func (job *Job) AddListener(output chan bool) {
	job.children = append(job.children, output)
}

func (job *Job) SetProcess(function func()) {
	go func() {
		for i := 0; i < len(job.dependencies); i++ {
			<-job.dependencies[i]
		}
		function()
		for i := 0; i < len(job.children); i++ {
			job.children[i] <- true
		}
	}()
}
