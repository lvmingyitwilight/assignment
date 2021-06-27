package biz

type Greeter interface {
	SayHello(name string) (string, error)
}

type defaultGreeter struct {
}

func (g *defaultGreeter) SayHello(name string) (string, error) {
	return "hello," + name, nil
}
