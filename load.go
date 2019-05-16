package godotenv

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"syscall"
)

// feature
// _test := godotenv.GetPanic("TEST_INT")
// _test = "5" || panic("TEST_INT")

// GetWarrning get variable or panic
func GetWarrning(key string) string {
	v, _ := syscall.Getenv(key)
	if v == "" {
		log.Println(key)
	}
	return v
}

// GetPanic get variable or panic
func GetPanic(key string) string {
	v, _ := syscall.Getenv(key)
	if v == "" {
		panic(key)
	}
	return v
}

// int value
//_test, err := godotenv.GetIntPanic("NATS_test")
// _test = 5 || panic("TEST_INT")

// GetInt get int variable
func GetInt(key string) (i int, err error) {
	v, _ := syscall.Getenv(key)
	if v == "" {
		err = errors.New(key)
		return
	}
	return strconv.Atoi(v)

}

// GetIntPanic get int variable or panic
func GetIntPanic(key string) int {
	i, err := GetInt(key)
	if err != nil {
		panic(key)
	}
	return i
}

// Gets returning a slice of substrings
// _test := godotenv.GetsPanic("TEST_INTS")
// _test =["1", "2", "3", "5"] || panic("TEST_INTS")
func Gets(key string) (l []string, err error) {
	v, _ := syscall.Getenv(key)
	if v == "" {
		err = errors.New(key)
		return
	}
	return strings.Fields(v), nil
}

// GetsPanic returning a slice of substrings
func GetsPanic(key string) (l []string) {
	slice, err := Gets(key)
	if err != nil {
		panic(key)
	}
	return slice
}

// _test := godotenv.GetIntsPanic("TEST_INTS")
//  _test = [1, 2, 3, 5] || panic("TEST_INTS")
