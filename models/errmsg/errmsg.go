package errmsg

type errmsg struct {
	Code int
	Text string
}

func Create(code int, text string) errmsg {
	n := errmsg{
		Code: code,
		Text: text,
	}
	return n
}
