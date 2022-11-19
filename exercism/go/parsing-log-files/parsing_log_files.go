package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
logMask := regexp.MustCompile(`^\[(?:TRC|DBG|INF|WRN|ERR|FTL)\].*$`)

return logMask.MatchString(text)
}

func SplitLogLine(text string) []string {
sepMask := regexp.MustCompile(`<(?:\~|\*|\=|\-)*>`)

return sepMask.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
pswdMask := regexp.MustCompile(`(?i)".*password.*"`)

var matches int 

for _, line := range lines {
	if pswdMask.MatchString(line) {
		matches += 1
	}
}
return matches
}

func RemoveEndOfLineText(text string) string {
	errorMask := regexp.MustCompile(`end-of-line\d*`)

	return errorMask.ReplaceAllString(text,"")
}

func TagWithUserName(lines []string) []string {
	userMask := regexp.MustCompile(`User\s+(\w+).*`)

	var parsedStrings []string 

	for _, line := range lines {
		matches := userMask.FindStringSubmatch(line) 

		if matches == nil {
			parsedStrings = append(parsedStrings, line)
		}	else {
			correctedLine := fmt.Sprintf("[USR] %s %s", matches[1], line)
			parsedStrings = append(parsedStrings, correctedLine)
		}
	}

	return parsedStrings
}
