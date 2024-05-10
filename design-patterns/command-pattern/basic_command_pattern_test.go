package command_pattern

import "testing"

func TestCreateUser(t *testing.T) {
	userAPI := UserAPI{"https://user-api.com"}
	createUserCommand := CreateUserCommand{userAPI, "username1", "testpass1", "Test User1"}
	createUserCommand.Execute()

	deleteUserCommand := DeleteUserCommand{userAPI, "username1"}
	deleteUserCommand.Execute()

	eventRouter := EventRouter{}
	getUserCommand := GetUserCommand{eventRouter, userAPI, "username1"}
	getUserCommand.Execute()
}
