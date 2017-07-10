package util

import (
	"regexp"
	"mafia-go/common/log"
)

var REG_URL string = "^((http|ftp|https)://)(([a-zA-Z0-9\\._-]+\\.[a-zA-Z]{2,6})|([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}))(:[0-9]{1,4})*(/[a-zA-Z0-9\\&%_\\./-~-]*)?"

func RegUrl(s string) bool {
	isMatch, err := regexp.Match(REG_URL, []byte(s))
	if err != nil {
		log.Error(err)
		return false
	}
	return isMatch
}
