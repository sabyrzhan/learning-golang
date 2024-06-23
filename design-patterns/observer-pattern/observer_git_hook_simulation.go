package observer_pattern

import (
	"fmt"
	"reflect"
	"strings"
)

// Simulates git hooks
// Whenever pull, push or commit is executed, registered hooks are invoked before any operations
// Hooks are implemented as observers
// And GitClient as observable

type GitClient struct {
	PullHooks []Observer
	PushHooks []Observer
	CommitHooks []Observer
	PullFiles []File
	Files []File
}

func (g* GitClient) Add(file string) {
	g.Files = append(g.Files, File{file})
}

func (g *GitClient) AddHook(hookType string, hook Observer) {
	switch hookType {
	case "pull":
		g.PullHooks = append(g.PullHooks, hook)
	case "push":
		g.PushHooks = append(g.PushHooks, hook)
	case "commit":
		g.CommitHooks = append(g.CommitHooks, hook)
	default:
		fmt.Println("Unsupported hook type:", hookType)
	}
}

func (g *GitClient) Pull() {
	fmt.Println("===== Running pull hooks =====")
	for _, hook := range g.PullHooks {
		fmt.Println(fmt.Sprintf("::::: Hook: %s", reflect.TypeOf(hook).Name()))
		hook.Notify(g.PullFiles)
	}
	g.PullFiles = make([]File, 0)
	fmt.Println("===== Complete running pull hooks =====")

	fmt.Println("Pulled the files")
}

func (g *GitClient) Push() {
	fmt.Println("===== Running push hooks =====")
	for _, hook := range g.PushHooks {
		fmt.Println("::::: Hook: " + reflect.TypeOf(hook).Name())
		hook.Notify(g.Files)
	}
	fmt.Println("===== Complete running push hooks =====")

	if len(g.Files) == 0 {
		fmt.Println("Nothing to push")
	} else {
		fmt.Println("Pushed the files")
	}
}

func (g *GitClient) Commit() {
	fmt.Println("===== Running commit hooks =====")
	for _, hook := range g.CommitHooks {
		fmt.Println("::::: Hook: " + reflect.TypeOf(hook).Name())
		hook.Notify(g.Files)
	}
	fmt.Println("===== Complete running commit hooks =====")
	if len(g.Files) == 0 {
		fmt.Println("Nothing to commit")
	} else {
		fmt.Println("Committed the files")
	}
}

type File struct {
	Name string
}

func NewGitClient() *GitClient {
	pullFiles := make([]File, 0)
	pullFileNames := []string{".gitignore", "pom.xml", "src/Main.java", "src/User.java", "src/Profile.java"}
	for _, fileName := range pullFileNames {
		pullFiles = append(pullFiles, File{Name: fileName})
	}

	return &GitClient{
		PullHooks: make([]Observer, 0),
		CommitHooks: make([]Observer, 0),
		PushHooks: make([]Observer, 0),
		Files:     make([]File, 0),
		PullFiles: make([]File, 0),
	}
}

type Observer interface {
	Notify(data interface{})
}

type UserMessageHook struct {
	Username string
	Email string
}

func (p UserMessageHook) Notify(data interface{}) {
	fmt.Println(fmt.Sprintf("%s::::: Username: %s", strings.Repeat(" ", 3), p.Username))
	fmt.Println(fmt.Sprintf("%s::::: Email: %s", strings.Repeat(" ", 3), p.Email))
}

type PrePushHook struct {
}

func (p PrePushHook) Notify(data interface{}) {
	files := data.([]File)
	for _, file := range files {
		fmt.Println(fmt.Sprintf("%s::::: File: %s", strings.Repeat(" ", 3), file.Name))
	}

	if len(files) == 0 {
		fmt.Println(fmt.Sprintf("%sStatus: no files to push", strings.Repeat(" ", 3)))
	}
}

type PreCommitHook struct {}
func (p PreCommitHook) Notify(data interface{}) {
	files := data.([]File)
	for _, file := range files {
		fmt.Println(fmt.Sprintf("%s::::: File: %s", strings.Repeat(" ", 3), file.Name))
	}

	if len(files) == 0 {
		fmt.Println(fmt.Sprintf("%sStatus: no files to commit", strings.Repeat(" ", 3)))
	}
}

type PrePullHook struct {}

func (p PrePullHook) Notify(data interface{}) {
	files := data.([]File)
	for _, file := range files {
		fmt.Println(fmt.Sprintf("%s::::: File: %s", strings.Repeat(" ", 3), file.Name))
	}

	if len(files) == 0 {
		fmt.Println(fmt.Sprintf("%sStatus: no files to pull", strings.Repeat(" ", 3)))
	}
}