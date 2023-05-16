package gquery

type Expr struct {
	column any
	op     string
	val    []any
}

func NewExpr(column Column, op string, val ...any) *Expr {
	return &Expr{column: column, op: op, val: val}
}

func (e *Expr) Expr() string {
	return ""
}

func Or(exs ...*Expr) *Expr {
	return &Expr{}
}

func EQ(c1 Column, c2 any) *Expr {
	switch c2.(type) {
	case Column:
		return NewExpr(c1, opEQ, c2)
	default:
		return NewExpr(c1, opEQ, c2)
	}
}

func On(exs ...*Expr) *Expr {
	return &Expr{}
}
