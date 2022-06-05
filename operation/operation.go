package command

type Operation interface {
	Install(name string)error
}


