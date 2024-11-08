package main

import (
	"fmt"

	"github.com/draffensperger/golp"
)

/*
	We want to minimize
	-1.0 * x1 - 2.0 * x2 + 0.1 * x3 + 3.0 * x4

	With these constraints:
	x1 + x2 ≤ 5
	2x1 − x2 ≥ 0
	x1 + 3x2 ≥ 0
	x3 + x4 ≥ 0.5
	x3 ≥ 1.1

	And x3 is an integer.
*/

func main() {
	lp := golp.NewLP(1, 4)
	lp.AddConstraintSparse([]golp.Entry{{Col: 0, Val: 1.0}, {Col: 1, Val: 1.0}}, golp.LE, 5.0)
	lp.AddConstraintSparse([]golp.Entry{{Col: 0, Val: 2.0}, {Col: 1, Val: -1.0}}, golp.GE, 0.0)
	lp.AddConstraintSparse([]golp.Entry{{Col: 0, Val: 1.0}, {Col: 1, Val: 3.0}}, golp.GE, 0.0)
	lp.AddConstraintSparse([]golp.Entry{{Col: 2, Val: 1.0}, {Col: 3, Val: 1.0}}, golp.GE, 0.5)
	lp.AddConstraintSparse([]golp.Entry{{Col: 2, Val: 1.0}}, golp.GE, 1.1)
	lp.SetObjFn([]float64{-1.0, -2.0, 0.1, 3.0})
	lp.SetInt(2, true)
	lp.Solve()

	fmt.Printf("Objective value: %v\n", lp.Objective())
	vars := lp.Variables()
	fmt.Printf("Variable values:\n")
	for i := 0; i < lp.NumCols(); i++ {
		fmt.Printf("x%v = %v\n", i+1, vars[i])
	}
}
