package stream

type T comparable

type EleHandler[T comparable] func(T) T
type FilterHandler[T comparable] func(T) bool

type streamContainer[T comparable] struct {
	list           []T
	filterHandlers []FilterHandler[T]
	eleHanlders    []EleHandler[T]
	actions        []string
}

const (
	MAP    = "MAP"
	FILTER = "FILTER"
)

func Stream[T comparable](list []T) *streamContainer[T] {
	return &streamContainer[T]{list, []FilterHandler[T]{}, []EleHandler[T]{}, []string{}}
}

func (streamRef *streamContainer[T]) Map(handler EleHandler[T]) *streamContainer[T] {
	streamRef.eleHanlders = append(streamRef.eleHanlders, handler)
	streamRef.actions = append(streamRef.actions, MAP)

	return streamRef
}

func (streamRef *streamContainer[T]) Filter(handler func(ele T) bool) *streamContainer[T] {
	streamRef.filterHandlers = append(streamRef.filterHandlers, handler)
	streamRef.actions = append(streamRef.actions, FILTER)

	return streamRef
}

func (streamRef *streamContainer[T]) Collect() []T {
	collect := streamRef.list
	result := []T{}

	for _, ele := range collect {
		exclude := false
		eleHandlers := streamRef.eleHanlders
		filterHandlers := streamRef.filterHandlers

		if len(collect) == 0 {
			break
		}

		for _, action := range streamRef.actions {

			switch action {
			case MAP:
				handler := eleHandlers[0]
				eleHandlers = eleHandlers[1:]

				ele = handler(ele)

			case FILTER:
				handler := filterHandlers[0]
				filterHandlers = filterHandlers[1:]

				exclude = !handler(ele)
			}

			if exclude {
				break
			}
		}

		if !exclude {
			result = append(result, ele)
		}
	}

	return result
}
