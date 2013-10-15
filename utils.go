package miniassert

import (
  "reflect"
)

func isNil(value interface{}) bool {
  if value == nil {
    return true
  }

  v := reflect.ValueOf(value)
  if v.Kind() == reflect.Ptr {
    return v.IsNil()
  }

  return false
}

