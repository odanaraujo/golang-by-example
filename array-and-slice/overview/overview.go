package overview

import "fmt"

func ExampleOverview() {
	//both number of elements and actual elements
	sample := [2]int{1, 2}
	fmt.Printf("Sample: Len: %d %v\n", sample, len(sample))

	//only actual elements
	sample2 := [...]int{1, 2, 3}
	fmt.Printf("Sample2: Len: %d %v\n", sample2, len(sample2))

	//only number of elements
	sample3 := [2]int{}
	fmt.Printf("Sample3: Len: %d %v\n", sample3, len(sample3))

	//whithout both number of elements and actual elements
	sample4 := [...]int{}
	fmt.Printf("Sample4: Len: %d %v\n", sample4, len(sample4))

	sample5 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Sprintln("first run")
	for _, value := range sample5 {
		for _, row := range value {
			fmt.Printf("sample5: len: %d %v\n", row, len(sample5))
		}
	}
}
