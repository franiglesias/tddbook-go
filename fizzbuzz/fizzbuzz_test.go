package fizzbuzz

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		number int
		expected string

	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{6, "Fizz"},
		{5, "Buzz"},
		{10, "Buzz"},
		{15, "FizzBuzz"},
		{16, "16"},
		{30, "FizzBuzz"},
	}
	for _, test := range tests {
		t.Run(string(test.number), func(t *testing.T) {
			fb := FizzBuzz()
			position := test.number - 1
			if fb[position] != test.expected	{
				t.Errorf("FizzBuzz %d should be %s, but found %s", test.number, test.expected, fb[position])
			}
		})
	}
}
