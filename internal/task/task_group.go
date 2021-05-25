package task

import (
	"errors"
	"github.com/lithammer/shortuuid"
)

type TaskGroup struct {
	Name, ID string
	Tasks map[string]bool
}

var (
	TaskGroupDoesNotExistErr = errors.New("task group does not exist")

	taskGroups = make(map[string]*TaskGroup)
)

// DoesTaskGroupExist determines if a task group is present
func DoesTaskGroupExist(id string) bool {
	_, ok := taskGroups[id]
	return ok
}

// CreateTaskGroup creates a new task group
func CreateTaskGroup(name string) string {
	id := shortuuid.New()

	taskGroups[id] = &TaskGroup{
		Name: name,
		ID: id,
		Tasks: make(map[string]bool),
	}

	return id
}

// RemoveTaskGroup removes a specified task group
func RemoveTaskGroup(id string) error {
	if !DoesTaskGroupExist(id) {
		return TaskGroupDoesNotExistErr
	}

	delete(taskGroups, id)

	return nil
}

// GetTaskGroup gets a task group from a specified id
func GetTaskGroup(id string) (*TaskGroup, error) {
	if !DoesTaskGroupExist(id) {
		return &TaskGroup{}, TaskGroupDoesNotExistErr
	}

	return taskGroups[id], nil
}

// GetTaskIDs gets all task ids of a specified group
func GetTaskIDs(id string) ([]string, error) {
	if !DoesTaskGroupExist(id) {
		return []string{}, TaskGroupDoesNotExistErr
	}

	ids := []string{}

	taskGroup := taskGroups[id]

	for id := range taskGroup.Tasks {
		ids = append(ids, id)
	}

	return ids, nil
}

// GetAllTaskGroupIDs gets all task group ids
func GetAllTaskGroupIDs() []string {
	ids := []string{}

	for id := range taskGroups {
		ids = append(ids, id)
	}

	return ids
}