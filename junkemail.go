package junkemail

import (
	"bufio"
	"embed"
	"strings"
)

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
