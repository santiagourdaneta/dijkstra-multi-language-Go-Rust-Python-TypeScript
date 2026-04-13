package main

import (
	"container/heap"
	"fmt"
)

type Edge struct { node string; weight int }
type Item struct { node string; dist int; index int }
type PriorityQueue []*Item

// Implementación de heap.Interface para PriorityQueue...
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i]; pq[i].index = i; pq[j].index = j }
func (pq *PriorityQueue) Push(x interface{}) { item := x.(*Item); item.index = len(*pq); *pq = append(*pq, item) }
func (pq *PriorityQueue) Pop() interface{} { old := *pq; n := len(old); item := old[n-1]; item.index = -1; *pq = old[0 : n-1]; return item }

func Dijkstra(graph map[string][]Edge, start string) map[string]int {
	dist := make(map[string]int)
	for node := range graph { dist[node] = 1e9 }
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, dist: 0})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Item)
		if curr.dist > dist[curr.node] { continue }

		for _, edge := range graph[curr.node] {
			if d := dist[curr.node] + edge.weight; d < dist[edge.node] {
				dist[edge.node] = d
				heap.Push(pq, &Item{node: edge.node, dist: d})
			}
		}
	}
	return dist
}

func main() {
	graph := map[string][]Edge{
		"A": {{node: "B", weight: 4}, {node: "C", weight: 2}},
		"B": {{node: "C", weight: 5}},
		"C": {{node: "B", weight: 1}},
	}
	fmt.Println(Dijkstra(graph, "A"))
}