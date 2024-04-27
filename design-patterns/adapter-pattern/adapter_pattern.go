package adapter_pattern

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"slices"
	"strings"
	"time"
)

type ExternalUser struct {
	id int
	username string
	createDate time.Time
}

// ExternalUserService is external API that accepts requests and returns results in XML format
type ExternalUserService struct {
	users []ExternalUser
}

func NewExternalUserService() *ExternalUserService {
	return new(ExternalUserService)
}

func (s *ExternalUserService) AddUser(username string) {
	user := new(ExternalUser)
	user.id = time.Now().Nanosecond()
	user.username = username
	user.createDate = time.Now().UTC()
	s.users = append(s.users, *user)
}

func (s *ExternalUserService) serialize(user ExternalUser) string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString("<user>")
	strBuilder.WriteString("<id>")
	strBuilder.WriteString(fmt.Sprintf("%d", user.id))
	strBuilder.WriteString("</id>")
	strBuilder.WriteString("<username>")
	strBuilder.WriteString(user.username)
	strBuilder.WriteString("</username>")
	strBuilder.WriteString("<create_date>")
	strBuilder.WriteString(user.createDate.Format("2006-01-02T15:04:05Z"))
	strBuilder.WriteString("</create_date>")
	strBuilder.WriteString("</user>")

	return strBuilder.String()
}

// GetUsers returns users list in xml format
func (s *ExternalUserService) GetUsers() string {
	result := "<users>"
	for _, user := range s.users {
		result += s.serialize(user)
	}
	result += "</users>"
	return result
}

func (s *ExternalUserService) GetUserById(id int) string {
	var foundUser *ExternalUser
	for _, user := range s.users {
		if user.id == id {
			foundUser = &user
			break
		}
	}

	if foundUser == nil {
		return ""
	}

	return s.serialize(*foundUser)
}

func (s *ExternalUserService) DeleteUserById(id int) {
	s.users = slices.DeleteFunc(s.users, func(u ExternalUser) bool { return u.id == id})
}

// DictionaryService is generic interface for all APIs that operate in dictionary items
type DictionaryService[T,R comparable] interface {
	Get(id int) *R
	List() []R
	Add(item T)
	Delete(id int)
}

type User struct {
	Id int `xml:"id" json:"id"`
	Username string `xml:"username" json:"username"`
	CreateDate time.Time `xml:"create_date" json:"createDate"`
}

func (u User) String() string {
	bytes, _ := json.Marshal(u)
	return string(bytes)
}

type users struct {
	Users []User `xml:"user"`
}

type UserServiceAdapter struct{
	userService *ExternalUserService
}

func NewUserServiceAdapter(userService *ExternalUserService) *UserServiceAdapter {
	return &UserServiceAdapter{userService}
}

func GetList(dictService DictionaryService[User, User]) []User {
	return dictService.List()
}

func (u *UserServiceAdapter) Get(id int) *User {
	xmlData := u.userService.GetUserById(id)

	fmt.Printf("Received xml data: %s\n", xmlData)

	var user User
	_ = xml.Unmarshal([]byte(xmlData), &user)


	return &user
}

func (u *UserServiceAdapter) List() []User {
	list := u.userService.GetUsers()
	fmt.Printf("Received list: %v\n", list)
	var users users
	err := xml.Unmarshal([]byte(list), &users)

	if err != nil {
		panic(err)
	}

	return users.Users
}

func (u *UserServiceAdapter) Add(item User) {
	u.userService.AddUser(item.Username)
}

func (u *UserServiceAdapter) Delete(id int) {
	u.userService.DeleteUserById(id)
}