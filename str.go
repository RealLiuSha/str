package str

import (
	"strings"
	"unsafe"
)

func StripIndent(multilineStr string) string {
	return strings.Replace(multilineStr, "\t", "", -1)
}

func JoinStrings(strs ...string) string {
	ln := 0
	for i := 0; i < len(strs); i++ {
		ln += len(strs[i])
	}
	bts := make([]byte, ln)
	ln = 0
	for _, str := range strs {
		ln += copy(bts[ln:], str)
	}

	return Bytes2str(bts)
}

func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func JoinBytes(bts ...[]byte) []byte {
	ln := 0
	for i := 0; i < len(bts); i++ {
		ln += len(bts[i])
	}
	ret := make([]byte, ln)
	ln = 0
	for _, b := range bts {
		ln += copy(ret[ln:], b)
	}

	return ret
}

func ReplaceStrings(s string, old []string, replace []string) string {
	if s == "" {
		return s
	}
	if len(old) != len(replace) {
		return s
	}

	for i, v := range old {
		s = strings.Replace(s, v, replace[i], 1000)
	}

	return s
}

func InStringSlice(slice []string, element string) bool {
	element = strings.TrimSpace(element)
	for _, v := range slice {
		if strings.TrimSpace(v) == element {
			return true
		}
	}

	return false
}
