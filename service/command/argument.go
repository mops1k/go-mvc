package command

type Argument struct {
	value string
}

func (a Argument) Value() string {
	return a.value
}
