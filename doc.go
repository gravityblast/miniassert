/*
package Miniassert is a small assert library.

Usage:

    package hello

    import (
      "testing"
      "errors"
      assert "github.com/pilu/miniassert"
    )

    type Foo struct {}

    func Hello() string { return "Hello World" }

    func TestSomething(t *testing.T) {
      assert.Equal(t, "Hello World", Hello())
      assert.True(t, true)
      assert.False(t, false)

      var err error
      assert.Nil(t, err)

      err = errors.New("foo")
      assert.NotNil(t, err)

      foo := &Foo{}
      assert.Type(t, "*hello.Foo", foo)
    }

Run `go test` as usual.
*/
package miniassert
