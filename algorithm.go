package main

import (
	"strconv"
	"unicode"
)

func getNumeric(str string) string  {
	var out = ""
	for _, rune := range str{
		if unicode.IsDigit(rune){
			out += string(rune)
		}
	}
	return out
}

func getAllSubstrings(str string, ln int) []string{
	var out []string
	max := len(str)+1
	i := 0
	f := ln
	for f < max {
		out = append(out, str[i:f])
		i += 1
		f += 1
	}
	return out
}

func LuhnAlgorithm(carNo string) bool{

	nDigits := len(carNo)
	nSum := 0
	var mustDuplicate = false

	for i := nDigits-1; i >= 0; i--{
		var d = rune(carNo[i]) - '0'

		if mustDuplicate{
			d = d*2
		}

		nSum += int(d) / 10
		nSum += int(d) % 10

		mustDuplicate = !mustDuplicate
	}

	return nSum % 10 == 0

}

func proofAlgorithm(str string)bool  {
	if len(str) < 12 {
		return  false
	}

	//Se itera en el rango de las posibles tarjetas de crédito
	for i:=12; i <= 19; i++{
		for _, subt := range getAllSubstrings(str, i){
			i, _ := strconv.Atoi(string([]rune(subt)[0]))
			if LuhnAlgorithm(subt) && i >= 3 && i <= 6{
				return true
			}
		}
	}
	return false
}

func isCreditCard(txt string) bool{
	return proofAlgorithm(getNumeric(txt))
}
