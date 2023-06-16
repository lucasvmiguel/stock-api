package pagination

// Result is the result of a paginated query
type Result[C any] struct {
	Items      []C
	NextCursor *uint
}
