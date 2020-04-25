package fizzbuzz

import "strconv"

type fbNum int

func FizzBuzz() [100]string {
	var fb [100]string

	for i := 0; i < 100; i++ {
		var number = fbNum(i + 1)
		fb[i] = number.convert()
	}

	return fb
}

func (number fbNum) convert() string {
	conversions := []struct {
		divisor int
		response string
	}{
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
	}

	current := strconv.Itoa(int(number))
	for _, conv := range conversions {
		if number.IsDivisibleBy(conv.divisor) {
			current = conv.response
		}
	}
	return current
}

func (number fbNum) IsDivisibleBy(divisor int) bool {
	return number%fbNum(divisor) == 0
}

