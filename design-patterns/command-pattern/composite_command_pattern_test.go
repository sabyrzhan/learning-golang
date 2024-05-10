package command_pattern

import "testing"

func TestCompositeCommandPattern(t *testing.T) {
	var commands []SQLCommand
	var dataSource = &DataSource{"postgresql:jdbc://localhost/testdb"}
	commands = append(commands, &UpsertSQLCommand{dataSource, "insert into users values(1,'new user1')", map[string]interface{}{}})
	commands = append(commands, &UpsertSQLCommand{dataSource, "insert into users values(2,'new user2')", map[string]interface{}{}})
	commands = append(commands, &UpsertSQLCommand{dataSource, "insert into users values(3,'new user3')", map[string]interface{}{}})
	commands = append(commands, &UpsertSQLCommand{dataSource, "insert into users values(4,'new user4')", map[string]interface{}{}})

	var transactionalCommand = NewTransactionalBatchSQLCommand(commands)
	transactionalCommand.Execute()
}