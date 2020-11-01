package series

import "fmt"

func init() {
	fmt.Println(11)
}
func Square(n int) int {
	fmt.Println(22)
	return n * n
}

func GetFibonacciSerie(n int) []int  {
	ret := []int{1,1}
	for i := 2; i < n; i++  {
		ret = append(ret,ret[i-2]+ret[i-1])
	}
	return ret
}