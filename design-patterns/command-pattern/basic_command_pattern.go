package command_pattern

import "fmt"

/*
Basic Command pattern implementation.
There are commands for creating, getting and deleting user by username.
Each command delegates commands to UserAPI client
 */

type UserAPI struct {
	url string
}

func (u *UserAPI) CreateUser(username string, password string, fio string) {
	fmt.Println(fmt.Sprintf("UserAPI: Create user with username=%s, password=%s, fio=%s", username, password, fio))
}

func (u *UserAPI) GetUser(username string) string {
	userDataTemplate := `
{
	"username": "%s",
	"fio": "Test User"
}
`
	userData := fmt.Sprintf(userDataTemplate, username)
	fmt.Println(fmt.Sprintf("UserAPI: Get user with username=%s: data=%s", username, userData))

	return userData
}

func (u *UserAPI) DeleteUser(username string) {
	fmt.Println("UserAPI: Deleted user with username=" + username + " successfully")
}

type EventRouter struct {
}

func (e *EventRouter) PublishEvent(event interface{}) {
	fmt.Println(fmt.Sprintf("EventRouter: Publishing event %v", event))
}


type Command interface {
	Execute()
}

type CreateUserCommand struct {
	userAPI UserAPI
	username string
	password string
	fio string
}

func (c *CreateUserCommand) Execute() {
	c.userAPI.CreateUser(c.username, c.password, c.fio)
}

type GetUserCommand struct {
	eventRouter EventRouter
	userAPI UserAPI
	username string
}

func (g *GetUserCommand) Execute() {
	user := g.userAPI.GetUser(g.username)
	g.eventRouter.PublishEvent(user)
}

type DeleteUserCommand struct {
	userAPI UserAPI
	username string
}

func (d *DeleteUserCommand) Execute() {
	d.userAPI.DeleteUser(d.username)
}

