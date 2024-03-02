package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
  t.Run("multiply a", func(t *testing.T) {
    repeated := Repeat("a", 5)
    expected := "aaaaa"

    if repeated != expected {
      t.Errorf("expected %q but got %q", expected, repeated)
    }
  })

  t.Run("clone string", func(t *testing.T) {
    counted := Repeat("other", 5)
    expected := "not count"
    if counted != expected {
      t.Errorf("expected %q but got %q", expected, counted)
    }
  })

  t.Run("multiply b", func(t *testing.T) {
    repeated := Repeat("b", 6)
    expected := "bbbbbb"

    if repeated != expected {
      t.Errorf("expected %q but got %q", expected, repeated)
    }
  })
}

func BenchmarkRepeat(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Repeat("a", 5)
  }
}

func ExampleRepeat() {
  result := Repeat("b", 3)
  fmt.Println(result)
  // Output: bbb
}

