package main

import (
	"fmt"
	"sort"
	"errors"
)

type CompanyInterface interface{
    AddWorkerInfo(name, position string, salary, experience uint) error
    SortWorkers() ([]string, error) 
}

type Company struct {
	workers []Worker
}

type Worker struct {
	name string
	position string
	salary uint
	experience uint
}

// Мапа для определения приоритета должностей (чем меньше число — тем выше должность)
var positionPriority = map[string]int{
	"директор":       0,
	"зам. директора": 1,
	"начальник цеха": 2,
	"мастер":         3,
	"рабочий":        4,
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	if _, exist := positionPriority[position]; !exist {
		return errors.New("invalid position")
	}

	if name == "" {
		return errors.New("invalid name")
	}

	c.workers = append(c.workers, Worker{
		name: name,
		position: position,
		salary: salary,
		experience: experience,
	})
	return nil
}

func (c *Company) SortWorkers() ([]string, error) {
	if len(c.workers) == 0 {
		return nil, errors.New("no workers")
	}

	workers := make([]Worker, len(c.workers))
	copy(workers, c.workers)

	sort.Slice(workers, func(i, j int) bool {
		incomeI := workers[i].salary * workers[i].experience
		incomeJ := workers[j].salary * workers[j].experience

		if incomeI != incomeJ {
			return incomeI > incomeJ
		}

		priorityI := positionPriority[workers[i].position]
		priorityJ := positionPriority[workers[j].position]
		if priorityI != priorityJ {
			return priorityI < priorityJ
		}

		return workers[i].name < workers[j].name
	})

	result := make([]string, len(workers))
	for i, w := range workers {
		income := w.salary * w.experience
		result[i] = fmt.Sprintf("%s — %d — %s", w.name, income, w.position)
	}

	return result, nil
}

func main() {
	var comp Company

	_ = comp.AddWorkerInfo("Иван", "рабочий", 50000, 12)
	_ = comp.AddWorkerInfo("Пётр", "мастер", 60000, 10)
	_ = comp.AddWorkerInfo("Мария", "директор", 150000, 24)
	_ = comp.AddWorkerInfo("Сергей", "зам. директора", 120000, 20)

	list, err := comp.SortWorkers()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	for _, s := range list {
		fmt.Println(s)
	}
}
