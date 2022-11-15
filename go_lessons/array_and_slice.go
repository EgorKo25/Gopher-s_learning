package main

import "fmt"

func main() {
	change_slice_elem()
	copy_slice_elem()
	slice_task_8()
}

func change_slice_elem() {
	s := []int{1, 2, 3} // [1 2 3]
	s = s[:len(s)-1]    // [1 2]

	s = append(s, 5) // [1 2 5]
	fmt.Println(s)
	// The append func not append elem to the old slice
	// and create new slice
}

func copy_slice_elem() {
	// For copy slice elem use the "copy" func
	// the len of copy is a min lenth from two slice
	var dest []int
	dest2, dest3 := make([]int, 3), make([]int, 5)
	src := []int{1, 2, 3, 4}
	copy(dest, src)
	copy(dest2, src)
	copy(dest3, src)
	fmt.Println(dest, dest2, dest3, src) // [] [1 2 3] [1 2 3 4 0] [1 2 3 4]
}
func slice_task_8() {
	// create slice with numeric 1-100
	// and reverse it
	MySlice := make([]int, 100)
	for i := 0; i < len(MySlice); i++ {
		MySlice[i] = i + 1

	}
	// [1 - 100]

	NewSlice := MySlice[10:89]

	tmp := len(NewSlice)
	tmp--
	fmt.Print(tmp)

	for i := 0; i < tmp/2; i++ {

		NewSlice[i], NewSlice[tmp] = NewSlice[tmp], NewSlice[i]

		//NewSlice[i] = NewSlice[i] ^ NewSlice[tmp]
		//NewSlice[tmp] = NewSlice[i] ^ NewSlice[tmp]
		//NewSlice[i] = NewSlice[i] ^ NewSlice[tmp]

		MySlice[i+10] = NewSlice[i]
		tmp--
	}

	fmt.Println(MySlice)

}
