package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var nodeNotFound = errors.New("node not found")

type node struct {
	value  int
	left   *node
	right  *node
	parent *node
	depth  int
}

var tree = &node{
	value: 30,
	left: &node{
		value: 8,
		left: &node{
			value: 3,
		},
		right: &node{
			value: 20,
			left: &node{
				value: 10,
			},
			right: &node{
				value: 29,
			},
		},
	},
	right: &node{
		value: 52,
	},
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)

	return i
}

func findNodeByValue(n *node, value int) (*node, error) {
	if n == nil {
		return nil, nodeNotFound
	}

	if n.value == value {
		return n, nil
	}

	if n.value > value {
		return findNodeByValue(n.left, value)
	} else {
		return findNodeByValue(n.right, value)
	}
}

func getLowestCommonAncestor(tree *node, value1, value2 int) int {
	node1, err := findNodeByValue(tree, value1)
	if err != nil {
		return tree.value
	}

	node2, err := findNodeByValue(tree, value2)
	if err != nil {
		return tree.value
	}

	minDepth := node1.depth
	if node2.depth < node1.depth {
		minDepth = node2.depth
	}

	if minDepth < 1 {
		minDepth = 1
	}

	fmt.Println("MIN DEPTH", minDepth)

	node1Parent, err := getParentNodeAtDepth(node1, node2.depth)
	if err != nil {
		node1Parent = tree
	}

	node2Parent, err := getParentNodeAtDepth(node2, node1.depth)
	if err != nil {
		node2Parent = tree
	}

	if node1Parent == node2 {
		return node2.value
	}

	if node2Parent == node1 {
		return node1.value
	}

	if node1Parent == node2Parent {
		return node1Parent.value
	} else {
		return tree.value
	}
}

func getParentNodeAtDepth(n *node, depth int) (*node, error) {
	if n == nil {
		return nil, nodeNotFound
	}

	if n.depth == depth {
		return n, nil
	}

	return getParentNodeAtDepth(n.parent, depth)
}

func populateHeuristics(n *node, p *node, depth int) {
	n.depth = depth
	n.parent = p

	if n.left != nil {
		populateHeuristics(n.left, n, depth+1)
	}

	if n.right != nil {
		populateHeuristics(n.right, n, depth+1)
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	populateHeuristics(tree, nil, 1)

	for scanner.Scan() {
		// line := scanner.Text()

		// values := strings.Split(string(line), " ")

		// fmt.Println(getLowestCommonAncestor(tree, atoi(values[0]), atoi(values[1])))
	}

	fmt.Println(getLowestCommonAncestor(tree, 8, 52), 30)
	fmt.Println(getLowestCommonAncestor(tree, 3, 29), 8)
	// fmt.Println(getLowestCommonAncestor(tree, 29, 20), 20)
	// fmt.Println(getLowestCommonAncestor(tree, 10, 29) == 20)
	// fmt.Println(getLowestCommonAncestor(tree, 10, 29) == 20)
}
