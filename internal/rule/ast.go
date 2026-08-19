package rule

type Node struct {
	Specs []Spec
}

func Compile(specs []Spec) Node {
	return Node{Specs: CloneSpecs(specs)}
}

func (n Node) Eval(attrs map[string]string) bool {
	return MatchAll(n.Specs, attrs)
}
