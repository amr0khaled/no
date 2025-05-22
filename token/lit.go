package token

import (
	"strconv"
	"strings"
)

// Utils for Literals

func IsInt(s string) bool {
	_, e := strconv.Atoi(s)
	if e == nil {
		return true
	}
	return false
}

func IsString(s string) bool {
	l := len(s)
	a := (s[0] == '"' && s[l-1] == '"') || (s[0] == '\'' && s[l-1] == '\'')
	return a
}

func IsUpperString(s string) bool {
	return strings.ContainsAny(s, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func IsLowerString(s string) bool {
	return strings.ContainsAny(s, "abcdefghijklmnopqrstuvwxyz")
}

func IsFloat(s string) bool {
	_, e := strconv.ParseFloat(s, 32)
	if e == nil {
		return true
	}
	return false
}

func IsBool(s string) bool {
	switch s {
	case "true":
		return true

	case "false":
		return true
	}
	return false
}

func GetValue(s string) interface{} {
	if IsString(s) {
		return string(s)
	} else if IsBool(s) {
		v, _ := strconv.ParseBool(s)
		return v
	} else if IsFloat(s) {
		v, _ := strconv.ParseFloat(s, 64)
		return float64(v)
	} else if IsInt(s) {
		v, _ := strconv.Atoi(s)
		return int(v)
	} else {
		return nil
	}
}
