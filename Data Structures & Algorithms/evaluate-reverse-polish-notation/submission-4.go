func evalRPN(tokens []string) int {
	valQueue := []int{}
	for _, val := range tokens {
		v, err := strconv.Atoi(val)
		if err == nil {
			valQueue = append(valQueue, v)
		}
		if val == "+" {
			var temp = valQueue[len(valQueue)-1] + valQueue[len(valQueue)-2]
			valQueue = valQueue[:len(valQueue)-2]
			valQueue = append(valQueue, temp)
		} else if val == "-" {
			var temp = valQueue[len(valQueue)-2] - valQueue[len(valQueue)-1]
			valQueue = valQueue[:len(valQueue)-2]
			valQueue = append(valQueue, temp)
		} else if val == "*" {
			var temp = valQueue[len(valQueue)-1] * valQueue[len(valQueue)-2]
			valQueue = valQueue[:len(valQueue)-2]
			valQueue = append(valQueue, temp)
		} else if val == "/" {
			var temp = valQueue[len(valQueue)-2] / valQueue[len(valQueue)-1]
			valQueue = valQueue[:len(valQueue)-2]
			valQueue = append(valQueue, temp)
		}
	}
	return valQueue[0]
}