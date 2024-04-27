package adapter_pattern

import (
	"fmt"
	"testing"
)

func TestXmlToJsonServiceAdapter(t *testing.T) {
	externalUserService := NewExternalUserService()
	adapter := NewUserServiceAdapter(externalUserService)
	adapter.Add(User{Username: "testuser1"})
	adapter.Add(User{Username: "testuser2"})
	adapter.Add(User{Username: "testuser3"})
	fmt.Println(GetList(adapter))
}