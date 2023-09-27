package copy_list_with_random_pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
copyRandomList function copy a linked list with random pointer
1. create a map to store the node and its copy
2. loop the linked list, create a copy node and store it in the map
3. loop the linked list again, set the copy node's next and random pointer
*/
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	hashmap := make(map[*Node]*Node)

	// loop the linked list, create a copy node and store it in the map
	for node := head; node != nil; node = node.Next {
		hashmap[node] = &Node{Val: node.Val}
	}

	// loop the linked list again, set the copy node's next and random pointer
	for node := head; node != nil; node = node.Next {
		hashmap[node].Next = hashmap[node.Next]
		hashmap[node].Random = hashmap[node.Random]
	}

	return hashmap[head]
}
