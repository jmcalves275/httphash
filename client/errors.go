package client

type EmptyURL struct {
}

func (e EmptyURL) Error() string {
	return "missing urls"
}

type InvalidFlag struct {
}

func (i InvalidFlag) Error() string {
	return "flag parallel must be greater than zero"
}
