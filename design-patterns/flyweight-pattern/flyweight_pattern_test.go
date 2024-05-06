package flyweight_pattern

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestFlyweight(*testing.T) {
	accountData := generateAccountData()
	tableView := NewTableView(accountData)
	tableView.ShowData()
}

func generateAccountData() []AccountData {
	var nameGenerator = func() string {
		name := ""
		for j := 0; j < 10; j++ {
			name += fmt.Sprintf("%c%c", rune(65 + rand.Intn(25)+1), rune(65 + rand.Intn(25)+1))
		}

		return name
	}
	var result []AccountData
	var data AccountData
	for i := 0; i < 100; i++ {
		data = AccountData{}
		data.AccountNumber = fmt.Sprintf("KZ%d", 1000000000000+rand.Intn(99999999999))
		data.UserFullName = nameGenerator() + " " + nameGenerator()
		data.Balance = float32(rand.Intn(100000) + 1000)
		data.Currency = []string{"KZT", "USD", "EUR"}[rand.Intn(3)]
		result = append(result, data)
	}

	return result
}
