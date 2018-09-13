package play_testify

import (
	"testing"

    "github.com/stretchr/testify/mock"
)


func TestCountsCorrectly(t *testing.T) {

    // Prepare a mock object that satisfies the MessageGetter interface.
    getter := MockMessageGetter{}

    // Notify it that it should expect its NextMessage method to be called
    // at least once, with no arguments. And that when so, it should return the
    // fixed string "fibble". This should make the calling-record assertions pass,
    // but the test per-se to fail, because the test
    // expects some of the returned messages among the first 100 read, to contain
    // the substring "01".
    getter.On("NextMessage").Return("fibble")

    // Call the function - passing in the mock message getter.
    count := CountSubstrings(&getter)

    // Make sure it got called as it was instructed to expect.
    getter.AssertExpectations(t)

    // Finish up with a conventional test of the object under test - which should
    // fail because the mock message getter doesn't deliver any messages with
    // "01" in them.
	if count != 43 {
		t.Errorf("Expected 43, got <%d>", count)
	}
}


type MockMessageGetter struct{
  mock.Mock
}


func (m *MockMessageGetter) NextMessage() (string) {
    args := m.Called()
    return args.String(0)
}
