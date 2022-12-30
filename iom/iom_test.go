package iom

import (
	"reflect"
	"testing"
)

func Test_ShuffleStrings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   []string
	}{
		{
			name: "empty",
			in:   []string{},
		},
		{
			name: "one",
			in:   []string{"one"},
		},
		{
			name: "two",
			in:   []string{"one", "two"},
		},
		{
			name: "three",
			in:   []string{"one", "two", "three"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ShuffleStrings(tt.in)
		})
	}
}

func Test_RemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want []string
	}{
		{
			name: "empty",
			in:   []string{},
			want: []string{},
		},
		{
			name: "one",
			in:   []string{"one"},
			want: []string{"one"},
		},
		{
			name: "ten",
			in:   []string{"one", "two", "two", "two", "two", "six", "seven", "eight", "nine", "ten"},
			want: []string{"one", "two", "six", "seven", "eight", "nine", "ten"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, got := RemoveDuplicates(tt.in)
			if len(got) != len(tt.want) {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetFileExtension(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "empty",
			in:   "",
			want: "",
		},
		{
			name: "one",
			in:   "one",
			want: "",
		},
		{
			name: "two",
			in:   "one.two",
			want: ".two",
		},
		{
			name: "three",
			in:   "one.two.three",
			want: ".three",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := GetFileExtension(tt.in); got != tt.want {
				t.Errorf("GetFileExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AppendSuffixToFilename(t *testing.T) {
	tests := []struct {
		name string
		in   string
		s    string
		want string
	}{
		{
			name: "empty",
			in:   "",
			s:    "",
			want: "",
		},
		{
			name: "three",
			in:   "one-two.txt",
			s:    "-four",
			want: "one-two-four.txt",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := AppendSuffixToFilename(tt.in, tt.s); got != tt.want {
				t.Errorf("AppendSuffixToFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SplitByAndPluckIDs(t *testing.T) {
	tests := []struct {
		name  string
		in    []string
		delim string
		ids   []int
		want  [][]string
	}{
		{
			name:  "empty",
			in:    []string{},
			delim: "",
			ids:   []int{},
			want:  [][]string{},
		},
		{
			name:  "one",
			in:    []string{"hello,this,is,a,test"},
			delim: ",",
			ids:   []int{1, 3},
			want:  [][]string{{"this", "a"}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if len(tt.in) == 0 {
				return
			}

			if got := SplitByAndPluckIDs(tt.in, tt.delim, tt.ids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitByAndPluckIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ChunkByLines(t *testing.T) {
	chunks := ChunkByLines([]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}, 3)
	if len(chunks) != 4 {
		t.Errorf("ChunkByLines() = %v, want %v", len(chunks), 4)
	}

	if !reflect.DeepEqual(chunks[0], []string{"one", "two", "three"}) {
		t.Errorf("ChunkByLines() = %v, want %v", chunks[0], []string{"one", "two", "three"})
	}

	if !reflect.DeepEqual(chunks[1], []string{"four", "five", "six"}) {
		t.Errorf("ChunkByLines() = %v, want %v", chunks[1], []string{"four", "five", "six"})
	}

	if !reflect.DeepEqual(chunks[2], []string{"seven", "eight", "nine"}) {
		t.Errorf("ChunkByLines() = %v, want %v", chunks[2], []string{"seven", "eight", "nine"})
	}

	if !reflect.DeepEqual(chunks[3], []string{"ten"}) {
		t.Errorf("ChunkByLines() = %v, want %v", chunks[3], []string{"ten"})
	}
}

func Test_ChunkByLines_EvenSplit(t *testing.T) {
	chunks := ChunkByLines([]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}, 5)
	if len(chunks) != 2 {
		t.Errorf("ChunkByLines() = %v, want %v", len(chunks), 2)
	}
}

func Test_Diff(t *testing.T) {
	base := map[string]struct{}{
		"one":   {},
		"two":   {},
		"three": {},
	}

	lines := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	n, diff := Diff(base, lines)
	if n != 7 {
		t.Errorf("Diff() = %v, want %v", n, 7)
	}

	if diff[0] != "four" {
		t.Errorf("Diff() = %v, want %v", diff[0], "four")
	}
}
