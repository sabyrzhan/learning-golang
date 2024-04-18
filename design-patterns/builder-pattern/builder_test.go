package builder_pattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullBuild(t *testing.T) {
	expected := `FROM java:latest
AUTHOR Author <author@example.com>
EXPOSE 8080 8081
USER user
COPY target/app.jar /app/app.jar
COPY target/resources/ /app/resources/
WORKDIR /app
CMD ["java", "-jar", "app.jar"]
ENTRYPOINT ["java", "-jar", "app.jar"]`

	builder := NewDockerfileBuilder()
	result, err := builder.
		From("java:latest").
		Author("Author <author@example.com>").
		AddPort(8080).AddPort(8081).
		User("user").
		Copy("target/app.jar", "/app/app.jar").
		Copy("target/resources/", "/app/resources/").
		Workdir("/app").
		Command([]string{"java", "-jar", "app.jar"}).
		Entrypoint([]string{"java", "-jar", "app.jar"}).
		Build()
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestBuildOnlyBasicDockerfile(t *testing.T) {
	expected := `FROM java:latest
EXPOSE 8080
COPY target/app.jar /app/app.jar
WORKDIR /app
CMD ["java", "-jar", "app.jar"]`
	builder := NewDockerfileBuilder()
	result, err := builder.
		From("java:latest").
		AddPort(8080).
		Copy("target/app.jar", "/app/app.jar").
		Workdir("/app").
		Command([]string{"java", "-jar", "app.jar"}).
		Build()

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestWithoutFrom(t *testing.T) {
	builder := NewDockerfileBuilder()
	_, err := builder.Build()

	assert.NotNil(t, err)
	assert.Equal(t, "from is required", err.Error())
}

func TestWithoutCmdOrEntrypoint(t *testing.T) {
	builder := NewDockerfileBuilder()
	_, err := builder.
		From("java:latest").
		Build()

	assert.NotNil(t, err)
	assert.Equal(t, "you must specify at least command or entrypoint", err.Error())
}
