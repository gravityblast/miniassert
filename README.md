# miniassert

Mini assert testing framework for the Go language

## Usage

    package hello

    import (
      "testing"
      assert "github.com/pilu/miniassert"
    )

    func TestSomething(t *testing.T) {
      var err error
      assert.Equal(t, "Hello World", Hello())
      assert.True(t, true)
      assert.False(t, false)
      assert.Nil(t, err)
    }

Run `go test` as usual.
