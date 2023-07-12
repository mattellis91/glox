package util

import "fmt";

func Check(e error) {
	if e != nil {
		panic(e)
	} 
}

var HadError = false

func ErrorMessage(line int, mesage string) string {
	return report(line, "", mesage)
}

func report(line int, where string, message string) string {
	HadError = true
	return fmt.Errorf("[line %d ] Error %s : %s", line, where, message).Error()
}


