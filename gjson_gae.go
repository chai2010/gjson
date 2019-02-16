//+build appengine js

package gjson

func getBytes(json []byte, path string) Result {
	return Get(string(json), path)
}
func fillIndex(json string, c *parseContext) {
	// noop. Use zero for the Index value.
}

func stringBytes(s string) []byte {
	return []byte(s)
}
