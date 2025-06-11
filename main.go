package main

import (
	"algoholism/week2"
	"fmt"
)

func main() {
	// fmt.Println(week1.TwoSum([]int{2, 7, 11, 15}, 9))
	// fmt.Println(week1.MaxNum([]int{3, 7, 2, 9, 4})) // 9
	// fmt.Println(week1.ContainsDuplicate([]int{1, 2, 3, 1}))
	// fmt.Println(week1.ValidAnagram("listen", "silent"))
	// fmt.Println(week1.MoveZeros([]int{0, 1, 0, 3, 12}))
	// fmt.Println(week1.BestBuy([]int{7, 1, 5, 3, 6, 4}))
	// fmt.Println(week1.RemoveDuplicates([]int{1, 1, 2, 2, 3}))
	// fmt.Println(week1.Intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	// fmt.Println(week1.NonRepeati([]int{2, 3, 4, 2, 3, 5, 4}))
	// fmt.Println(week1.TargetSubarraySum([]int{1, 2, 3}, 3))

	root := &week2.TreeNode{Val: 3, Left: &week2.TreeNode{Val: 9}, Right: &week2.TreeNode{Val: 20, Left: &week2.TreeNode{Val: 15}, Right: &week2.TreeNode{Val: 7}}}
	fmt.Println(week2.LevelOrderTraversal(root))
}
