package main

import "strconv"

func evalRPN(tokens []string) int {
	res := 0
	stacks := make([]int, len(tokens))
	nowStackTop := -1
	for i := range tokens {
		nowToken := tokens[i]
		if nowToken == "+" || nowToken == "-" || nowToken == "*" || nowToken == "/" {
			nowVal1, nowVal2 := stacks[nowStackTop], stacks[nowStackTop-1]
			nowStackTop -= 2
			switch nowToken {
			case "+":
				res = nowVal1 + nowVal2
			case "-":
				res = nowVal2 - nowVal1
			case "*":
				res = nowVal1 * nowVal2
			case "/":
				res = nowVal2 / nowVal1
			}
			nowStackTop++
			stacks[nowStackTop] = res
		} else {
			nowStackTop++
			stacks[nowStackTop], _ = strconv.Atoi(nowToken)
		}
	}
	if res == -1 { // only number
		return stacks[nowStackTop]
	}
	return res
}
