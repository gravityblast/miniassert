package miniassert

import (
  "testing"
  "fmt"
)

type fakeSuite struct {
  lastError string
}

func init() {
  testSuite = &fakeSuite{}
}

func (s *fakeSuite) Errorf(t *testing.T, message string, args ...interface{}) {
  m := fmt.Sprintf(message, args...)
  s.lastError = m
}

func (s fakeSuite) LastErrorMessage() string {
  return s.lastError
}

func (s *fakeSuite) Reset() {
  s.lastError = ""
}

func check(t *testing.T, description, expectedMessage string, f func() ) {
  f()
  t.Log(description)
  if expectedMessage != testSuite.LastErrorMessage() {
    t.Errorf("%s \nExpected: `%s`\ngot: `%s`", description, expectedMessage, testSuite.LastErrorMessage())
  }
}

func TestTrue(t *testing.T) {
  check(t,
    "True() passing a true value",
    "",
    func() {
      True(t, true)
    },
  )
  check(t,
    "True() passing a false value",
    "Expected <false>(bool) to be true",
    func() {
      True(t, false)
    },
  )
  check(t,
    "True() passing a not bool value",
    "Expected <false>(string) to be true",
    func() {
      True(t, "false")
    },
  )
}

func TestEqual(t *testing.T) {
  check(t,
    "Equal() passing 2 different values",
    "Expected <1>(int), got <2>(int)",
    func() {
      Equal(t, 1, 2)
    },
  )
  check(t,
    "Equal() passing 2 equal values",
    "",
    func() {
      Equal(t, 2, 2)
    },
  )
  check(t,
    "Equal() passing 2 different types",
    "Expected <2>(int), got <2>(string)",
    func() {
      Equal(t, 2, "2")
    },
  )
}

func TestNil(t *testing.T) {
  check(t,
    "Nil() passing nil value",
    "",
    func() {
      Nil(t, nil)
    },
  )

  check(t,
    "Nil() passing not nil value",
    "Expected <true>(bool) to be nil",
    func() {
      Nil(t, true)
    },
  )
}
