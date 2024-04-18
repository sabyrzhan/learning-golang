package builder_pattern

import (
	"fmt"
	"strings"
)

type file struct {
	source string
	target string
}

type dockerfile struct {
	from       string
	author     string
	expose     []string
	user       string
	copy       []file
	workdir    string
	command    []string
	entrypoint []string
}

type DockerfileBuilder struct {
	dockerfile dockerfile
}

func NewDockerfileBuilder() *DockerfileBuilder {
	return &DockerfileBuilder{dockerfile{}}
}

func (b *DockerfileBuilder) From(from string) *DockerfileBuilder {
	b.dockerfile.from = from
	return b
}

func (b *DockerfileBuilder) Author(author string) *DockerfileBuilder {
	b.dockerfile.author = author
	return b
}

func (b *DockerfileBuilder) AddPort(port int) *DockerfileBuilder {
	b.dockerfile.expose = append(b.dockerfile.expose, fmt.Sprintf("%d", port))
	return b
}

func (b *DockerfileBuilder) User(user string) *DockerfileBuilder {
	b.dockerfile.user = user
	return b
}

func (b *DockerfileBuilder) Copy(source string, dst string) *DockerfileBuilder {
	b.dockerfile.copy = append(b.dockerfile.copy, file{source, dst})
	return b
}

func (b *DockerfileBuilder) Command(command []string) *DockerfileBuilder {
	b.dockerfile.command = command
	return b
}

func (b *DockerfileBuilder) Entrypoint(entrypoint []string) *DockerfileBuilder {
	b.dockerfile.entrypoint = entrypoint
	return b
}

func (b *DockerfileBuilder) Workdir(workdir string) *DockerfileBuilder {
	b.dockerfile.workdir = workdir
	return b
}

func (b *DockerfileBuilder) validate() error {
	if b.dockerfile.from == "" {
		return fmt.Errorf("from is required")
	}

	if len(b.dockerfile.command) == 0 && len(b.dockerfile.entrypoint) == 0 {
		return fmt.Errorf("you must specify at least command or entrypoint")
	}

	return nil
}

func (b *DockerfileBuilder) Build() (string, error) {
	validationError := b.validate()
	if validationError != nil {
		return "", validationError
	}

	var result strings.Builder
	result.WriteString("FROM ")
	result.WriteString(b.dockerfile.from)
	if b.dockerfile.author != "" {
		result.WriteString("\nAUTHOR ")
		result.WriteString(b.dockerfile.author)
	}
	if len(b.dockerfile.expose) != 0 {
		result.WriteString("\nEXPOSE ")
		result.WriteString(strings.Join(b.dockerfile.expose, " "))
	}
	if b.dockerfile.user != "" {
		result.WriteString("\nUSER ")
		result.WriteString(b.dockerfile.user)
	}
	if len(b.dockerfile.copy) != 0 {
		for _, file := range b.dockerfile.copy {
			result.WriteString("\nCOPY " + file.source + " " + file.target)
		}
	}
	if b.dockerfile.workdir != "" {
		result.WriteString("\nWORKDIR " + b.dockerfile.workdir)
	}

	if len(b.dockerfile.command) != 0 {
		if len(b.dockerfile.command) != 0 {
			var cmds []string
			for _, cmd := range b.dockerfile.command {
				cmds = append(cmds, fmt.Sprintf("\"%s\"", cmd))
			}
			result.WriteString("\nCMD [" + strings.Join(cmds, ", ") + "]")
		}
	}

	if len(b.dockerfile.entrypoint) != 0 {
		if len(b.dockerfile.entrypoint) != 0 {
			var cmds []string
			for _, cmd := range b.dockerfile.entrypoint {
				cmds = append(cmds, fmt.Sprintf("\"%s\"", cmd))
			}
			result.WriteString("\nENTRYPOINT [" + strings.Join(cmds, ", ") + "]")
		}
	}

	return result.String(), nil
}
