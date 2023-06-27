package validator

import "strings"

func GetValidateEmail(val string) bool {
	return strings.HasSuffix(val, "@gmail.com")
}

func GetValidationLength(col string) (int, int) {
	if col == "username" {
		return 6, 36
	} else if col == "email" {
		return 10, 75
	} else if col == "password" {
		return 6, 36
	} else if col == "first_name" {
		return 1, 36
	} else if col == "last_name" {
		return 0, 36
	}
	return 0, 0
}
