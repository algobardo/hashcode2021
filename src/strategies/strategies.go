package strategies

type Input interface {

}

type Output interface {
	ToStrings() []string
}

type Strategy interface {
	Apply(input Input) Output
}
