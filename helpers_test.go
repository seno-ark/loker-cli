package main

import "testing"

func TestUnpackInput(t *testing.T) {

	var command, arg1, arg2 string

	command, arg1, arg2 = unpackInput([]string{})
	if command != "" || arg1 != "" || arg2 != "" {
		t.Errorf(`unpackInput([]string{} = %s %s %s`, command, arg1, arg2)
	}

	command, arg1, arg2 = unpackInput([]string{"status"})
	if command != "status" || arg1 != "" || arg2 != "" {
		t.Errorf(`unpackInput([]string{"status"} = %s %s %s`, command, arg1, arg2)
	}

	command, arg1, arg2 = unpackInput([]string{"init", "5"})
	if command != "init" || arg1 != "5" || arg2 != "" {
		t.Errorf(`unpackInput([]string{"init", "5"} = %s %s %s`, command, arg1, arg2)
	}

	command, arg1, arg2 = unpackInput([]string{"input", "SIM", "123"})
	if command != "input" || arg1 != "SIM" || arg2 != "123" {
		t.Errorf(`unpackInput([]string{"input", "SIM", "123"} = %s %s %s`, command, arg1, arg2)
	}
}