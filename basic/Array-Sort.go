package main

import "fmt"

// 冒泡排序，复杂度O(n2)
func bubbleSort(arr []int) {
    for i := 0; i < len(arr)-1; i++ {
        for j := 0; j < len(arr)-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    fmt.Println(arr)
}

// 选择排序，时间复杂度O(n2)
func selectSort(arr []int) {
    for i := 0; i < len(arr)-1; i++ {
        for j := i + 1; j < len(arr)-1; j++ {
            if arr[i] > arr[j] {
                arr[j], arr[i] = arr[i], arr[j]
            }
        }
    }
    fmt.Println(arr)
}

// 插入排序，时间复杂度O(n2)
func insertSort(arr []int) {
    for i := 0; i < len(arr)-1; i++ {
        for j := i; j > 0; j-- {
            if arr[j-1] > arr[j] {
                arr[j-1], arr[j] = arr[j], arr[j-1]
            }
        }
    }
    fmt.Println(arr)
}

// 快速排序，时间复杂度O(nlog2n)
// func quickSort() {

// }

func main() {
    alanList := []int{3432, 2, 234, 5756, 67, 787}
    // bubbleSort(alanList)
    selectSort(alanList)
    // insertSort(alanList)
}
