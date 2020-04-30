package wrkhub_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/eriktate/wrkhub"
)

func Test_TaskJSON(t *testing.T) {
	task := wrkhub.Task{
		ID:          wrkhub.NewUniqueID(),
		Title:       "Test Task",
		Description: "Test Description",
	}

	data, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("failed to marshal task: %s", err)
	}

	var decodedTask wrkhub.Task
	if err := json.Unmarshal(data, &decodedTask); err != nil {
		t.Fatalf("failed to unmarshal task: %s", err)
	}

	if !task.ID.IsEqual(decodedTask.ID) {
		t.Fatal("expected matching data")
	}
}

func Test_TaskJSON(t *testing.T) {
	task := wrkhub.Task{
		ID:          ,
		Title:       "Test Task",
		Description: "Test Description",
	}

	data, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("failed to marshal task: %s", err)
	}

	fmt.Println(string(data))
	var decodedTask wrkhub.Task
	if err := json.Unmarshal(data, &decodedTask); err != nil {
		t.Fatalf("failed to unmarshal task: %s", err)
	}

	if !task.ID.IsEqual(decodedTask.ID) {
		t.Fatal("expected matching data")
	}
}
