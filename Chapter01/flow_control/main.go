package main

func main() {
	ten := 10
	if ten == 20 {
		println("This shouldn't be printed as 10 isn't equal to 20")
	} else {
		println("10 is not equal to 20")
	}

	ten_a := 10
	ten_b := 10

	// OR operator is "||", so just one condition needs to be true
	if "a" == "b" || ten_a == ten_b || true == false {
		// Enter here. Although "a" isn't "b" and true isn't false, 10 is equal
		// to 10, so at least one condition is satisfied
		println("10 is equal to 10")

		// AND operator is &&, so BOTH conditions must be satisfied
		// } else if 11 == 11 && "go" == "go" {
		// 	println("This isn't printed, because previous condition was satisfied")
	} else {
		println("In case no condition is satisfied, print this")
	}
}
