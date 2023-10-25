package main

func main(){
	res := AddTwoNumbers(
		&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}},
		&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}})

	for {
		if res == nil {
			return
		}
		println(res.Val)

		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lres := &ListNode{}

	pointer := lres
	rem := 0

	for {
		var res ListNode

		tmp := sum(l1, l2, rem)

		rem = calcRemainder(tmp)

		lres.Val = tmp % 10

		if l1.Next != nil || l2.Next != nil {
			lres.Next = &res
		}

		lres = &res

		var end bool
		end, l1, l2 = getNext(l1, l2)
		if end {
			break
		}
	}

	return pointer
}

func getNext(l1 *ListNode, l2 *ListNode) (bool, *ListNode, *ListNode) {
	if l1.Next == nil && l2.Next == nil {
		return true, nil, nil
	}

	if l1.Next == nil {
		l1.Next = &ListNode{0, nil}
	}

	if l2.Next == nil {
		l2.Next = &ListNode{0, nil}
	}

	return false, l1.Next, l2.Next
}

func calcRemainder(value int) int {
	if value >= 10 {
		return value%9 + 1
	}

	return 0
}

func sum(v1 *ListNode, v2 *ListNode, rem int) int {
	println(v1.Val, v2.Val, rem)
	return v1.Val + v2.Val + rem
}

/**

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

9+9+0 = 18 	0
9+9+1 = 19	9	1
9+9+1 = 19	9	1

*/
