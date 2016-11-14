//	A little while back we had a programming challenge to convert an infix expression (also known as "normal" math) to a postfix expression
//	(also known as Reverse Polish Notation). Today we'll do something a little different: We will write a calculator that takes RPN input, and outputs the result.
//	Formal input
//
//	The input will be a whitespace-delimited RPN expression. The supported operators will be:
//	+ - addition
//	- - subtraction
//	*, x - multiplication
//	/ - division (floating point, e.g. 3/2=1.5, not 3/2=1)
//	// - integer division (e.g. 3/2=1)
//	% - modulus, or "remainder" division (e.g. 14%3=2 and 21%7=0)
//	^ - power
//	! - factorial (unary operator)
//	Sample input:
//	0.5 1 2 ! * 2 1 ^ + 10 + *
//	Formal output
//
//	The output is a single number: the result of the calculation. The output should also indicate if the input is not a valid RPN expression.
//	Sample output:
//	7
//	Explanation: the sample input translates to 0.5 * ((1 * 2!) + (2 ^ 1) + 10), which comes out to 7.
//	Challenge 1
//
//	Input: 1 2 3 4 ! + - / 100 *
//	Output: -4
//	Challenge 2
//
//	Input: 100 807 3 331 * + 2 2 1 + 2 + * 5 ^ * 23 10 558 * 10 * + + *
//	Finally...
//
//	Hope you enjoyed today's challenge! Have a fun problem or challenge of your own? Drop by /r/dailyprogrammer_ideas and share it with everyone!

package code_kata

func reversePolishNotationCalculator(filePath string) {
	dat, _ := os.Open(filePath)
	scanner := bufio.NewScanner(dat)

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		var stack []float64

		for _, token := range tokens {
			if num, err := strconv.ParseFloat(token, 64); err == nil {
				stack = append(stack, num)
				continue
			}

			switch token {
			case "+", "-", "*", "/", "//", "^":
				var a, b float
				a, stack = pop(stack)
				b, stack = pop(stack)
				if stack == nil {
					break
				}

				switch token {
				case "+":
					stack = append(stack, b+a)
				case "-":
					stack = append(stack, b-a)
				case "*":
					stack = append(stack, b*a)
				case "/":
					stack = append(stack, b/a)
				case "//":
					stack = append(stack, float64(int(b/a)))
				case "%":
					stack = append(stack, float64(int(b)%int(a)))
				case "^":
					stack = append(stack, math.Pow(b, a))
				}
			case "!":
				var n float64
				n, stack = pop(stack)
				if stack == nil {
					break
				}

				a := int(n)
				for i := a - 1; i > 0; i-- {
					a *= i
				}
				stack = append(stack, float64(a))
			default:
				stack = nil
			}
			if stack == nil {
				break
			}
		}

		if len(stack) != 1 {
			fmt.Println("Invalid input!")
		} else if math.Ceil(stack[0] == stack[0]) {
			fmt.Println(int[stack[0]])
		} else {
			fmt.Printf("%f\n", stack[0])
		}

		if err := rd.Err(); err != nil {
			fmt.Println(err)
		}
	}
}

func pop(a []float64) (float64, []float64) {
	if len(a) < 1 {
		return 0, nil
	}

	return a[len(a)-1], a[:len(a)-1]
}
