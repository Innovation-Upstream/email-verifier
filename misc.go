package emailverifier

import (
	"strings"
)

var (
	disposableDomains       []string        // map to store disposable domains data
	disposableDomainsLoaded bool            //  whether disposableDomains is loaded or not
	freeDomains             map[string]bool // map to store free domains data
	roleAccounts            map[string]bool // map to store role-based accounts data
)

// checks if a string array contains a string
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// IsRoleAccount checks if username is a role-based account
func (v *Verifier) IsRoleAccount(username string) bool {
	return roleAccounts[strings.ToLower(username)]
}

// IsFreeDomain checks if domain is a free domain
func (v *Verifier) IsFreeDomain(domain string) bool {
	return freeDomains[domain]
}

// IsDisposable checks if domain is a disposable domain
func (v *Verifier) IsDisposable(domain string) bool {
	domain = domainToASCII(domain)
	d := parsedDomain(domain)
	found := contains(disposableDomains, d)
	return found
}
