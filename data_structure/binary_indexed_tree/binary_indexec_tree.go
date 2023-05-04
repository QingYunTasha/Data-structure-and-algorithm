package binaryindexedtree

/*
Use: Compute dynamic prefix sum
Complexity:
	Time: O(logN)
	Space: O(N)
*/

type binaryIndexedTree []int

func NewBinaryIndexedTree(length int) binaryIndexedTree {
	return make(binaryIndexedTree, length)
}

// Set A[i] += v
func (t binaryIndexedTree) Update(i, v int) {
	n := len(t)
	for i <= n {
		t[i]
	}
}

// Set A[i] = v
func (t binaryIndexedTree) Set(i, v int) {
	n := len(t)
	for i <= n {
		t[i]
	}
}

// A[0] + ... + A[i]
func (t binaryIndexedTree) QueryPrefixSum(i int) {

}

// B[i] - B[i-1]
func (t binaryIndexedTree) Query(i int) {

}
