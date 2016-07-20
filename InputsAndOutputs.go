package gotest

type InputsAndOutputs struct {
	MapOfValues map[string]string
}

func NewInputsAndOutputs() *InputsAndOutputs {
	return &InputsAndOutputs{make(map[string]string)}
}

func (inputsAndOutputs *InputsAndOutputs) Add(key string, value string) {
	inputsAndOutputs.MapOfValues[key] = value
}

func (inputsAndOutputs *InputsAndOutputs) Merge(src *InputsAndOutputs) {
	for k, v := range src.MapOfValues {
		inputsAndOutputs.MapOfValues[k] = v
	}
}
