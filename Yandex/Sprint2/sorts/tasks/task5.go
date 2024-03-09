package main

import (
	"fmt"
	"sort"
)

type Worker struct {
	name       string
	position   string
	salary     uint
	experience uint
}

type Company struct {
	workers []Worker
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) {
	newWorker := Worker{name, position, salary, experience}
	c.workers = append(c.workers, newWorker)
}

func (c *Company) SortWorkers() ([]string, error) {
	sort.Slice(c.workers, func(i, j int) bool {
		if c.workers[i].salary*c.workers[i].experience != c.workers[j].salary*c.workers[j].experience {
			return c.workers[i].salary*c.workers[i].experience > c.workers[j].salary*c.workers[j].experience
		}
		if c.workers[i].position != c.workers[j].position {
			return getPositionRank(c.workers[i].position) > getPositionRank(c.workers[j].position)
		}
		return c.workers[i].name < c.workers[j].name
	})

	result := make([]string, len(c.workers))
	for i, worker := range c.workers {
		result[i] = fmt.Sprintf("%s - %d - %s", worker.name, worker.salary*worker.experience, worker.position)
	}

	return result, nil
}

func getPositionRank(position string) int {
	switch position {
	case "директор":
		return 5
	case "зам. директора":
		return 4
	case "начальник цеха":
		return 3
	case "мастер":
		return 2
	case "рабочий":
		return 1
	default:
		return 0
	}
}
