package factory_pattern

import "strings"

type Configuration struct {
	data string
}

func (c *Configuration) String() string {
	return c.data
}

type Parser interface {
	Parse() (*Configuration, error)
}

type NoParser struct {
}

func (p *NoParser) Parse() (*Configuration, error) {
	parser := &Configuration{}
	parser.data = "empty parser"
	return parser, nil
}

type YAMLConfigParser struct {
}

func (p *YAMLConfigParser) Parse() (*Configuration, error) {
	config := &Configuration{}
	config.data = `
server:
	port: 8080
deployment:
	image: my-app:latest
	ports:
		- 8080
		- 8081
`
	return config, nil
}

type JSONConfigParser struct {
}

func (p *JSONConfigParser) Parse() (*Configuration, error) {
	result := &Configuration{}
	result.data = `
{
	"server": {
		"port": 8080
	},
	"deployment": {
		"image": "my-app:latest",
		"ports": [8080, 8081]
	}
}
`
	return result, nil
}

func NewParser(path string) Parser {
	switch {
	case strings.HasSuffix(path, ".json"):
		return &JSONConfigParser{}
	case strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml"):
		return &YAMLConfigParser{}
	default:
		return &NoParser{}
	}
}