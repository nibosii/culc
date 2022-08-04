package culc

import (
   "strconv"
)


func Add(x, y int)int{
	return x + y
}

func Sub(x, y int)int{
	return x - y
}

func Mul(x, y int)int{
	return x * y
}

func Div(x, y int)(int, int){
	a, z := x / y, x % y
	return a, z
}

func AddFive(x int)int{
	return x + 5 
}

// 奇数ならtrueを返す関数
func Deep2(n int)bool{
	return n & 1 == 1 
}

func Deep3(s string)int{
	i, err:= strconv.Atoi(s)
	if err != nil{
		return 0
	}
	return i
}
