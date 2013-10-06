package miniassert

import (
  "fmt"
  "runtime"
  "testing"
  "reflect"
)

type suiteInterface interface {
  Errorf(t *testing.T, message string, args ...interface{})
  LastErrorMessage() string
  Reset()
}

type suite struct {
  lastError string
}

func (s *suite) Errorf(t *testing.T, message string, args ...interface{}) {
  s.lastError = fmt.Sprintf(message, args...)
  text := s.lastError
  _, file, line, ok := runtime.Caller(2)
  if ok {
    text = fmt.Sprintf("%s, in %s:%d", text, file, line)
  }
  t.Error(text)
}

func (s *suite) Reset() {
  s.lastError = ""
}

func (s suite) LastErrorMessage() string {
  return s.lastError
}

var testSuite suiteInterface

func init() {
  testSuite = &suite{}
}

func True(t *testing.T, value interface{}) {
  testSuite.Reset()
  if (value != true) {
    testSuite.Errorf(t, "Expected <%v>(%s) to be true", value, reflect.TypeOf(value))
  }
}

func False(t *testing.T, value interface{}) {
  testSuite.Reset()
  if (value != false) {
    testSuite.Errorf(t, "Expected <%v>(%s) to be false", value, reflect.TypeOf(value))
  }
}

func Nil(t *testing.T, value interface{}) {
  testSuite.Reset()
  if (value != nil) {
    testSuite.Errorf(t, "Expected <%v>(%s) to be nil", value, reflect.TypeOf(value))
  }
}

func Type(t *testing.T, expectedType string, value interface{}) {
  testSuite.Reset()
  valueType := fmt.Sprintf("%v", reflect.TypeOf(value))
  if expectedType != valueType {
    testSuite.Errorf(t, "Expected <%v>(%v) to be of type '%s'", value, reflect.TypeOf(value), expectedType)
  }
}

func Equal(t *testing.T, expected, actual interface{}) {
  testSuite.Reset()
  if !reflect.DeepEqual(expected, actual) {
    testSuite.Errorf(t, "Expected <%v>(%s), got <%v>(%s)", expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual))
  }
}
