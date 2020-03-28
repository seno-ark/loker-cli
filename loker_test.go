package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestInitLoker(t *testing.T) {
	res, _ := initLoker(5)

	expect := fmt.Sprintf(message["init_success"], 5)
	if res != expect {
		t.Errorf("createLoker(5) = %s; %s", res, expect)
	}
}

func TestLokerStatus(t *testing.T) {
	res, _ := lokerStatus()

	if res == "" {
		t.Errorf("lokerStatus() = %s;", res)
	}
}

func TestInputIdCard(t *testing.T) {
	var cardType string
	var cardNo int

	for i := 0; i<5; i++ {

		if i%2 == 0 {
			cardType = "SIM"
		} else {
			cardType = "ktp"
		}

		cardNo = 1000+i
		res, _ := inputIdCard(cardType, cardNo)

		expect := fmt.Sprintf(message["input_success"], i+1)
		if res != expect {
			t.Errorf(`inputIdCard("%s", %d) = %s; %s`, cardType, cardNo, res, expect)
		}
	}
}

func TestLeaveIdCard(t *testing.T) {
	res, _ := leaveIdCard(1)

	expect := fmt.Sprintf(message["leave_success"], 1)
	if res != expect {
		t.Errorf(`leaveIdCard(1) = %s; %s`, res, expect)
	}
}

func TestFindIdCard(t *testing.T) {
	res, _ := findIdCard(1001)

	expect := fmt.Sprintf(message["find_result"], 2)
	if res != expect {
		t.Errorf(`findIdCard(1001) = %s; %s`, res, expect)
	}
}

func TestSearchIdCard(t *testing.T) {
	res, _ := searchIdCard("sim")

	cardNums := strings.Join([]string{"1002", "1004"}, ", ")
	expect := fmt.Sprintf("No Identitas: %s", cardNums)
	if res != expect {
		t.Errorf(`searchIdCard("sim") = %s; %s`, res, expect)
	}
}