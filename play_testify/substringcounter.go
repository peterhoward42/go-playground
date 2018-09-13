// Package play_testify contains code that explores the use of mocking
// services in the "github.com/stretchr/testify/mock" package.
package play_testify

import (
    "strings"
)


// The function CountSubstrings is the code that we are going to write
// a test for. It depends on a message getter object passed in - which we
// are going to provide in the test using a testify mock object.
func CountSubstrings(getter MessageGetter) int {
    // Returns how many of the first hundred messages retreived from the
    // message getter, contain the substring "01".
    count :=0
    for i := 0; i < 100; i++ {
        message := getter.NextMessage()
        if strings.Contains(message, "01") {
            count += 1
        }
    }
    return count
}


// The MessageGetter interface offers to fetch messages from its source
// one at a time.
type MessageGetter interface {
	NextMessage() string
}
