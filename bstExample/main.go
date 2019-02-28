// Working with a BST (Binary Search Tree)
package main

import (
	"fmt"
	"log"

	"github.com/bsnux/go-cookbook/bst"
)

func main() {
	values := []int{7, 1, 0, 3, 2, 5, 4, 6, 9, 8, 10}
	tree := &bst.Tree{}

	// Inserting values
	for i := 0; i < len(values); i++ {
		err := tree.Insert(values[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	// in-order traversal:  0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	tree.Traverse(tree.Root)
	fmt.Println("---")
	// post-order traversal: 0, 2, 4, 6, 5, 3, 1, 8, 10, 9, 7
	tree.TraversePos(tree.Root)
	fmt.Println("---")
	// pre-order traversal: 7, 1, 0, 3, 2, 5, 4, 6, 9, 8, 10
	tree.TraversePre(tree.Root)
	fmt.Println("---")

	// level-order: 7, 1, 9, 0, 3, 8, 10, 2, 5, 4, 6
	tree.TraverseLevel(tree.Root)

	// Find a value
	log.Println(tree.Exists(11))
	log.Println(tree.Exists(4))
}
