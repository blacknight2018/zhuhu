package configure

import "strings"

func GetCookieValue(cookie string, key string) string {
	pos := strings.Index(cookie, key)
	if -1 != pos {
		ss := []rune(cookie)
		pos += len(key)
		pos++
		i := pos
		for i <= len(cookie)-1 {
			if ss[i] == ';' {
				break
			}
			i++
		}
		if i == len(cookie) || ss[i] == ';' {
			return cookie[pos:i]
		}
	}
	return ""
}
func SetCookieValue(cookie string, key string, value string) string {
	if GetCookieValue(cookie, key) == "" {
		return cookie + ";" + key + "=" + value
	} else {
		src := GetCookieValue(cookie, key)
		return strings.Replace(cookie, src, value, -1)
	}
}
