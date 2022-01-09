package validation

import (
	"regexp"
)

type Validator struct {

}

var (
	URLRegexp = regexp.MustCompile("([a-z]*:\\/\\/)?[a-zA-Z0-9@:%._\\+~#=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%_\\+.~#?&//=]*)")
	ProtocolRegexp = regexp.MustCompile("[a-z]*:\\/\\/[a-zA-Z0-9@:%._\\+~#=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%_\\+.~#?&//=]*)")
)

func (v *Validator) URLValidation(url string) (string, bool) {
	ok := URLRegexp.Match([]byte(url))
	if !ok {
		return "", false
	}

	withProtocol := ProtocolRegexp.Match([]byte(url))
	if !withProtocol {
		url = "http://" + url
	}

	return url, true
}

