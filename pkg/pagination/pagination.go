// pagination is a package responsible for paginating
package pagination

// Result is the result of a paginated query
type Result[C any] struct {
	Items      []C
	NextCursor *int
}
