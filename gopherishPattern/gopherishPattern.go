package gopherishPattern

import "regexp"

var Dictionary map[*regexp.Regexp]string

func init() {
	Dictionary = make(map[*regexp.Regexp]string, 3)

	var re *regexp.Regexp

	re = regexp.MustCompile(`^((?:[rRtTpPsSdDfFgGhHjJkKlLzZxXcCvVbBnNmM]{1}qu)|(?:[qQrRtTpPsSdDfFgGhHjJkKlLzZxXcCvVbBnNmM]+))(\w+)([^a-zA-Z]*)$`)
	Dictionary[re] = "${2}${1}ogo${4}"

	re = regexp.MustCompile(`^([aAeEiIoOuUyYwW]+\w+[^a-zA-Z]*)$`)
	Dictionary[re] = "g${1}"

	re = regexp.MustCompile(`^(xr\w+[^a-zA-Z]*)$`)
	Dictionary[re] = "ge${1}"
}
