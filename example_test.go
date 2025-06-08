package junkemail_test

import (
	"fmt"

	"github.com/cristalhq/junkemail"
)

func ExampleCheck() {
	email := "drop@10-minute-mail.com"

	isJunk := junkemail.Check(email)
	fmt.Printf("is '%s' junk? - %v", email, isJunk)

	// Output: is 'drop@10-minute-mail.com' junk? - true
}
