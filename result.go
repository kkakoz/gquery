package gquery

type execResult struct {
	err       error
	rowEffect int64
}

type result[T any] struct {
	result T
	*execResult
}

func newExecResult(err error, rowEffect int64) *execResult {
	return &execResult{err: err, rowEffect: rowEffect}
}

func newResult[T any](res T, err error, rowEffect int64) *result[T] {
	return &result[T]{result: res, execResult: &execResult{err: err, rowEffect: rowEffect}}
}

func (r *result[T]) Result() (T, error) {
	return r.result, r.err
}

func (r *execResult) Err() error {
	return r.err
}

func (r *execResult) RowEffect() int64 {
	return r.rowEffect
}
