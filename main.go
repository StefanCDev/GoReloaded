package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// quits if there's an error
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// function compares two strings
func compare(a, b string) int {
	return strings.Compare(a, b)
}

// function converts word to uppercase
func to_upper(str string) string {
	return strings.ToUpper(str)
}

// converts word to lowercase
func to_lower(str string) string {
	return strings.ToLower(str)
}

// gets the first rune of a string
func first_rune(s string) string {
	a := []rune(s)
	return string(a[0])
}

// capitalises the first charactersacter of a word
func capitalise(str string) string {
	return strings.Title(str)
}

// seperate string by spaces and appends to string list
func split_white_spaces(s string) []string {
	return strings.Fields(s)
}

// removes the space before the quotation mark (from a[space]' to a') character by character
func quotes(s string) string {
	str := ""
	var removeSpace bool
	for i, characters := range s {
		if (i == 0 && characters == 39) || (characters == 39 && s[i-1] == ' ') {
			//if characters == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(characters)
				removeSpace = false
			} else {
				str = str + string(characters)
				removeSpace = true
			}
		} else if i > 1 && s[i-2] == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(characters)
			} else {
				str = str + string(characters)
			}
		} else {
			str = str + string(characters)
		}
	}
	return str
}

// function removes words by ranging through the array of strings

func remove_tags(s []string) string {
	str := ""
	for i, word := range s {
		if word == "(cap," || word == "(low," || word == "(up," {
			s[i] = ""
			s[i+1] = ""
		} else if word != "(up)" && word != "(hex)" && word != "(bin)" && word != "(cap)" && word != "(low)" && word != "" {
			if i == 0 {
				str = str + word
			} else {
				str = str + " " + word
			}
		}
	}
	return str
}

// function trims spaces to the right
func remove_spaces(s string) string {
	return strings.TrimRight(s, " ")

}
func main() {
	//ACCESSING FILE
	data, err := os.ReadFile("sample.txt") //reads the file and returns a byte slice which is stored in 'data'
	//error checker
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	check(err) //
	input := string(data)
	result := split_white_spaces(input)
	// runs range loop to modify result
	for i, v := range result {
		// replaces the word before with its decimal version(compare function)
		// if v == "(hex)"
		if v == "(hex)" {
			j, _ := strconv.ParseInt(result[i-1], 16, 64)
			result[i-1] = fmt.Sprint(j)
		}
		// replaces the word before with its decimal version
		if compare(v, "(bin)") == 0 {
			j, _ := strconv.ParseInt(result[i-1], 2, 64)
			result[i-1] = fmt.Sprint(j)
		}
		// converts the word before to lowercase
		if compare(v, "(low)") == 0 {
			result[i-1] = to_lower(result[i-1])
		}
		// converts the number of words before to lowercase
		if compare(v, "(low,") == 0 {
			result[i-1] = to_lower(result[i-1])
			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			check(err)
			for j := 1; j <= nu; j++ {
				result[i-j] = to_lower(result[i-j])
			}
		}
		// converts the word before to uppercase
		if compare(v, "(up)") == 0 {
			result[i-1] = to_upper(result[i-1])
		}
		// converts the number of words before to uppercase
		if compare(v, "(up,") == 0 {
			result[i-1] = to_upper(result[i-1])
			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			check(err)
			for j := 1; j <= nu; j++ {
				result[i-j] = to_upper(result[i-j])
			}
		}
		// capitalises the word before cap
		if compare(v, "(cap)") == 0 {
			result[i-1] = capitalise(result[i-1])
		}
		// capitalises the number of words before
		if compare(v, "(cap,") == 0 {
			result[i-1] = capitalise(result[i-1])
			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			check(err)
			for j := 1; j <= nu; j++ {
				result[i-j] = capitalise(result[i-j])
			}
		}
		// converts 'a' into 'an' when the next word begins with a vowel or 'h'.

		if (compare(v, "a") == 0 && first_rune(result[i+1]) == "a") || (compare(v, "A") == 0 && first_rune(result[i+1]) == "a") {
			result[i] += "n"
		}
		if (compare(v, "a") == 0 && first_rune(result[i+1]) == "e") || (compare(v, "A") == 0 && first_rune(result[i+1]) == "e") {
			result[i] += "n"
		}
		if (compare(v, "a") == 0 && first_rune(result[i+1]) == "i") || (compare(v, "A") == 0 && first_rune(result[i+1]) == "i") {
			result[i] += "n"
		}
		if (compare(v, "a") == 0 && first_rune(result[i+1]) == "o") || (compare(v, "A") == 0 && first_rune(result[i+1]) == "o") {
			result[i] += "n"
		}
		if (compare(v, "a") == 0 && first_rune(result[i+1]) == "u") || (compare(v, "A") == 0 && first_rune(result[i+1]) == "u") {
			result[i] += "n"
		}
		if (compare(v, "a") == 0 && first_rune(result[i+1]) == "h") || (compare(v, "A") == 0 && first_rune(result[i+1]) == "h") {
			result[i] += "n"
		}

	}
	// calls remove_words() and split_white_spaces() and gets a new result variable
	nowordResult := remove_tags(result)
	result2 := split_white_spaces(nowordResult)
	str := ""
	for _, word := range result2 {
		str = str + word + " "
	}
	// remove spaces from string(punctuation function)
	str = remove_spaces(str)
	word := ""
	for i, characters := range str {
		if i == len(str)-1 {
			if characters == '.' || characters == ',' || characters == '!' || characters == '?' || characters == ';' || characters == ':' {
				if str[i-1] == ' ' {
					word = word[:len(word)-1]
					word = word + string(characters)
				} else {
					word = word + string(characters)
				}
			} else {
				word = word + string(characters)
			} // checking characteres from begining up the end of a string
		} else if characters == '.' || characters == ',' || characters == '!' || characters == '?' || characters == ';' || characters == ':' {
			if str[i-1] == ' ' {
				word = word[:len(word)-1]
				word = word + string(characters)
			} else {
				word = word + string(characters)
			} //it adds a space, if it's not there
			if str[i+1] != ' ' && str[i+1] != '.' && str[i+1] != ',' && str[i+1] != '!' && str[i+1] != '?' && str[i+1] != ';' && str[i+1] != ':' {
				word = word + " "
			}
		} else {
			word = word + string(characters)
		}
	}
	// calls quotes() to remove additional spaces
	word = quotes(word)
	output := []byte(word)
	OurData := os.Args[2]
	// WriteFile writes data to the named file, creating it if necessary.
	// If the file does not exist, WriteFile creates it with permissions perm (before umask);
	// otherwise WriteFile truncates it before writing, without changing permissions.
	words := os.WriteFile(OurData, output, 0644)
	check(words)
}
