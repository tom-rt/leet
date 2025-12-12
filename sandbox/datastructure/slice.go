package datastructure

import "fmt"

func SliceSandbox() {
	var strSlice []string = []string{"hello"}
	var intArr []int = []int{1, 2, 3}

	appendArrPtr(&strSlice, "toto")
	appendArrPtr(&intArr, 4)

	fmt.Println(strSlice)
	fmt.Println(intArr)

	strSlice = appendArr(strSlice, "toto")
	intArr = appendArr(intArr, 5)

	fmt.Println(strSlice)
	fmt.Println(intArr)
}

func appendArrPtr[T IntOrString](arr *[]T, elem T) {
	*arr = append(*arr, elem)
}

func appendArr[T IntOrString](arr []T, elem T) []T {
	arr = append(arr, elem)
	return arr
}
