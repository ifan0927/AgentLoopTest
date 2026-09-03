package greeting

import "testing"

func TestMessage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "named greeting",
			input: "Ada",
			want:  "Hello, Ada!",
		},
		{
			name:  "guest fallback",
			input: "",
			want:  "Hello, guest!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Message(tt.input); got != tt.want {
				t.Fatalf("Message(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
