package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println(IsValidLine("[ERR] A good error here")) // => true
	fmt.Println(IsValidLine("Any old [ERR] text"))      // => false
	fmt.Println(IsValidLine("[BOB] Any old text"))      // => false

	fmt.Println(SplitLogLine("section 1<*>section 2<~~~>section 3")) // => []string{"section 1", "section 2", "section 3"},

	lines := []string{
		`[INF] passWord`, // contains 'password' but not surrounded by quotation marks
		`"passWord"`,     // count this one
		`[INF] User saw error message "Unexpected Error" on page load.`,          // does not contain 'password'
		`[INF] The message "Please reset your password" was ignored by the user`, // count this one
	}
	fmt.Println(CountQuotedPasswords(lines))

	result := TagWithUserName([]string{
		"[WRN] User James123 has exceeded storage space.",
		"[WRN] Host down. User   Michelle4 lost connection.",
		"[INF] Users can login again after 23:00.",
		"[DBG] We need to check that user names are at least 6 chars long.",
	})

	fmt.Println(result)

}

func IsValidLine(text string) bool {
	match := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return match.MatchString(text)
}

func SplitLogLine(text string) []string {
	match := regexp.MustCompile(`<[-=~*]*>`)
	return match.Split(text, -1)
}

func CountQuotedPasswords(lines []string) (count int) {
	match := regexp.MustCompile(`(?i)".*password.*"`)
	for _, line := range lines {
		if match.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	match := regexp.MustCompile(`end-of-line\d*`)
	return match.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) (ret []string) {
	match := regexp.MustCompile(`User\s+(\w+)`)
	for _, line := range lines {
		founds := match.FindStringSubmatch(line)
		if founds != nil {
			line = fmt.Sprintf("[USR] %s %s", founds[1], line)
		}
		ret = append(ret, line)
	}
	return ret
}
