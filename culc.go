package culc

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
