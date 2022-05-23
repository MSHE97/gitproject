package products

import "fmt"

//"testing"

func ExampleCalculate() {
	fmt.Println(CalculateDeposit(-100, "TJS"))
	fmt.Println(CalculateDeposit(0, "USD"))
	fmt.Println(CalculateDeposit(500_000_00, "TJS"))
	fmt.Println(CalculateDeposit(500_000_00, "USD"))

	// Output:
	// 0 0
	// 0  0
	// 7001400 8501700
	// 1000200 2000400
}

// func TestCalculate(t *testing.T) {
// 	var minWant int64 = 140
// 	var maxWant int64 = 170
// 	minGot, maxGot := Calculate(1000, "TJS")

// 	if minWant != minGot && maxWant != maxGot {
// 		t.Fatalf("\nGOT: min = %v, max = %v\nWANT: min = %v, max = %v", minGot, maxGot, minWant, maxWant)
// 	}
// 	return
// }
