package main

import (
	"fmt"

	emailverifier "github.com/innovation-upstream/email-verifier"
)

func main() {
	verifi := emailverifier.NewVerifier()
	domain := "mailinator.com"
	ret := verifi.IsDisposable(domain)
	fmt.Println("misc validation result: ", ret)
}
