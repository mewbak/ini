// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package ini

import (
	"fmt"
	"strconv"
)

// Section represents a single ini section.
// It holds key/value pairs of ini configuration data.
type Section map[string]interface{}

// Set sets the given value for the given key.
func (s Section) Set(key string, value interface{}) {
	s[key] = fmt.Sprint(value)
}

// SetList sets the specified list for the given key.
//
//    s := file.Section("test")
//    s.SetList("foo", "a", "b", "c")
//    s.SetList("bar")
//
//    [test]
//    foo < a
//    foo < b
//    foo < c
//    bar <
//
func (s Section) SetList(key string, argv ...interface{}) {
	l := make([]string, len(argv))

	for i := range argv {
		l[i] = fmt.Sprint(argv[i])
	}

	s[key] = l
}

// S returns the string value for the given key.
func (s Section) S(key, defaultval string) string {
	if v, ok := s[key]; ok {
		if str, ok := v.(string); ok {
			return str
		}
	}
	return defaultval
}

// B returns the boolean value for the given key.
func (s Section) B(key string, defaultval bool) bool {
	if n, err := strconv.ParseBool(s.S(key, "")); err == nil {
		return n
	}
	return defaultval
}

// I returns the int value for the given key.
func (s Section) I(key string, defaultval int) int {
	if n, err := strconv.ParseInt(s.S(key, ""), 0, 32); err == nil {
		return int(n)
	}
	return defaultval
}

// I8 returns the int8 value for the given key.
func (s Section) I8(key string, defaultval int8) int8 {
	if n, err := strconv.ParseInt(s.S(key, ""), 0, 8); err == nil {
		return int8(n)
	}
	return defaultval
}

// I16 returns the int16 value for the given key.
func (s Section) I16(key string, defaultval int16) int16 {
	if n, err := strconv.ParseInt(s.S(key, ""), 0, 16); err == nil {
		return int16(n)
	}
	return defaultval
}

// I32 returns the int32 value for the given key.
func (s Section) I32(key string, defaultval int32) int32 {
	if n, err := strconv.ParseInt(s.S(key, ""), 0, 32); err == nil {
		return int32(n)
	}
	return defaultval
}

// I64 returns the int64 value for the given key.
func (s Section) I64(key string, defaultval int64) int64 {
	if n, err := strconv.ParseInt(s.S(key, ""), 0, 64); err == nil {
		return n
	}
	return defaultval
}

// U returns the uint value for the given key.
func (s Section) U(key string, defaultval uint) uint {
	if n, err := strconv.ParseUint(s.S(key, ""), 0, 32); err == nil {
		return uint(n)
	}
	return defaultval
}

// U8 returns the uint8 value for the given key.
func (s Section) U8(key string, defaultval uint8) uint8 {
	if n, err := strconv.ParseUint(s.S(key, ""), 0, 8); err == nil {
		return uint8(n)
	}
	return defaultval
}

// U16 returns the uint16 value for the given key.
func (s Section) U16(key string, defaultval uint16) uint16 {
	if n, err := strconv.ParseUint(s.S(key, ""), 0, 16); err == nil {
		return uint16(n)
	}
	return defaultval
}

// U32 returns the uint32 value for the given key.
func (s Section) U32(key string, defaultval uint32) uint32 {
	if n, err := strconv.ParseUint(s.S(key, ""), 0, 32); err == nil {
		return uint32(n)
	}
	return defaultval
}

// U64 returns the uint64 value for the given key.
func (s Section) U64(key string, defaultval uint64) uint64 {
	if n, err := strconv.ParseUint(s.S(key, ""), 0, 64); err == nil {
		return n
	}
	return defaultval
}

// F32 returns the float32 value for the given key.
func (s Section) F32(key string, defaultval float32) float32 {
	if n, err := strconv.ParseFloat(s.S(key, ""), 32); err == nil {
		return float32(n)
	}
	return defaultval
}

// F64 returns the float64 value for the given key.
func (s Section) F64(key string, defaultval float64) float64 {
	if n, err := strconv.ParseFloat(s.S(key, ""), 64); err == nil {
		return n
	}
	return defaultval
}

// List returns a list of string values for the given key.
func (s Section) List(key string) []string {
	if v, ok := s[key]; ok {
		if str, ok := v.([]string); ok {
			return str
		}
	}
	return nil
}
