package tasks

import "fmt"

type TableOfMultiplications struct{
	Table [][]int64
}
// New creates a new instance of TableOfMultiplications with a size of n+1.
// The first row and column are filled with numbers from 0 to n.
// The remaining cells are initialized with 0.
func New(n int) *TableOfMultiplications{
	var table = make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		table[i] = make([]int64, n+1)
		table[i][0] = int64(i)
		table[0][i] = int64(i)
	}
	return &TableOfMultiplications{
		Table: table,
	}
}
// Fill fills the table with multiplication results.
// It starts from the second row and column (index 1) and multiplies the row index with the column index.
// The result is stored in the corresponding cell in the table.
func (t *TableOfMultiplications) Fill(){
	for i := 1; i < len(t.Table); i++ {
        for j := 1; j < len(t.Table); j++ {
            t.Table[i][j] = int64(i * j)
        }
    }
}
// Print prints the table of multiplications in a tabular format.
func (t *TableOfMultiplications) Print(){
	for i := 0; i < len(t.Table); i++ {
        for j := 0; j < len(t.Table); j++ {
            fmt.Print(t.Table[i][j], "\t")
        }
        fmt.Println()
    }
}