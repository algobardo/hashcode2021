package strategies

type Naive struct {

}

func (n *Naive) Apply(input Input) Output {
	//for _, x := range input. {
	//
	//}
	return nil
}

func NewNaiveStrategy() Strategy {
	return &Naive{}
}


