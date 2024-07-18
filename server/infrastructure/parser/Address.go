package parser

import (
	"regexp"
)

func ParseAddress(address string) [3]string {
	// 住所を都道府県、市町村区、それ以降に分割する
	re := regexp.MustCompile(`^(.*?[都道府県])((?:.*?[市町村])?(?:.*?区)?)(.*)$`)

	matches := re.FindStringSubmatch(address)

	if len(matches) != 4 {
		return [3]string{
			"",
			"",
			"",
		}
	}
	return [3]string{
		matches[1],
		matches[2],
		matches[3],
	}
}
