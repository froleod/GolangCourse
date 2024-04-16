package main

import (
	"errors"
	"fmt"
)

type MergeDictsJob struct {
	Dicts      []map[string]string
	Merged     map[string]string
	IsFinished bool
}

var (
	ErrNotEnoughDicts = errors.New("at least 2 dictionaries are required")
	ErrNilDict        = errors.New("nil dictionary")
)

/*
Реализуйте функцию ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error),
которая выполняет джобу MergeDictsJob и возвращает ее.

Алгоритм обработки джобы следующий:

перебрать по порядку все словари job.Dicts и записать каждое ключ-значение в результирующую мапу job.Merged
если в структуре job.Dicts меньше 2-х словарей, возвращается ошибка
errNotEnoughDicts = errors.New("at least 2 dictionaries are required")
если в структуре job.Dicts встречается словарь в виде нулевого значения nil, то возвращается
ошибка errNilDict = errors.New("nil dictionary")
независимо от успешного выполнения или ошибки в возвращаемой структуре MergeDictsJob
поле IsFinished должно быть заполнено как true
*/

// BEGIN (write your solution here)
func ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error) {

	defer func() { job.IsFinished = true }()

	if len(job.Dicts) < 2 {
		return job, ErrNotEnoughDicts
	}

	job.Merged = make(map[string]string)

	for _, dict := range job.Dicts {
		if dict == nil {
			return job, ErrNilDict
		}
		for key, value := range dict {
			job.Merged[key] = value
		}
	}

	return job, nil
}

// END

func main() {
	fmt.Println(ExecuteMergeDictsJob(&MergeDictsJob{}))
	fmt.Println(ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{{"a": "b"}, nil}}))
	fmt.Println(ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{{"a": "b"}, {"b": "c"}}}))
}
