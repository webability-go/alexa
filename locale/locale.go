package locale

import (
	"fmt"
)

type Lang map[string]string

func (l Lang) Get(entry string) string {
	e, ok := l[entry]
	if !ok {
		fmt.Println("Lang entry not found: ", entry)
	}
	return e
}

func Get(locale string) Lang {
	switch locale {
	// natives:
	case "en-US":
		return en_US
	case "es-MX":
		return es_MX
	case "fr-FR":
		return fr_FR

	// others (to translate well as native)
	case "en-AU":
		return en_US
	case "en-CA":
		return en_US
	case "en-GB":
		return en_US
	case "en-IN":
		return en_US
	case "es-ES":
		return es_MX
	case "fr-CA":
		return fr_FR
	case "de-DE":
		return en_US
	case "ja-JP":
		return en_US
	}
	return en_US
}
