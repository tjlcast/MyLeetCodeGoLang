package main

import (
	"fmt"
	"strconv"
)



/**
	desc：	在一个二维数组中，每一行都按照从左到右递的顺序排序，
			每一列都按照从上到下递增的顺序排序。请完成一个函数，
			输入这样的一个二维数组和一个整数，判断数组中是否含有
			该整数。

	tips:	二维数组的创建、指针、切片
 */

func main() {
	nums := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}

	target := 7
	x, y := find_var_in_matrix(nums[:][:], target)
	fmt.Printf("the x: %d y: %d\n", x, y)

	// --- other test ----
	print_matrix(nums)
	nums = add_one_matrix(nums[:][:])
	fmt.Println("--- result ----")
	print_matrix(nums)
}

func find_var_in_matrix(matrix [][]int, target int) (int, int) {
	rows := len(matrix)
	cols := len(matrix[0])

	x, y := 0, cols-1

	for x<rows && y>=0 {
		num := matrix[x][y]
		if target < num {
			y-=1
		} else if target > num {
			x+=1
		} else {
			// equal
			return x, y
		}
	}

	return -1, -1
}

func print_matrix(matrix [][]int)  {
	for _, line := range matrix {
		for _, num := range line {
			fmt.Print(strconv.Itoa(num) + "\t")
		}
		fmt.Println()
	}
}

func add_one_matrix(matrix [][]int) ([][]int) {
	for i, line := range matrix {
		for j, _ := range line {
			matrix[i][j] += 1
		}
	}
	return matrix
}