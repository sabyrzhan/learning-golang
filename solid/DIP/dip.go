package DIP

// DIP states high level classes or modules should not depend on low level classes or modules. They must interact
// using abstractions. For example, we have a DB storage that provides access to data through DB driver. Instead of
// accessing directly through driver specific or low level API, we should access it using separate abstraction layer
// like interface. This way we wont depend on specific low level storage. If the storage type will be file or network,
// we can swap the implementation and still communicate through abstract interfaces.

type Person struct {
	Name string
}

type QueryManager struct {
	personQuery PersonQuery
}

func (qm *QueryManager) GetPersonByName(name string) *Person {
	return qm.personQuery.GetPerson(name)
}

func NewQueryManager(personQuery PersonQuery) *QueryManager {
	return &QueryManager{
		personQuery: personQuery,
	}
}

type PersonQuery interface {
	GetPerson(name string) *Person
}

type PersonStorage struct {
	Persons []Person
}

func (storage PersonStorage) GetPerson(name string) *Person {
	for _, person := range storage.Persons {
		if person.Name == name {
			return &person
		}
	}

	return nil
}
