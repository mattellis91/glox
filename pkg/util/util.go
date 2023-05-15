package util

import "fmt";

func Check(e error) {
	if e != nil {
		panic(e)
	} 
}

var HadError = false

func ErrorMessage(line int, mesage string) {
	report(line, "", mesage)
}

func report(line int, where string, message string) {
	HadError = true
	fmt.Println(fmt.Errorf("[line %d ] Error %s : %s", line, where, message).Error())
}


