package main

import (
	"bufio"
	"fmt"
	"os"
	"report_card/report"
	"strconv"
	"strings"
)

func readFromTerminal(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	str, err := reader.ReadString('\n')
	trimmedStr := strings.Trim(str, "\n")
	return trimmedStr, err
}
func main() {
	studentName, _ := readFromTerminal("Enter the student's name")
	newReport := report.CreateReport(studentName)

	for {
		letter, _ := readFromTerminal("Press a - to add result, k - to add comment, s - save ")
		if letter == "a" {
			subject, _ := readFromTerminal("Enter the subject")
			getScore, _ := readFromTerminal("Enter the score ")
			score, _ := strconv.Atoi(getScore)
			newReport.AddResult(subject, score)
		} else if letter == "k" {
			comment, _ := readFromTerminal("Enter your comments")
			newReport.AddComment(comment)
		} else if letter == "s" {
			newReport.Save()
			break
		} else {
			fmt.Println("That is an invalid key")
		}
	}

}
