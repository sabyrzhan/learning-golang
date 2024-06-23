package observer_pattern

import (
	"fmt"
	"testing"
)

func TestHooks(t *testing.T) {
	g := NewGitClient()
	userInfoHook := UserMessageHook{Username: "John Doe", Email: "johndoe@gmail.com"}
	g.AddHook("pull", userInfoHook)
	g.AddHook("push", userInfoHook)
	g.AddHook("commit", userInfoHook)

	prePullHook := PrePullHook{}
	g.AddHook("pull", prePullHook)

	prePushHook := PrePushHook{}
	g.AddHook("push", prePushHook)

	commitHook := PreCommitHook{}
	g.AddHook("commit", commitHook)

	g.Add(".gitignore")
	g.Add("pom.xml")
	g.Add("src/Main.java")
	g.Add("src/User.java")
	g.Add("src/Profile.java")

	fmt.Println()
	g.Commit()
	fmt.Println()
	g.Pull()
	fmt.Println()
	g.Push()
}