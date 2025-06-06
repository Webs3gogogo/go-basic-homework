package main

func isValid(s string) bool {
	stack := []byte{}
	strArr := []byte(s)
	for _, b := range strArr {
		switch currStr := string(b); currStr {
		case "(", "[", "{":

			stack = append(stack, b)
			continue
		case ")":
			if len(stack) > 0 && string(stack[len(stack)-1]) == "(" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case "]":
			if len(stack) > 0 && string(stack[len(stack)-1]) == "[" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case "}":
			if len(stack) > 0 && string((stack[len(stack)-1])) == "{" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
