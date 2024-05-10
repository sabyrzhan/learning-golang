package command_pattern

import (
	"fmt"
	"strings"
	"time"
)

/*
Implementation of the SQL commands.
SELECT query is implemented using QuerySQLCommand.
INSERT/UPDATE/DELETE queries are implemented using UpsertSQLCommand.

TransactionalBatchSQLCommand aggregates SQLCommands and executes as batch operations
them using Composite design pattern
 */

type SQLCommandResult struct {
	resultCode int
	query string
	resultData interface{}
}

type SQLCommand interface {
	Execute() *SQLCommandResult
}

type DataSource struct {
	url string
}

type QuerySQLCommand struct {
	ds *DataSource
	query string
}

func (q *QuerySQLCommand) Execute() *SQLCommandResult {
	fmt.Println("QuerySQLCommand:")
	return &SQLCommandResult{100, q.query, []map[string]interface{} {
		{"id": time.Now().String(), "username": "somename1", "fio": "Some Fullname1"},
		{"id": time.Now().String(), "username": "somename2", "fio": "Some Fullname2"},
		{"id": time.Now().String(), "username": "somename3", "fio": "Some Fullname3"},
	}}
}

type UpsertSQLCommand struct {
	ds *DataSource
	query string
	params map[string]interface{}
}

func (q *UpsertSQLCommand) Execute() *SQLCommandResult {
	fmt.Println("UpsertSQLCommand:", q.query, ", params:", q.params)
	return &SQLCommandResult{101, q.query, 1}
}

type TransactionalBatchSQLCommand struct {
	commands []SQLCommand
}

func (t *TransactionalBatchSQLCommand) Execute() *SQLCommandResult {
	fmt.Println("BEGIN Transaction")
	for _, command := range t.commands {
		result := command.Execute()
		fmt.Println(fmt.Sprintf("%s %s: result=%v", strings.Repeat(" ", 3), result.query, result.resultData))
	}
	fmt.Println("COMMIT Transaction")

	return &SQLCommandResult{3, "composite command", "done"}
}

func NewTransactionalBatchSQLCommand(commands []SQLCommand) SQLCommand {
	cmd := TransactionalBatchSQLCommand{}
	cmd.commands = commands
	return &cmd
}