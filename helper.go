package helper

func If[T any](cond bool, rtrue T, rfalse T) T {
	if cond {
		return rtrue
	}
	return rfalse
}
