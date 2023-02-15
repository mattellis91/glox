package reporting

import "fmt"

var HadError = false

func ErrorMessage(line int, mesage string) {
	report(line, "", mesage)
}

func report(line int, where string, message string) error {
	HadError = true
	return fmt.Errorf("[line %d ] Error %s : %s", line, where, message)
}