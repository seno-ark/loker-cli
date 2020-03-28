package main

import (
	"fmt"
	"testing"
)

func TestCreateLoker(t *testing.T) {
	res, err := createLoker(0)
	if err == nil || err.Error() != message["invalid_argument"] {
		t.Errorf("createLoker(0); want %s", message["invalid_argument"])
	}

	res, err = createLoker(3)
	res_should := fmt.Sprintf(message["init_success"], 3)
	if res != res_should {
		t.Errorf("createLoker(3) = %s; want %s", res, res_should)
	}
}