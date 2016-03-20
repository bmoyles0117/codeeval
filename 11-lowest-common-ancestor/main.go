package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func getLowestCommonAncestor(node1, node2 *node) int {
	if node1.parent == nil {
		return node1.value
	} else if node2.parent == nil {
		return node2.value
	} else if node1.parent == node2 {
		return node2.value
	} else if node2.parent == node1 {
		return node1.value
	} else if node1.parent == node2.parent {
		return node1.parent.value
	}

	if node1.depth > node2.depth {
		return getLowestCommonAncestor(node1.parent, node2)
	} else {
		return getLowestCommonAncestor(node1, node2.parent)
	}
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
		line := scanner.Text()

		values := strings.Split(string(line), " ")

		node1, err := findNodeByValue(tree, atoi(values[0]))
		if err != nil {
			fmt.Println(tree.value)
			continue
		}

		node2, err := findNodeByValue(tree, atoi(values[1]))
		if err != nil {
			fmt.Println(tree.value)
			continue
		}

		fmt.Println(getLowestCommonAncestor(tree, node1, node2))
	}
}
