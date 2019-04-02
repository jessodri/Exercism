package bob

import "strings"

func question(remark string) bool {
	trimRemark := strings.TrimSpace(remark)
	question := strings.HasSuffix(trimRemark, "?")
	return question
}

func sayNothing(remark string) bool {
	if strings.TrimSpace(remark) == "" {
		return true
	}
	return false
}

func yelling(remark string) bool {
	if remark == strings.ToUpper(remark) && remark != strings.ToLower(remark){
		return true
	}
	return false
}

// Hey is a function that takes in a remark and returns Bob's answer.
func Hey(remark string) string {

	switch {
	case question(remark) && yelling(remark):
		return "Calm down, I know what I'm doing!"
	case sayNothing(remark):
		return "Fine. Be that way!"
	case question(remark):
		return "Sure."
	case yelling(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
