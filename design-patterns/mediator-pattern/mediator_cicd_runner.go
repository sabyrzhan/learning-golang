package mediator_pattern

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Here I am simulating CICD server with runner and jobs, much like GitLab CICD
// CICD server acts as a mediator between runners and jobs.
// Runners are registered to server. And runners submit the jobs to server.
// Server executes the jobs and set job results to runner.
// Runners regularly can pull the results from server.

type CICDServer struct {
	Runners map[string]*CICDRunner
	runnersMutex sync.RWMutex
	JobResults map[string][]string
	jobResultsMutex sync.RWMutex
}

func (s *CICDServer) RegisterRunner(runner *CICDRunner) {
	s.runnersMutex.Lock()
	defer s.runnersMutex.Unlock()
	s.Runners[runner.RunnerId] = runner
	fmt.Println(fmt.Sprintf("[SERVER][INFO] Registered new runner=%s", runner.RunnerId))
}

func (s *CICDServer) SubmitJobs(runnerId string, jobs []string) {
	for _, jobScript := range jobs {
		jobId := fmt.Sprintf("%d", 1000000 + rand.Intn(10000000))
		job := CICDRunnerJob{
			JobId: jobId,
			Server: s,
			RunnerId: runnerId,
			Name: "JOB_" + jobId,
		}
		job.Execute(jobScript)
		time.Sleep(1 * time.Second)
	}

	s.runnersMutex.Lock()
	defer s.runnersMutex.Unlock()
	runner := s.Runners[runnerId]
	if runner != nil {
		message := "Finished all jobs"
		if runner.FinishMessage != "" {
			message = runner.FinishMessage
		}

		fmt.Println(fmt.Sprintf("[SERVER][INFO] Runner=%s: %s", runner.RunnerId, message))
	} else {
		fmt.Println(fmt.Sprintf("[SERVER][ERROR] Runner not found to show finish message!"))
	}
}

func (s *CICDServer) ReceiveJobResult(job *CICDRunnerJob, result string) {
	s.jobResultsMutex.Lock()
	s.JobResults[job.RunnerId] = append(s.JobResults[job.RunnerId], result)
	fmt.Println(fmt.Sprintf("[SERVER][INFO] Job %s finished job", job.JobId))
	s.jobResultsMutex.Unlock()
}

func (s *CICDServer) GetRunnerResults(runnerId string) []string {
	s.jobResultsMutex.Lock()
	defer s.jobResultsMutex.Unlock()
	if len(s.JobResults[runnerId]) != 0 {
		result := s.JobResults[runnerId]
		s.JobResults[runnerId] = []string{}
		return result
	}

	return []string{}
}

type CICDRunnerJob struct {
	JobId string
	Server *CICDServer
	RunnerId string
	Name string
}

func (j *CICDRunnerJob) Execute(script string) {
	fmt.Println(fmt.Sprintf("[JOB][INFO] Job=%s: exec: %s", j.JobId, script))
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	j.Server.ReceiveJobResult(j, fmt.Sprintf("[RUNNER=%s][JOB=%s][INFO][Success] %s", j.RunnerId, j.JobId, script))
}


type CICDRunner struct {
	Server   *CICDServer
	RunnerId string
	Scripts []string
	FinishMessage string
}

func (r *CICDRunner) AddScript(script string) {
	r.Scripts = append(r.Scripts, script)
}

func (r *CICDRunner) SetFinishMessage(message string) {
	r.FinishMessage = message
}

func (r *CICDRunner) Run() {
	r.Server.SubmitJobs(r.RunnerId, r.Scripts)
}

func (r *CICDRunner) GetResults() {
	results := r.Server.GetRunnerResults(r.RunnerId)
	if len(results) > 0 {
		for _, result := range results {
			fmt.Println(fmt.Sprintf("[RUNNER=%s][INFO] Result: %s", r.RunnerId, result))
		}
	} else {
		fmt.Println(fmt.Sprintf("[RUNNER=%s][INFO] No results found", r.RunnerId))
	}
}