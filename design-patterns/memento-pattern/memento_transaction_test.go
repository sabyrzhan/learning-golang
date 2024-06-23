package memento_pattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveNewUser_Success(t *testing.T) {
	pm := NewPersistenceManager()
	validUser := NewUsersTable()
	validUser.SetId(1)
	validUser.SetFullName("Jason Bourne")
	validUser.SetUsername("jasonbourne")
	validUser.SetEmail("jasonbourne@gmail.com")
	validUser.SetPassword("jasonbournepass")

	err := pm.Save(validUser)
	assert.Nil(t, err)
	savedEntity, err := pm.GetById("users", 1)
	assert.Nil(t, err)
	savedUser := savedEntity.(UsersTable)
	assert.Equal(t, "Jason Bourne", savedUser.GetFullName())
	assert.Equal(t, "jasonbourne", savedUser.GetUsername())
	assert.Equal(t, "jasonbourne@gmail.com", savedUser.GetEmail())
	assert.Equal(t, "jasonbournepass", savedUser.GetPassword())
}

func TestMementoTransaction(t *testing.T) {
	pm := NewPersistenceManager()
	validUser := NewUsersTable()
	validUser.SetId(2)
	validUser.SetFullName("James Bond")
	validUser.SetUsername("jamesbond")
	validUser.SetEmail("jamesbond@mi6.gov.uk")
	validUser.SetPassword("bond007pass")
	err := pm.Save(validUser)
	assert.Nil(t, err)

	validUser = NewUsersTable()
	validUser.SetId(2)
	validUser.SetUsername("invaliduser")
	_, err = pm.Update(validUser)
	assert.NotNil(t, err)

	rollbackEntity, err := pm.GetById("users", validUser.GetId())
	assert.Nil(t, err)
	rollbackUser := rollbackEntity.(UsersTable)
	assert.Equal(t, "James Bond", rollbackUser.GetFullName())
	assert.Equal(t, "jamesbond", rollbackUser.GetUsername())
	assert.Equal(t, "jamesbond@mi6.gov.uk", rollbackUser.GetEmail())
	assert.Equal(t, "bond007pass", rollbackUser.GetPassword())
}
