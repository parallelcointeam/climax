package climax

import (
	"fmt"
	"strings"
)

// Dir returns a string containing all of the tags, and the tags inside tags, to the edges off the tree
func (
	s *KV,
) Dir(
	path string,
) (
	out string,
) {

	//
	return
}

// GetTags returns the Tags item at the given path in the list structure expressed as a space separated list in a string, containing the labels from the KV's structure, and path is interpreted as relative to the node itself so it is possible to cut and paste nodes into each other.
/*
	The primary label of the root KV does not need to be specified in the path, and the root item is meant to be a single KV pair.

	Note that this structure does not have reverse linking as it is declared depth-first and whichever Tags you start from has no back reference (in theory this could be added in a pre-processing pass if it was needed)
*/
func (
	s *KV,
) Get(
	path string,
) (
	out PathVariant,
) {

	var ok bool
	*out.Tags, ok = s.Data.(Tags)
	pathParts := strings.Split(path, " ")
	// Root element is a name and a Tags array
	if ok {
		for _, x := range pathParts {
			for _, y := range *out.Tags {
				if y.Key == x {

					fmt.Println("found", y.Key, y.Data)
					switch cursor := y.Data.(type) {
					case Tags:
						// If we found a match the cursor is moved to the inner Tags item, and this assignment was done in the switch; and now place the newly found Tags item into the return value
						*out.Tags = cursor

					case KV:
						// The path includes or ends with a KV, meaning there is no child Tags below it, meaning there is no Tags matching the path, but instead, a KV pair. Thus return the value here
						*out.KV = y.Data.(KV)
					}
					// Then we advance the position in the path parts string slice
					goto nextTags
				}
			}
			// if we got to here looking for the item in the current Tags it means a path element was not found. The out value may contain the deepest Tags item that *was* found, and the caller may want to use that anyway
			return
		nextTags:
		}
	}
	//
	return
}

// GetKV returns the KV located at the path specified by a space separated string
/*

 */

// InitApp scans the declared KV tree and inserts the back-references to enable relative paths
func (
	s *KV,
) InitApp() (
	out *KV,
) {
	return
}
