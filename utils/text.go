package utils

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
