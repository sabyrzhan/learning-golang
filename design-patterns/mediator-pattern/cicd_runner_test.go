package mediator_pattern

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRunnerJobs(t *testing.T) {
	server := &CICDServer{
		JobResults: map[string][]string{},
		Runners: map[string]*CICDRunner{},
	}
	runner1 := &CICDRunner{
		Server: server,
		RunnerId: "1",
	}

	runner2 := &CICDRunner{
		Server: server,
		RunnerId: "2",
	}

	runner3 := &CICDRunner{
		Server: server,
		RunnerId: "3",
	}

	server.RegisterRunner(runner1)
	server.RegisterRunner(runner2)
	server.RegisterRunner(runner3)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Add(1)
	wg.Add(1)

	go func() {
		defer wg.Done()
		runner1.AddScript("mv clean compile")
		runner1.AddScript("kubectl apply -f java_backend.yaml")
		runner1.SetFinishMessage("Finished deploy backend!")
		runner1.Run()
	}()

	go func() {
		defer wg.Done()
		runner2.AddScript("npm run build")
		runner2.AddScript("aws s3 cp ./build/ s3://service.frontend")
		runner2.SetFinishMessage("Finished deploy frontend!")
		runner2.Run()
	}()
	go func() {
		defer wg.Done()
		runner3.AddScript("terraform apply -f")
		runner3.AddScript("ansible-playbook -i prod.inventory prod.playbook")
		runner3.SetFinishMessage("Finished provisioning prod environment!")
		runner3.Run()
	}()

	wg.Wait()
	fmt.Println("\n\nResults:")
	wg.Add(1)
	wg.Add(1)
	wg.Add(1)

	go func() {
		defer wg.Done()
		for range []int{1,2,3} {
			runner1.GetResults()
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		for range []int{1,2,3} {
			runner2.GetResults()
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		for range []int{1,2,3} {
			runner3.GetResults()
			time.Sleep(1 * time.Second)
		}
	}()


	wg.Wait()
}