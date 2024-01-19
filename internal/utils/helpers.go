package utils

import "fmt"

func GameKeyWithID(ID int, gameResponse string) string {
	result := fmt.Sprintf("\"game_%d\":%v", ID, gameResponse)
	return result
}
