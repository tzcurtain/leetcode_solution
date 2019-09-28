package main

func main() {
	arr := []int{1, 4, 3, 2, 5, 2}
	tmp := new(ListNode)

	tmp = makeLinkList(tmp, arr)
	printLinkList(partition(tmp, 3))
	// fmt.Println(partition(tmp, 3))
}
