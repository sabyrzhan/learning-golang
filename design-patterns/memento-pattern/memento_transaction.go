package memento_pattern

import (
	"errors"
	"fmt"
)

// Database transaction rollback simulation
// PersistenceManager - ORM responsible for data persistence
// Memento - stores data snapshot
// Caretaker - manages adding and rollback of Memento data. Also stores current pointer to active memento
// It saves snapshots whenever PM saves or updates entities
// Entity, UsersTable - ORM entities
// In this simulation when the user sets username=invaliduser, PM rollback the saved change and returns error
// When PM rolls back it asks caretaker the previous state of the table records.
// Caretaker decrements the pointer to set to previous memento and returns that object

type Memento struct {
	tableName string
	rowValues []interface{}
}

func NewMemento(tableName string, values []interface{}) *Memento {
	m := &Memento{
		tableName: tableName,
		rowValues: values,
	}

	return m
}

type Caretaker struct {
	mementos map[string][]*Memento
	rollbackIndex int
}

func NewCaretaker() *Caretaker {
	return &Caretaker{
		mementos: map[string][]*Memento{},
		rollbackIndex: -1,
	}
}

func (c *Caretaker) AddMemento(tableName string, memento *Memento) {
	if c.mementos[tableName] == nil {
		c.mementos[tableName] = []*Memento{}
	}
	c.mementos[tableName] = append(c.mementos[tableName], memento)
	c.rollbackIndex++
}

func (c *Caretaker) Rollback(tableName string) *Memento {
	if len(c.mementos[tableName]) == 0 || c.rollbackIndex < 0 {
		return nil
	}

	c.rollbackIndex--
	memento := c.mementos[tableName][c.rollbackIndex]
	return memento
}

type Entity interface {
	GetId() int
	GetTableName() string
	GetTableColumns() []string
	SetValue(columnIndex int, value interface{})
	GetValue(columnIndex int) interface{}
}

type UsersTable struct {
	columns []string
	values []interface{}
}

func NewUsersTable() UsersTable {
	columns := []string{"id", "full_name", "username", "email", "password"}
	return UsersTable{
		columns: columns,
		values: make([]interface{}, len(columns)),
	}
}

func (u UsersTable) GetId() int {
	if len(u.values) == 0 {
		return 0
	}

	return u.values[0].(int)
}

func (u UsersTable) GetTableName() string {
	return "users"
}

func (u UsersTable) GetTableColumns() []string {
	return []string{"id", "full_name", "username", "email", "password"}
}

func (u UsersTable) SetId(id int) {
	u.SetValue(0, id)
}

func (u UsersTable) SetFullName(fullName string) {
	u.SetValue(1, fullName)
}
func (u UsersTable) GetFullName() string {
	return u.GetValue(1).(string)
}

func (u UsersTable) SetUsername(username string) {
	u.SetValue(2, username)
}

func (u UsersTable) GetUsername() string {
	return u.GetValue(2).(string)
}

func (u UsersTable) SetEmail(email string) {
	u.SetValue(3, email)
}

func (u UsersTable) GetEmail() string {
	return u.GetValue(3).(string)
}

func (u UsersTable) SetPassword(password string) {
	u.SetValue(4, password)
}

func (u UsersTable) GetPassword() string {
	return u.GetValue(4).(string)
}

func (u UsersTable) SetValue(columnIndex int, value interface{}) {
	u.values[columnIndex] = value
}

func (u UsersTable) GetValue(columnIndex int) interface{} {
	if len(u.values) == 0 {
		return nil
	}

	return u.values[columnIndex]
}

type PersistenceManager struct {
	entities map[string][]Entity
	careTaker *Caretaker
}

func NewPersistenceManager() *PersistenceManager {
	return &PersistenceManager{
		careTaker: NewCaretaker(),
		entities:  map[string][]Entity{},
	}
}

func (m *PersistenceManager) GetUserByUsername(tableName string, username string) (Entity, error) {
	switch tableName {
	case "users":
		users := m.entities[tableName]
		for _, user := range users {
			if user.GetValue(2) == username {
				return user, nil
			}
		}

		return nil, errors.New(fmt.Sprintf("user with username=%s not found", username))
	default:
		return nil, errors.New("unsupported table")
	}
}

func (m *PersistenceManager) GetById(tableName string, id int) (Entity, error) {
	switch tableName {
	case "users":
		users := m.entities[tableName]
		for _, user := range users {
			if user.GetValue(0) == id {
				return user, nil
			}
		}

		return nil, errors.New(fmt.Sprintf("user with id=%d not found", id))
	default:
		return nil, errors.New("unsupported table")
	}
}

func (m *PersistenceManager) Update(entity Entity) (Entity, error) {
	switch entity.GetTableName() {
	case "users":
		_, ok := m.GetById("users", entity.GetValue(0).(int))
		if ok != nil {
			return nil, ok
		}

		var index int
		for i, en := range m.entities[entity.GetTableName()] {
			if en.GetId() == entity.GetId() {
				index = i
				break
			}
		}

		m.entities[entity.GetTableName()][index] = entity
		m.addToMemento(entity)

		if entity.GetValue(2) == "invaliduser" {
			entity = m.rollBackFromMemento(entity)
			m.entities[entity.GetTableName()][index] = entity
			return nil, errors.New(fmt.Sprintf("transaction rolled back: username cannot be named invaliduser"))
		}

		return entity, nil
	default:
		return nil, errors.New("unsupported table")
	}
}

func (m *PersistenceManager) Save(entity Entity) error {
	switch entity.GetTableName() {
	case "users":
		_, ok := m.GetUserByUsername(entity.GetTableName(), entity.GetValue(2).(string))
		if ok == nil {
			return fmt.Errorf("user with this username already exists")
		}

		m.addToMemento(entity)
		m.entities[entity.GetTableName()] = append(m.entities[entity.GetTableName()], entity)

		return nil
	default:
		return errors.New("unsupported table")
	}
}

func (m *PersistenceManager) addToMemento(entity Entity) {
	values := make([]interface{}, 0)
	for i := range len(entity.GetTableColumns()) {
		values = append(values, entity.GetValue(i))
	}
	m.careTaker.AddMemento(entity.GetTableName(), NewMemento(entity.GetTableName(), values))
}

func (m *PersistenceManager) rollBackFromMemento(entity Entity) Entity {
	oldValue := m.careTaker.Rollback(entity.GetTableName())
	for i, value := range oldValue.rowValues {
		entity.SetValue(i, value)
	}

	return entity
}
