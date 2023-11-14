package srcfile

import "testing"

// TestSourceFile_is_File tests if SourceFile implements the File interface.
func TestSourceFile_is_File(t *testing.T) {
	sourcefile := New("test")

	if !isInterface[File](sourcefile) {
		t.Error("SourceFile does not implement the File interface")
	}
}

func isInterface[t any](x any) bool {
	_, ok := x.(t)
	return ok
}

// TODO: Add more Tests
