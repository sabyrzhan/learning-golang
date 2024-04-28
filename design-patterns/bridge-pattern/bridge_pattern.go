package bridge_pattern

import (
	"fmt"
	"time"
)

/*
We have the following components after applying Bridge pattern:
- Abstraction: DeployCLI - is the main API that is used by client to make deployments
- Implementor: Deployer - is the interface and single entrypoint through which DeployCLI makes
  deployments to various platforms.
- ConcreteImplementor: We have KubernetesDeployer, BareMetalDockerDeployer and BareMetalSystemD deployers which
  all implement Deployer interface. If there will be any extension to Deployer, all the concrete implementors
  will have to implement new functionalities. But this way we hide the concrete implementations from the client by
  providing interface.
 */

type DeployTarget int
const (
	Kubernetes DeployTarget = iota
	BareMetalDocker
	BareMetalSystemD
)

type DeployCLI struct {}

func NewDeployCLIInstance() *DeployCLI {
	return new(DeployCLI)
}

func (c *DeployCLI) CreateDeployer(deployTarget DeployTarget) Deployer {
	switch deployTarget {
	case Kubernetes:
		return new(KubernetesDeployer)
	case BareMetalDocker:
		return new(BareMetalDockerDeployer)
	case BareMetalSystemD:
		return new(BareMetalSystemDDeployer)
	default:
		return nil
	}
}

type Deployer interface {
	Deploy(name string, pathToArtifact string, scale int)
	Undeploy(name string)
}

type KubernetesDeployer struct{}
func (d *KubernetesDeployer) Deploy(name string, pathToArtifact string, scale int) {
	fmt.Println(fmt.Sprintf("Kubernetes: Deploying %s: path=%s, scale=%d...", name, pathToArtifact, scale))
	time.Sleep(1 * time.Second)
	fmt.Println("Kubernetes: Deployment finished")
}
func (d *KubernetesDeployer) Undeploy(name string) {
	fmt.Println(fmt.Sprintf("Kubernetes: Undeploying %s...", name))
	time.Sleep(1 * time.Second)
	fmt.Println("Kubernetes: Undeploy finished")
}

type BareMetalDockerDeployer struct{}
func (d *BareMetalDockerDeployer) Deploy(name string, pathToArtifact string, scale int) {
	fmt.Println(fmt.Sprintf("BaremetalDocker: Deploying %s: path=%s, scale=%d...", name, pathToArtifact, scale))
	time.Sleep(1 * time.Second)
	fmt.Println("BaremetalDocker: Deployment finished")
}

func (d *BareMetalDockerDeployer) Undeploy(name string) {
	fmt.Println(fmt.Sprintf("BaremetalDocker: Undeploying %s...", name))
	time.Sleep(1 * time.Second)
	fmt.Println("BaremetalDocker: Undeploy finished")
}

type BareMetalSystemDDeployer struct{}
func (d *BareMetalSystemDDeployer) Deploy(name string, pathToArtifact string, scale int) {
	fmt.Println(fmt.Sprintf("BaremetalSystemD: Deploying %s: path=%s, scale=%d...", name, pathToArtifact, scale))
	time.Sleep(1 * time.Second)
	fmt.Println("BaremetalSystemD: Deployment finished")
}
func (d *BareMetalSystemDDeployer) Undeploy(name string) {
	fmt.Println(fmt.Sprintf("BaremetalSystemD: Undeploying %s...", name))
	time.Sleep(1 * time.Second)
	fmt.Println("BaremetalSystemD: Undeploy finished")
}