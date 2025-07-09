package junkemail_test

import (
	"fmt"

	"github.com/cristalhq/junkemail"
)

func ExampleIsValidFormat() {
	{
		email := "drop@10-minute-mail.com"
		isValid := junkemail.IsValidFormat(email)
		fmt.Printf("is '%s' valid? - %v\n", email, isValid)
	}
	{
		email := "drop@.com"
		isValid := junkemail.IsValidFormat(email)
		fmt.Printf("is '%s' valid? - %v\n", email, isValid)
	}
	{
		email := "hehe"
		isValid := junkemail.IsValidFormat(email)
		fmt.Printf("is '%s' valid? - %v\n", email, isValid)
	}
	{
		email := ""
		isValid := junkemail.IsValidFormat(email)
		fmt.Printf("is '%s' valid? - %v\n", email, isValid)
	}

	// Output:
	// is 'drop@10-minute-mail.com' valid? - true
	// is 'drop@.com' valid? - false
	// is 'hehe' valid? - false
	// is '' valid? - false
}

func ExampleCheck() {
	email := "drop@10-minute-mail.com"

	isJunk := junkemail.Check(email)
	fmt.Printf("is '%s' junk? - %v", email, isJunk)

	// Output: is 'drop@10-minute-mail.com' junk? - true
}
