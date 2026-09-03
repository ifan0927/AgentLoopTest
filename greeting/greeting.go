// Package greeting provides simple greeting messages.
package greeting

// Message returns a greeting for name. When name is empty, it greets a guest.
func Message(name string) string {
	if name == "" {
		name = "guest"
	}

	return "Hello, " + name + "!"
}
