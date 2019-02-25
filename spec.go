package climax

import (
	"fmt"
	"strings"
)

// Get returns the object at the given path in the list structure expressed as a space separated list in a string, containing the labels from the KV's structure, and path is interpreted as relative to the node itself so it is possible to cut and paste nodes into each other.
/*
	The primary label of the root KV does not need to be specified in the path, and the root item is meant to be a single KV pair.
*/
func (
	s *KV,
) Get(
	path string,
) (
	out *KV,
) {
	pathParts := strings.Split(path, " ")
	// Root element is a name and a Tags array
	rootElem, ok := s.Data.(Tags)
	if ok {
		for _, x := range pathParts {
			for _, y := range rootElem {
				fmt.Println(x, y.Key)
			}
		}
	}
	//
	return
}
