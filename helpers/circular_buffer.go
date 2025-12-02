package helpers

type CircularBuffer[T any] struct {
	data         []*T
	size         int
	currentIndex int
}

func NewBuffer[T any]() *CircularBuffer[T] {
	return &CircularBuffer[T]{
		data:         make([]*T, 0),
		size:         0,
		currentIndex: 0,
	}
}

func (b *CircularBuffer[T]) SetData(data []*T) {
	if data == nil {
		b.data = make([]*T, 0)
	} else {
		b.data = data
	}
	b.size = len(b.data)
	b.currentIndex = 0
}

func (b *CircularBuffer[T]) Enqueue(el *T) {
	b.data = append(b.data, el)
	b.size = len(b.data)
	if b.size == 1 {
		b.currentIndex = 0
	}
}

func (b *CircularBuffer[T]) Pop() *T {
	if b.size == 0 {
		return nil
	}
	el := b.data[b.size-1]
	b.data = b.data[:b.size-1]
	b.size--
	if b.size == 0 {
		b.currentIndex = 0
	} else if b.currentIndex >= b.size {
		b.currentIndex = b.size - 1
	}
	return el
}

func (b *CircularBuffer[T]) PopFront() *T {
	if b.size == 0 {
		return nil
	}
	el := b.data[0]
	b.data = b.data[1:]
	b.size--
	if b.size == 0 {
		b.currentIndex = 0
	} else if b.currentIndex >= b.size {
		b.currentIndex = b.size - 1
	}
	return el
}

func (b *CircularBuffer[T]) Get() *T {
	if b.size == 0 {
		return nil
	}
	if b.currentIndex < 0 || b.currentIndex >= b.size {
		b.currentIndex = 0
	}
	return b.data[b.currentIndex]
}

func (b *CircularBuffer[T]) SetIndex(i int) {
	if b.size == 0 {
		b.currentIndex = 0
		return
	}
	i = i % b.size
	if i < 0 {
		i += b.size
	}
	b.currentIndex = i
}

func (b *CircularBuffer[T]) MoveRight(move int) *T {
	if b.size == 0 {
		return nil
	}
	b.SetIndex(b.currentIndex + move)
	return b.Get()
}

func (b *CircularBuffer[T]) Next() *T {
	return b.MoveRight(1)
}

func (b *CircularBuffer[T]) MoveLeft(move int) *T {
	if b.size == 0 {
		return nil
	}
	b.SetIndex(b.currentIndex - move)
	return b.Get()
}

func (b *CircularBuffer[T]) Previous() *T {
	return b.MoveLeft(1)
}

func (b *CircularBuffer[T]) Len() int {
	return b.size
}
