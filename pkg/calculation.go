package calculation

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")

	if !isValidExpression(expression) {
		return 0, errors.New("некорректное выражение")
	}

	numStack := []float64{}
	opStack := []rune{}

	for i := 0; i < len(expression); i++ {
		ch := rune(expression[i])

		if unicode.IsDigit(ch) || ch == '.' {
			numStr := string(ch)
			for i+1 < len(expression) && (unicode.IsDigit(rune(expression[i+1])) || rune(expression[i+1]) == '.') {
				i++
				numStr += string(expression[i])
			}
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				return 0, err
			}
			numStack = append(numStack, num)
		} else if ch == '(' {
			opStack = append(opStack, ch)
		} else if ch == ')' {
			for len(opStack) > 0 && opStack[len(opStack)-1] != '(' {
				if len(numStack) < 2 {
					return 0, errors.New("некорректное выражение: недостаточно чисел")
				}
				numStack, opStack = evaluate(numStack, opStack)
			}
			if len(opStack) == 0 {
				return 0, errors.New("некорректное выражение: пропущена скобка")
			}
			opStack = opStack[:len(opStack)-1] // Убираем '('
		} else if isOperator(ch) {
			if i == 0 || isOperator(rune(expression[i-1])) || expression[i-1] == '(' {
				return 0, errors.New("некорректное выражение: неожиданная операция")
			}
			for len(opStack) > 0 && precedence(opStack[len(opStack)-1]) >= precedence(ch) {
				if len(numStack) < 2 {
					return 0, errors.New("некорректное выражение: недостаточно чисел")
				}
				numStack, opStack = evaluate(numStack, opStack)
			}
			opStack = append(opStack, ch)
		} else {
			return 0, errors.New("некорректный символ")
		}
	}

	for len(opStack) > 0 {
		if len(numStack) < 2 {
			return 0, errors.New("некорректное выражение: недостаточно чисел")
		}
		numStack, opStack = evaluate(numStack, opStack)
	}

	if len(numStack) != 1 {
		return 0, errors.New("некорректное выражение")
	}

	return numStack[0], nil
}

func isValidExpression(expr string) bool {
	for i, ch := range expr {
		if !isDigit(ch) && !isOperator(ch) && ch != '(' && ch != ')' && ch != '.' {
			return false
		}
		if ch == '(' {
			if i > 0 && isDigit(rune(expr[i-1])) {
				return false
			}
		}
		if ch == ')' {
			if i == 0 || expr[i-1] == '(' {
				return false
			}
		}
	}
	return true
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch) || ch == '.'
}

func isOperator(ch rune) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

func evaluate(numStack []float64, opStack []rune) ([]float64, []rune) {
	b := numStack[len(numStack)-1]
	a := numStack[len(numStack)-2]
	op := opStack[len(opStack)-1]

	numStack = numStack[:len(numStack)-2]
	opStack = opStack[:len(opStack)-1]

	var result float64
	switch op {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		if b == 0 {
			return nil, nil
		}
		result = a / b
	}
	numStack = append(numStack, result)
	return numStack, opStack
}
