package factory_pattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYamlParser(t *testing.T) {
	expected := `
server:
	port: 8080
deployment:
	image: my-app:latest
	ports:
		- 8080
		- 8081
`
	parser := NewParser("/path/to/dev.yaml")
	config, err := parser.Parse()
	assert.Nil(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, expected, config.String())
}

func TestJsonParser(t *testing.T) {
	expected := `
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
	parser := NewParser("/path/to/dev.json")
	config, err := parser.Parse()
	assert.Nil(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, expected, config.String())
}

func TestNoParser(t *testing.T) {
	parser := NewParser("/path/to/invalid.config")
	config, err := parser.Parse()
	assert.Nil(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, "empty parser", config.String())
}