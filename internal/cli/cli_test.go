package cli

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCli(t *testing.T) {
	path := "task_list_test.csv"

	defer os.Remove(path)

	data := []byte("6ba7b810-9dad-11d1-80b4-00c04fd430c8,Test Task 1,This is a test task,Not Started\n" + "4e8027c4-3f46-4a26-bac2-1faee7735bea,Test Task 2,This is another test task,In Progress\n")

	err := os.WriteFile(path, data, 0666)
	assert.NoError(t, err)

	app := TaskManagerCli(path)

	err = app.Run([]string{"godo", "list"})
	assert.NoError(t, err)

	err = app.Run([]string{"godo", "list", "shouldnt-be-here"})
	assert.Error(t, err)

	err = app.Run([]string{"godo", "add", "Test task", "Test task description"})
	assert.NoError(t, err)

	err = app.Run([]string{"godo", "add", "Just the title"})
	assert.Error(t, err)

	err = app.Run([]string{"godo", "add", "1", "Can't be a number"})
	assert.Error(t, err)

	err = app.Run([]string{"godo", "add", "Can't be a number", "1"})
	assert.Error(t, err)

	err = app.Run([]string{"godo", "remove", "6ba7b810-9dad-11d1-80b4-00c04fd430c8"})
	assert.NoError(t, err)

	err = app.Run([]string{"godo", "update", "4e8027c4-3f46-4a26-bac2-1faee7735bea", "Completed"})
	assert.NoError(t, err)
}
