package utils

/*
*
	Validate the string is non empty
*/
func ValidateStringInput(input string) bool {
	if len(input) == 0 || input == "" {
		return false
	}
	return true
}
