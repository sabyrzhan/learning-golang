package flyweight_pattern

import (
	"fmt"
	"strings"
)
/*
Here I am simulating the TableView's ReusableCell of iOS.
To generate the content in a TableView I am reusable the same Cell to render the content
instead of creating new cell for each row
 */

type AccountData struct {
	UserFullName string
	AccountNumber string
	Balance float32
	Currency string
}

func (a *AccountData) GetIcon() string {
	usernames := strings.Split(a.UserFullName, " ")
	return strings.ToUpper(fmt.Sprintf("%c.%c.", usernames[0][0], usernames[1][0]))
}

type TableView struct {
	Accounts []AccountData
	ReusableCell *Cell
}

func NewTableView(accountData []AccountData) *TableView {
	return &TableView{accountData, &Cell{width: 200, height: 70, backgroundColor: "white"}}
}

func (t* TableView) ShowData() {
	var cell* Cell
	for _, account := range t.Accounts {
		cell = t.ReusableCell
		cell.Icon = account.GetIcon()
		cell.Label = account.UserFullName
		content := fmt.Sprintf("%s\n", account.AccountNumber)
		content += fmt.Sprintf("%f %s", account.Balance, account.Currency)
		cell.Content = content
		cell.Render()
	}
}

type Cell struct {
	width int
	height int
	backgroundColor string

	Label string
	Content string
	Icon string
}

func (c* Cell) Render() {
	fmt.Printf("================================================\n")
	fmt.Printf("| %s |\n", c.Icon)
	fmt.Printf("%s\n", c.Label)
	fmt.Printf("%s\n", c.Content)
	fmt.Printf("================================================\n\n")
}