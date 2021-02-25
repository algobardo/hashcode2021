package strategies

import "hashcode2021/m/v2/src/datastructures"

type Output interface {
	ToStrings() []string
}

type Strategy interface {
	Apply(input *datastructures.Input) Output
}
