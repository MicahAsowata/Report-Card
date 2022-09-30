package report

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Report struct {
	studentName string
	results     map[string]int
	comment     string
}

func CreateReport(studentName string) Report {
	newStudentReport := Report{
		studentName: studentName,
		results:     map[string]int{},
		comment:     " ",
	}
	return newStudentReport
}

func (r *Report) AddResult(subject string, score int) {
	r.results[subject] = score
}

func (r *Report) AddComment(comment string) {
	r.comment = comment
}

func (r *Report) format() string {
	formattedReport := fmt.Sprintf("%s Result \n", r.studentName)

	for subject, score := range r.results {
		formattedReport += fmt.Sprintf("%s ...... %d \n", subject, score)
	}

	formattedReport += fmt.Sprintf("Teacher's Comment ....... %s \n", r.comment)

	todayYear, todayMonth, todayDay := time.Now().Date()
	formattedReport += fmt.Sprintf("Result given today, %v-%v-%v \n", strconv.Itoa(todayYear), strconv.Itoa(int(todayMonth)), strconv.Itoa(todayDay))

	return formattedReport
}

func (r *Report) Save() {
	reportData := []byte(r.format())
	err := os.WriteFile("reports/"+r.studentName+".txt", reportData, 0644)
	if err != nil {
		panic(err)
	}
}
