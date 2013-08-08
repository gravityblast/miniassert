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

func TestFalse(t *testing.T) {
  check(t,
    "False() passing a false value",
    "",
    func() {
      False(t, false)
    },
  )
  check(t,
    "False() passing a true value",
    "Expected <true>(bool) to be false",
    func() {
      False(t, true)
    },
  )
  check(t,
    "False() passing a not bool value",
    "Expected <false>(string) to be false",
    func() {
      False(t, "false")
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

  a := map[string]string { "foo": "bar" }
  b := map[string]string { "foo": "bar" }
  c := map[string]string { "bar": "baz" }

  check(t,
    "Equal() passing 2 equal maps",
    "",
    func() {
      Equal(t, a, b)
    },
  )

  check(t,
    "Equal() passing 2 different maps",
    "Expected <map[foo:bar]>(map[string]string), got <map[bar:baz]>(map[string]string)",

    func() {
      Equal(t, a, c)
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

func TestType(t *testing.T) {
  check(t,
    "Type() passing 'string' and a string",
    "",
    func() {
      Type(t, "string", "hello")
    },
  )

  check(t,
    "Type() passing 'string' and an int",
    "Expected <1>(int) to be of type 'string'",
    func() {
      Type(t, "string", 1)
    },
  )

  check(t,
    "Type() passing 'string' and an nil",
    "Expected <1>(int) to be of type 'string'",
    func() {
      Type(t, "string", 1)
    },
  )
  check(t,
    "Type() passing 'string' and nil",
    "Expected <<nil>>(<nil>) to be of type 'string'",
    func() {
      Type(t, "string", nil)
    },
  )
}
