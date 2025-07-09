package junkemail

import (
	"bufio"
	"embed"
	"regexp"
	"strings"
)

var emailRe = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// IsValidFormat reports whether email has a valid-ish format.
func IsValidFormat(email string) bool {
	switch idx := strings.IndexByte(email, '@'); {
	case email == "":
		return false
	case idx == -1 || idx > 64 || len(email) > 254:
		return false
	case !emailRe.MatchString(email):
		return false
	default:
		return true
	}
}

// Check if a given email is from a junk domain.
func Check(email string) bool {
	idx := strings.IndexByte(email, '@')
	domain := email[idx+1:]
	domain = strings.ToLower(domain)
	_, ok := domainMap[domain]
	return ok
}

//go:embed junkemails.txt
var staticFS embed.FS

var domainMap = make(map[string]struct{}, 4500)

func init() {
	f, err := staticFS.Open("junkemails.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for scanner := bufio.NewScanner(f); scanner.Scan(); {
		domainMap[scanner.Text()] = struct{}{}
	}
}
