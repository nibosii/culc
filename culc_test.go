package culc

import	(
	"testing"
)

func TestCulc(t *testing.T){
	
	actual := Add(5, 2)
	expect := 7
	if expect != actual{
		t.Error("add error")
	}

	actual = Sub(5, 2)
	expect = 3
	if expect != actual{
		t.Error("sub error")
	}

	actual = Mul(5, 2)
	expect = 10
	if expect != actual{
		t.Error("add error")
	}

	actual , actual1 := Div(5, 2)
	expect   = 2
	expect1 := 1
	if expect != actual{
		t.Error("syou error")
	}

	if expect1 != actual1{
		t.Error("amari error")
	}
}