package main

import "fmt"

func main() {
	fmt.Println(levelCount(makeTree(5, []int{4, -1, 4, 1, 1})))
	fmt.Println(levelCount(makeTree(5, []int{-1, 0, 4, 0, 3})))
	fmt.Println(levelCount(makeTree(5, []int{-1, 0, 1, 2, 3})))
}

func makeTree(nodeCount int, data []int) *tree {
	var root *tree

	relationDict := map[int][]*tree{}
	nodes := []*tree{}

	for _, parent := range data {
		nextNode := &tree{}
		nodes = append(nodes, nextNode)

		if parent == -1 {
			root = nextNode
		} else {
			if _, ok := relationDict[parent]; !ok {
				relationDict[parent] = []*tree{}
			}
			relationDict[parent] = append(relationDict[parent], nextNode)
		}
	}

	for index, node := range nodes {
		node.children = relationDict[index]
	}

	return root
}

type tree struct {
	children []*tree
}

func levelCount(node *tree) int {
	if len(node.children) == 0 {
		return 1
	}

	maxLevelCount := 0

	for _, child := range node.children {
		childLevelCount := levelCount(child)
		if childLevelCount > maxLevelCount {
			maxLevelCount = childLevelCount
		}
	}

	return 1 + maxLevelCount
}
