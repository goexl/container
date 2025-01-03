package kernel

type Queue[T any] interface {
	Collection

	// Enqueue 入队
	Enqueue(T, ...T)

	// Dequeue 出队
	Dequeue() T
}
