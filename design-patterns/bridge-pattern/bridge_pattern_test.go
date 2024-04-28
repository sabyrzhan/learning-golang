package bridge_pattern

import (
	"fmt"
	"testing"
	"time"
)


func TestK8sDeployer(t *testing.T) {
	cli := NewDeployCLIInstance()
	k8sDeployer := cli.CreateDeployer(Kubernetes)
	k8sDeployer.Deploy("mainApp", "registry.myapp.com/main/app:v1.0.0", 2)
	time.Sleep(1 * time.Second)
	fmt.Println()
	k8sDeployer.Undeploy("mainApp")
}

func TestBareMetalDockerDeployer(t *testing.T) {
	cli := NewDeployCLIInstance()
	dockerDeployer := cli.CreateDeployer(BareMetalDocker)
	dockerDeployer.Deploy("mainApp", "registry.myapp.com/main/app:v1.0.0", 2)
	time.Sleep(1 * time.Second)
	fmt.Println()
	dockerDeployer.Undeploy("mainApp")
}

func TestBareMetalSystemDDeployer(t *testing.T) {
	cli := NewDeployCLIInstance()
	systemDDeployer := cli.CreateDeployer(BareMetalSystemD)
	systemDDeployer.Deploy("mainApp", "/home/user/mainApp/mainApp.jar", 2)
	time.Sleep(1 * time.Second)
	fmt.Println()
	systemDDeployer.Undeploy("mainApp")
}