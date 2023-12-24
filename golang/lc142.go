package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fastNode, slowNode := head, head
	for {
		if fastNode == nil || fastNode.Next == nil {
			return nil
		}
		fastNode = fastNode.Next.Next
		slowNode = slowNode.Next
		if fastNode == slowNode {
			break
		}
	}
	fastNode = head
	for slowNode != fastNode {
		slowNode = slowNode.Next
		fastNode = fastNode.Next
	}
	return fastNode
}
