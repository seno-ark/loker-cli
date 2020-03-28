package main

import (
	"fmt"
	"errors"
	"strings"
)

func initLoker(capacity int) (res string, err error){

	if len(loker) > 0 {
		err = errors.New(message["init_duplicate"])
		return	
	}

	// slice berisi element berupa zero value idCard
	loker = make([]idCard, capacity)

	res = fmt.Sprintf(message["init_success"], len(loker))
	return
}

func lokerStatus() (res string, err error) {

	if len(loker) == 0 {
		err = errors.New(message["loker_slot_empty"])
		return
	}

	tableRows := []string{
		"============================================",
		"No Loker\tTipe Identitas\tNo Identitas",
		"--------------------------------------------",
	}

	for i, card := range loker {
		if card != (idCard{}) {
			tableRows = append(tableRows, fmt.Sprintf("%d\t\t%s\t\t%d", i+1, card.Type, card.No))
		}
	}

	res = strings.Join(tableRows, "\n")
	return
}

func inputIdCard(cardType string, cardNo int) (res string, err error){

	if len(loker) == 0 {
		err = errors.New(message["loker_slot_empty"])
		return
	}

	// check existing idCard, prevent duplicate idCard Num
	for i, v := range loker {
		if v.No == cardNo {
			err = errors.New(fmt.Sprintf(message["input_duplicate"], i+1))
			return
		}
	}

	cardType = strings.ToUpper(cardType)

	lokerId := 0
	for i, v := range loker {
		if v == (idCard{}) {
			loker[i] = idCard{cardType, cardNo}
			lokerId = i+1
			break
		}
	}

	if lokerId == 0 {
		err = errors.New(message["input_failed"])
		return
	}

	res = fmt.Sprintf(message["input_success"], lokerId)
	return
}

func leaveIdCard(lokerId int) (res string, err error){

	if len(loker) == 0 {
		err = errors.New(message["loker_slot_empty"])
		return
	}

	if lokerId > len(loker) {
		// nomor loker tidak tersedia
		err = errors.New(fmt.Sprintf(message["leave_failed"], lokerId))
		return
	}
	if loker[lokerId-1] == (idCard{}) {
		// nomor loker tersedia tapi kosong
		err = errors.New(fmt.Sprintf(message["leave_empty"], lokerId))
		return
	}

	loker[lokerId-1] = idCard{}

	res = fmt.Sprintf(message["leave_success"], lokerId)
	return
}

func findIdCard(cardNo int) (res string, err error){

	if len(loker) == 0 {
		err = errors.New(message["loker_slot_empty"])
		return
	}

	for i, card := range loker {
		if card.No == cardNo {
			res = fmt.Sprintf(message["find_result"], i+1)
			return
		}
	}
	res = message["find_empty"]
	return
}

func searchIdCard(cardType string) (res string, err error){

	if len(loker) == 0 {
		err = errors.New(message["loker_slot_empty"])
		return
	}

	cardType = strings.ToUpper(cardType)

	var searchResults []string
	for _, card := range loker {
		if card.Type == cardType {
			searchResults = append(searchResults, fmt.Sprintf("%d", card.No))
		}
	}

	if len(searchResults) == 0 {
		err = errors.New(fmt.Sprintf(message["search_empty"], cardType))
		return
	}

	res = fmt.Sprintf("No Identitas: %s", strings.Join(searchResults, ", "))
	return
}
