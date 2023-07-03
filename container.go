package container

// elements should implement this interface to support
// comparison in the container package
type Comparable interface {
	CompareTo(other any) bool
}
