package gotest

type Givens struct {
	MapOfValues map[string]string
}

func NewGivens() *Givens {
	return &Givens{make(map[string]string)}
}

func (givens *Givens) Add(key string, value string) {
	givens.MapOfValues[key] = value
}

func (givens *Givens) Merge(src *Givens) {
	for k, v := range src.MapOfValues {
		givens.MapOfValues[k] = v
	}
}
