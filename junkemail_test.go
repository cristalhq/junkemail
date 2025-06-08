package junkemail

import (
	"testing"
)

func TestCheck(t *testing.T) {
	testCases := []struct {
		email string
		want  bool
	}{
		{email: "lol@10-minute-mail.com", want: true},
		{email: "10-minute-mail.com", want: true},
		{email: "kek@zoemail.com", want: true},
		{email: "zoemail.com", want: true},
		{email: "mail.com", want: false},
		{email: "gmail.com", want: false},
	}

	for _, tc := range testCases {
		have := Check(tc.email)
		if have != tc.want {
			t.Errorf("email: '%s' have: %v; want: %v", tc.email, have, tc.want)
		}
	}
}

// something from the middle of the junkemails.txt file.
var benchEmail = "KeK@letTHEmeatspam.com"

func BenchmarkCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isJunk := Check(benchEmail)
		if !isJunk {
			b.Fail()
		}
	}
}
