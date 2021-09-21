package gopherishPattern

import "regexp"

var dictionary map[*regexp.Regexp]string

func GetGopherishDictionary() map[*regexp.Regexp]string {

	if dictionary != nil {
		return dictionary
	}

	dictionary = make(map[*regexp.Regexp]string, 3)

	var re *regexp.Regexp

	re = regexp.MustCompile(`^(?:(x|(?:(?:x[^r]|[^xaeiouyw])[trpsdfghjklzxcvbnm]*)|[trpsdfghjklzxcvbnm]qu)([aeiouyw]+\w*)([^a-zA-Z]*))$`)
	dictionary[re] = "${2}${1}ogo${3}"

	re = regexp.MustCompile(`^([aeiouyw]+\w*[^a-z]*)$`)
	dictionary[re] = "g${1}"

	re = regexp.MustCompile(`^(xr\w+[^a-zA-Z]*)$`)
	dictionary[re] = "ge${1}"

	return dictionary
}
