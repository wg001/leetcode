package main

import "fmt"

func main() {
	//funv1()
	arr_1 := []int{1, 2, 3, 5, 7, 00}
	arr_2 := []int{2, 4, 6,}
	merge(arr_1, 5, arr_2, 3)
}

func funv1() {
	arr_1 := []int{1, 2, 3, 5, 7}
	arr_2 := []int{2, 4, 6,}
	arr_1_len := len(arr_1)
	arr_2_len := len(arr_2)

	var result []int
	i, j := 0, 0

	for {
		if i < arr_1_len && j < arr_2_len {
			if arr_1[i] < arr_2[j] {
				result = append(result, arr_1[i])
				i++
			}
			if arr_1[i] <= arr_2[j] {
				result = append(result, arr_1[i], arr_2[j])
				i++
				j++
			}
			if arr_1[i] > arr_2[j] {
				result = append(result, arr_2[j])
				j++
			}
		} else if i < arr_1_len {
			result = append(result, arr_1[i])
			i++
		} else if j < arr_2_len {
			result = append(result, arr_2[j])
			j++
		} else {
			break
		}
	}
	fmt.Println(result)
}
//func merge(nums1 []int, m int, nums2 []int, n int) {
////	i, j, k := m-1, n-1, m+n-1
////
////	for {
////		if(nums1[i]<)
////	}
////}
