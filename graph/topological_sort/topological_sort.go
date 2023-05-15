package topologicalsort

/* @tags: sort,graph,schedular */

/*
type node int
type edge int

The representation of graph: map[node][]edge
*/

// T: O(V + E)
// S: O(V)
func TopoSortRemoveVertix(graph map[int][]int) (seq []int) {
	// initialize indegree

	// T: O(E)
	inDeg := map[int]int{}
	for _, edges := range graph {
		for _, edge := range edges {
			if _, ok := inDeg[edge]; !ok {
				inDeg[edge] = 1
			} else {
				inDeg[edge]++
			}
		}
	}

	// T: O(V)
	for node, degree := range inDeg {
		if degree == 0 {
			seq = append(seq, node)
		}
	}

	// T: O(E)
	for i := 0; i < len(seq); i++ {
		for _, edge := range graph[inDeg[i]] {
			inDeg[edge]--
			if inDeg[edge] == 0 {
				seq = append(seq, edge)
			}
		}
	}

	if len(seq) < len(graph) {
		return nil
	}

	return
}

func TopoSortDfs(graph map[int][]int) []int {
	seq := []int{}

	return seq
}
