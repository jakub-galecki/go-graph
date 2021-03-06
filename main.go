package main

import (
	"container/list"
)

func main() {
	var adjList []list.List
	l1 := list.New()
	l1.PushFront(0)
	l1.PushBack(1)
	l1.PushBack(2)
	l2 := list.New()
	l2.PushFront(1)
	l2.PushBack(2)
	l3 := list.New()
	l3.PushFront(2)
	l3.PushBack(0)
	l3.PushBack(3)
	l4 := list.New()
	l4.PushBack(3)
	l4.PushBack(3)
	graph := []list.List{*l1, *l2, *l3, *l4}
	adjList = append(adjList, graph...)

	//printAdjList(&adjList)
	//adjMatrix := adjListToAdjMatrix(&adjList)
	//printAdjMatrix(&adjMatrix)
	//adjList2 := adjMatrixToAdjList(&adjMatrix)
	//printAdjList(&adjList2)
	DFS(&adjList)
}
