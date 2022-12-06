package main

import (
	"fmt"
)

func do(p1 []int, p2 []int, p3 []int, p4 []int) bool {

	edge1 := (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
	edge2 := (p2[0]-p3[0])*(p2[0]-p3[0]) + (p2[1]-p3[1])*(p2[1]-p3[1])
	edge3 := (p1[0]-p3[0])*(p1[0]-p3[0]) + (p1[1]-p3[1])*(p1[1]-p3[1])
	edge_p4, bol := make([]int, 0, 3), true

	fmt.Println(edge1, edge2, edge3)

	if edge1 != edge2 && edge2 != edge3 && edge1 != edge3 {
		return false
	}

	if !(edge1+edge2 == edge3 || edge1+edge3 == edge2 || edge2+edge3 == edge1) {
		return false
	}

	edge_p4 = append(edge_p4, (p1[0]-p4[0])*(p1[0]-p4[0])+(p1[1]-p4[1])*(p1[1]-p4[1]))
	edge_p4 = append(edge_p4, (p2[0]-p4[0])*(p2[0]-p4[0])+(p2[1]-p4[1])*(p2[1]-p4[1]))
	edge_p4 = append(edge_p4, (p3[0]-p4[0])*(p3[0]-p4[0])+(p3[1]-p4[1])*(p3[1]-p4[1]))
	for i := 0; i < len(edge_p4); i++ {
		if edge_p4[i] == edge1 || edge_p4[i] == edge2 || edge_p4[i] == edge3 {
			bol = bol && true
		} else {
			bol = bol && false
		}
	}
	return bol
}
func main() {
	fmt.Println(do([]int{0, 0}, []int{1, 1}, []int{0, 0}, []int{0, 0}))
}
