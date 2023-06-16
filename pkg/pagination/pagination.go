package pagination

// Result is the result of a paginated query
type Result[C any] struct {
	Items      []C   `json:"items"`
	NextCursor *uint `json:"next_cursor"`
}
