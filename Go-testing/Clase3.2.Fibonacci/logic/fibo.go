package logic

func Fibonacci(number int) []int {
	fibo := []int{}
	for i := 0; i < number; i++ {
		if i < 2 {
			fibo = append(fibo, i)
		} else {
			fibo = append(fibo, fibo[i-2]+fibo[i-1])
		}
	}
	return fibo
}
