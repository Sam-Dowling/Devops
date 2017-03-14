package main

import (
	"sync"
	"time"
)

var currentID int
var mutex = &sync.Mutex{}

var tasks Tasks

// Give us some seed data
func init() {
	RepoCreateTask(Task{Name: "Write presentation", Completed: true, Due: time.Now().AddDate(0, 0, -1)}) // due yesterday
	RepoCreateTask(Task{Name: "Study for Devops", Completed: false, Due: time.Now().AddDate(0, 0, 14)})  // due 14 days in future
	RepoCreateTask(Task{Name: "Wash Car", Completed: true, Due: time.Now().AddDate(0, 1, 0)})            // due in 1 month
}

func RepoFindTask(id int) Task {
	for _, t := range tasks {
		if t.Id == id {
			return t
		}
	}
	// return empty Task if not found
	return Task{}
}

func RepoCreateTask(t Task) Task {

	mutex.Lock()

	currentID++
	t.Id = currentID
	tasks = append(tasks, t)

	mutex.Unlock()
	return t
}

func RepoDestroyTask(id int) bool {
	for i, t := range tasks {
		if t.Id == id {
			mutex.Lock()
			tasks = append(tasks[:i], tasks[i+1:]...)
			currentID--
			mutex.Unlock()
			return true
		}
	}
	return false
}

func RepoUpdateTask(id int, task Task) Task {
	for i, t := range tasks {
		if t.Id == id {
			mutex.Lock()
			tasks[i].Completed = task.Completed
			tasks[i].Due = task.Due
			tasks[i].Name = task.Name
			mutex.Unlock()
			return tasks[i]
		}
	}
	return Task{}
}
