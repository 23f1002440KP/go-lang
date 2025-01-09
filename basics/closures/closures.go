package main


func adder() func() int {
	sum := 0
	return func() int {
		sum += 1
		return sum
	}
}

func subs() func() int {
	sum := 0
	return func() int {
		sum -= 1
		return sum
	}
}

func main() {
	pos, neg := adder(), subs()
	for i := 0; i < 10; i++ {
		println(
			pos(),
			neg(),
		)
	}

}