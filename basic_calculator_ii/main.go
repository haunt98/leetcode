package main

func main() {}

func calculate_1(s string) int {
	exprs := parseReversePolishNotation(s)

	result, ok := calcReversePolishNotation(exprs)
	if !ok {
		return 0
	}

	return result
}

type stackString struct {
	arr []string
}

func newStackString() *stackString {
	return &stackString{
		arr: make([]string, 0, 10),
	}
}

func (s *stackString) push(v string) {
	s.arr = append(s.arr, v)
}

func (s *stackString) peek() (string, bool) {
	if len(s.arr) == 0 {
		return "", false
	}

	result := s.arr[len(s.arr)-1]
	return result, true
}

func (s *stackString) pop() (string, bool) {
	if len(s.arr) == 0 {
		return "", false
	}

	result := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return result, true
}

type stackInt struct {
	arr []int
}

func newStackInt() *stackInt {
	return &stackInt{
		arr: make([]int, 0, 10),
	}
}

func (s *stackInt) push(v int) {
	s.arr = append(s.arr, v)
}

func (s *stackInt) pop() (int, bool) {
	if len(s.arr) == 0 {
		return 0, false
	}

	result := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return result, true
}

// https://en.wikipedia.org/wiki/Shunting-yard_algorithm
func parseReversePolishNotation(s string) []string {
	if s == "" {
		return nil
	}

	result := make([]string, 0, 10)

	num := ""
	encounterNum := false

	stack := newStackString()

	for _, c := range s {
		if c == ' ' {
			continue
		}

		if c == '+' || c == '-' || c == '*' || c == '/' {
			if encounterNum {
				result = append(result, num)
				encounterNum = false
				num = ""
			}

			for {
				op, ok := stack.peek()
				if !ok {
					break
				}

				if getPrecendence(op) < getPrecendence(string(c)) {
					break
				}

				stack.pop()
				result = append(result, op)
			}

			stack.push(string(c))
			continue
		}

		if c >= '0' && c <= '9' {
			encounterNum = true
			num += string(c)
			continue
		}
	}

	if encounterNum {
		result = append(result, num)
	}

	for {
		v, ok := stack.pop()
		if !ok {
			break
		}

		result = append(result, v)
	}

	return result
}

func getPrecendence(operator string) int {
	switch operator {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

func atoi(v string) (int, bool) {
	result := 0

	for _, c := range v {
		if c < '0' || c > '9' {
			return 0, false
		}

		result = result*10 + int(c) - int('0')
	}

	return result, true
}

// http://www-stone.ch.cam.ac.uk/documentation/rrf/rpn.html
func calcReversePolishNotation(exprs []string) (int, bool) {
	stack := newStackInt()

	for _, expr := range exprs {
		if expr == "+" || expr == "-" || expr == "*" || expr == "/" {
			second, ok := stack.pop()
			if !ok {
				return 0, false
			}

			first, ok := stack.pop()
			if !ok {
				return 0, false
			}

			switch expr {
			case "+":
				stack.push(first + second)
			case "-":
				stack.push(first - second)
			case "*":
				stack.push(first * second)
			case "/":
				stack.push(first / second)
			default:
				return 0, false
			}

			continue
		}

		num, ok := atoi(expr)
		if !ok {
			return 0, false
		}

		stack.push(num)
	}

	result, ok := stack.pop()
	if !ok {
		return 0, false
	}

	return result, true
}
