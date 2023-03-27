package utils

import "strings"

func ClearColor(str string) string {
	result := []rune("")
	color := false
	for i := 0; i < len(str); i++ {
		if str[i] == 27 {
			color = true
		} else {
			if color && str[i] == 'm' {
				color = false
			} else if !color {
				result = append(result, rune(str[i]))
			}
		}
	}

	return string(result)
}

func ClearBackspace(str string) string {
	result := []rune("")
	backspace := 0
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == 7 {
			continue
		}

		if str[i] == '\b' {
			backspace++
		} else {
			if backspace > 0 {
				backspace--
			} else {
				result = append(result, rune(str[i]))
			}
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func ConvertCRLFToLf(str string) string {
	return strings.ReplaceAll(str, "\r\n", "\n")
}

func ClearWindowsEscapeKey(str string) string {
	result := []rune("")
	escape := false
	escapeType := 0

	for i := 0; i < len(str); i++ {
		if str[i] == 27 {
			if str[i+1] == '[' {
				escapeType = 1
			} else if str[i+1] == ']' {
				escapeType = 2
			}

			escape = true
		} else {
			if escape {
				if escapeType == 1 && (str[i] == 'X' || str[i] == 'J' || str[i] == 'm' || str[i] == 'H' || str[i] == 'h' || str[i] == 'l') {
					escape = false
				} else if escapeType == 2 && str[i] == 7 {
					escape = false
				}
			} else if !escape {
				result = append(result, rune(str[i]))
			}
		}
	}

	return string(result)
}
