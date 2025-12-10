package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func CreateList(list []int) *ListNode {
	var in *ListNode//Create a nil adress ListNode
	var out *ListNode//Craate a nil adress ListNode
	for i :=len(list); i != 0; i-- {//Cycle 4-1
		out = new(ListNode)//Create a plase for a value
		*out = ListNode{
			Val: list[i-1],//Val get a number in a plase adress
			Next: in,
			}
		in = out//Assign in "in" a value from out
	}
	return out
}

func deleteDuplicates (list *ListNode, a int) *ListNode{
	slise := make([]int, 0, a)//
	for list != nil {//A cycle until the list is nil
		bo := 0//Assign to "bo" the number "0"
		for _, g := range slise {//Cycle from values in the slise. To 'g' assign a value from the slise
			if list.Val == g || bo < 1  {//If the value in list equars a value from slise, or "bo" less them "2".
				bo++//Add to "bo" 1 
			}
		}
		if bo < 2 {//If the value have in the slise, when it skip the value
			slise = append(slise, list.Val)//Add value in the slise
		}
		list = list.Next//Reasigning a other adress for list
	}
	list = CreateList(slise)//Create a new list with values
	return list
}

func main () {
	slise := []int {1, 2, 3, 2, 4, 4}
	list := CreateList(slise)
	fmt.Println(list)
	list = deleteDuplicates(list, len(slise))
	for list != nil {//
		fmt.Println(list)
		list = list.Next
	}
}
