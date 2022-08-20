package do

import (
	"testing"
	"unicode"

	"github.com/domcermak/generally/do"
)

func TestMap(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		t.Parallel()

		for _, tc := range []struct {
			name     string
			argument string
			expected string
		}{
			{
				name:     "empty",
				argument: "",
				expected: "",
			},
			{
				name:     "short",
				argument: "string",
				expected: "STRING",
			},
			{
				name:     "long",
				argument: "stringstringstringstringstringstringstringstringstringstring",
				expected: "STRINGSTRINGSTRINGSTRINGSTRINGSTRINGSTRINGSTRINGSTRINGSTRING",
			},
		} {
			t.Run(tc.name, func(t *testing.T) {
				input := []rune(tc.argument)
				uppercase := string(do.Map(input, unicode.ToUpper))
				if uppercase != tc.expected {
					t.Fatalf("expected \"%s\" but got \"%s\"", tc.expected, uppercase)
				}
				if string(input) != tc.argument {
					t.Fatalf("should not have changed from \"%s\" but changed to \"%s\"", tc.argument, string(input))
				}
			})
		}
	})
}
