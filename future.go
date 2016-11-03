package future

type X interface{}

type future_Result struct {
	val X
	err error
}

func _Future(fn func() (X, error)) func() (X, error) {
	c := make(chan future_Result, 1)
	go func() {
		defer close(c)
		result, err := fn()
		c <- future_Result{result, err}
	}()

	return func() (X, error) {
		r := <-c
		return r.val, r.err
	}
}
