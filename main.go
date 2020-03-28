package main

import (
	"fmt"
	"strings"
	"strconv"
	"errors"
	"regexp"
)

type idCard struct {
	Type string
	No int
}
var loker []idCard

func initHandler(capacity string) {
	// Fungsi untuk menangani perintah `init`

	var res string
	var err error
	var lokerCap int

	if capacity == "" {
		err = errors.New(message["empty_argument"])
		showResult(res, err)
		return
	}

	lokerCap, err = strconv.Atoi(capacity)
	if err != nil || lokerCap < 1 {
		err = errors.New(message["invalid_argument"])
		showResult(res, err)
		return
	}

	res, err = initLoker(lokerCap)
	showResult(res, err)
}

func statusHandler() {
	// Fungsi untuk menangani perintah `status`

	res, err := lokerStatus()
	showResult(res, err)
}

func inputHandler(cardType, cardNo string) {
	// Fungsi untuk menangani perintah `input`

	var res string
	var err error
	var cardNumber int

	if cardType == "" || cardNo == "" {
		err = errors.New(message["empty_argument"])
		showResult(res, err)
		return
	}

	regex, _ := regexp.Compile(`^[a-zA-Z]+$`)
	cardTypeIsAlphabet := regex.MatchString(cardType)

	cardNumber, err = strconv.Atoi(cardNo)

	if err != nil || cardNumber < 1 || !cardTypeIsAlphabet {
		err = errors.New(message["invalid_argument"])
		showResult(res, err)
		return
	}

	res, err = inputIdCard(cardType, cardNumber)
	showResult(res, err)
}

func leaveHandler(lokerNo string) {
	// Fungsi untuk menangani perintah `leave`

	var res string
	var err error
	var lokerId int

	if lokerNo == "" {
		err = errors.New(message["empty_argument"])
		showResult(res, err)
		return
	}

	lokerId, err = strconv.Atoi(lokerNo)
	if err != nil || lokerId < 1 {
		err = errors.New(message["invalid_argument"])
		showResult(res, err)
		return
	}

	res, err = leaveIdCard(lokerId)
	showResult(res, err)
}

func findHandler(cardNo string) {
	// Fungsi untuk menangani perintah `find`

	var res string
	var err error
	var cardNumber int

	if cardNo == "" {
		err = errors.New(message["empty_argument"])
		showResult(res, err)
		return
	}

	cardNumber, err = strconv.Atoi(cardNo)
	if err != nil || cardNumber < 1 {
		err = errors.New(message["invalid_argument"])
		showResult(res, err)
		return
	}

	res, err = findIdCard(cardNumber)
	showResult(res, err)
}

func searchHandler(cardType string) {
	// Fungsi untuk menangani perintah `search`
	
	var res string
	var err error

	if cardType == "" {
		err = errors.New(message["empty_argument"])
		showResult(res, err)
		return
	}

	res, err = searchIdCard(cardType)
	showResult(res, err)
}

func process(userInput []string) bool {
	// Fungsi untuk memroses inputan user

	command, arg1, arg2 := unpackInput(userInput)
	command = strings.ToLower(command)
	var err error

	switch command {

		case "init":
			initHandler(arg1)
		case "status":
			statusHandler()
		case "input":
			inputHandler(arg1, arg2)
		case "leave":
			leaveHandler(arg1)
		case "find":
			findHandler(arg1)
		case "search":
			searchHandler(arg1)
		case "exit":
			return false
		case "":
			err = errors.New(message["empty_command"])
			showResult("", err)
		default:
			err = errors.New(message["invalid_command"])
			showResult("", err)
	}

	return true
}

func main() {	

	var userInput []string
	fmt.Println(message["welcome"])

	i := 0
	for {
		userInput = getInput()
		if !process(userInput) {
			break
		}
		i++
	}
	fmt.Println(message["exit"])
}