package gquery

const (
	opEQ    = "="
	opLT    = "<"
	opGT    = ">"
	opLTE   = "<="
	opGTE   = ">="
	opIN    = "IN"
	opExist = "EXIST"
	opAND   = "AND"
	opOR    = "OR"
	opNOT   = "NOT"
	opAdd   = "+"
	opMulti = "*"
)

func C(name string) Column {
	return Column(name)
}

type Column string

func (c Column) EQ(v any) *Expr {
	return NewExpr(c, opEQ, v)
}

func (c Column) LT(v any) *Expr {
	return NewExpr(c, opLT, v)
}

func (c Column) LTE(v any) *Expr {
	return NewExpr(c, opLTE, v)
}

func (c Column) GT(v any) *Expr {
	return NewExpr(c, opGT, v)
}

func (c Column) GTE(v any) *Expr {
	return NewExpr(c, opGTE, v)
}

func (c Column) In(v any) *Expr {
	return NewExpr(c, opIN, v)
}

func (c Column) Exist(v any) *Expr {
	return NewExpr(c, opExist, v)
}

func (c Column) And(v any) *Expr {
	return NewExpr(c, opAND, v)
}

func (c Column) Or(v any) *Expr {
	return NewExpr(c, opOR, v)
}

func (c Column) Not(v any) *Expr {
	return NewExpr(c, opNOT, v)
}

func (c Column) Add(v any) *Expr {
	return NewExpr(c, opAdd, v)
}

func (c Column) Multi(v any) *Expr {
	return NewExpr(c, opMulti, v)
}
