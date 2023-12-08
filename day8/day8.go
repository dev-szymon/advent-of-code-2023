package main

import (
	"bufio"
	"io"
	"strings"
)

type Node struct {
	id    string
	left  *Node
	right *Node
}

func parseMap(file io.Reader) (map[string]*Node, string) {
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	directions := scanner.Text()
	scanner.Scan()

	nodeIdToNode := map[string]*Node{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "=")
		nodeId := strings.TrimSpace(parts[0])
		children := strings.Split(parts[1], ", ")
		leftId := strings.TrimLeft(strings.TrimSpace(children[0]), "(")
		rightId := strings.TrimRight(strings.TrimSpace(children[1]), ")")

		if nodeIdToNode[leftId] == nil {
			nodeIdToNode[leftId] = &Node{id: leftId}
		}
		if nodeIdToNode[rightId] == nil {
			nodeIdToNode[rightId] = &Node{id: rightId}
		}
		if nodeIdToNode[nodeId] == nil {
			nodeIdToNode[nodeId] = &Node{id: nodeId}
		}
		nodeIdToNode[nodeId].left = nodeIdToNode[leftId]
		nodeIdToNode[nodeId].right = nodeIdToNode[rightId]
	}

	return nodeIdToNode, directions

}

func solvePartOne(file io.Reader) int {
	nodeIdToNode, directions := parseMap(file)

	from := "AAA"
	to := "ZZZ"
	currentNode := nodeIdToNode[from]
	steps := 0
	for currentNode.id != to {
		for _, direction := range directions {
			if direction == 'R' {
				currentNode = currentNode.right
			} else {
				currentNode = currentNode.left
			}
			steps++
		}
	}

	return steps
}

func solvePartTwo(file io.Reader) int {
	startSuffix := 'A'
	endSuffix := 'Z'
	nodeIdToNode, directions := parseMap(file)

	startingNodes := []*Node{}
	for id, node := range nodeIdToNode {
		if rune(node.id[len(id)-1]) == startSuffix {
			startingNodes = append(startingNodes, node)
		}
	}

	largestStepsRequired := 0
	nodeIdToSteps := map[string]int{}
	for _, startingNode := range startingNodes {
		steps := 0
		currentNode := startingNode
	WalkLoop:
		for {
			for _, direction := range directions {
				if direction == 'R' {
					currentNode = currentNode.right
				} else {
					currentNode = currentNode.left
				}

				steps++
				if rune(currentNode.id[len(currentNode.id)-1]) == endSuffix && nodeIdToSteps[startingNode.id] == 0 {
					nodeIdToSteps[startingNode.id] = steps
					if steps > largestStepsRequired {
						largestStepsRequired = steps
					}
					break WalkLoop
				}
			}
		}
	}

	lcm := 0
	for {
		lcm += largestStepsRequired
		foundLcm := true
		for _, steps := range nodeIdToSteps {
			if lcm%steps != 0 {
				foundLcm = false
			}
		}
		if foundLcm {
			break
		}
	}

	return lcm
}
