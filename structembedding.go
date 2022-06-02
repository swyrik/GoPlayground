package main

import "fmt"

func structembeddingtest() {

	govtEmp := employee{
		job: job{
			department: "education",
			hr:         "harish",
		},
		empname: "swyrik",
		empage:  27,
	}

	fmt.Println(govtEmp)

	fmt.Println(govtEmp.describe())

	type describer interface {
		describe() string
	}

	var d describer = govtEmp
	fmt.Println(d.describe())
}

type job struct {
	department string
	hr         string
}

func (j job) describe() string {
	return fmt.Sprintf("job department %v, hr %v\n", j.department, j.hr)
}

type employee struct {
	job
	empname string
	empage  int
}
