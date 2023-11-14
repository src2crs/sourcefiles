package srcfile

// SourceFile represents a text file that contains source code.
// It implements the File interface.
type SourceFile struct {
	name  string
	lines []Line
}

// NewSourceFile creates a new SourceFile.
func New(name string) *SourceFile {
	return &SourceFile{
		name:  name,
		lines: make([]Line, 0),
	}
}

// Name returns the name of the file.
func (file *SourceFile) Name() string {
	return file.name
}

// Lines returns all lines of the file.
func (file *SourceFile) Lines() []Line {
	return file.lines
}

// String returns the text of the file.
func (file *SourceFile) String() string {
	// TODO
	return ""
}

// Blocks returns the blocks of the file.
func (file *SourceFile) Blocks() []Block {
	// TODO
	return []Block{}
}

// CodeBlocks returns the blocks of the file that are code.
func (file *SourceFile) CodeBlocks() []Block {
	// TODO
	return []Block{}
}

// CommentBlocks returns the blocks of the file that are comments.
func (file *SourceFile) CommentBlocks() []Block {
	// TODO
	return []Block{}
}

// Tags returns the tags contained in the file.
func (file *SourceFile) Tags() []string {
	// TODO
	return []string{}
}

// TagBlocks returns the blocks of the file that are tagged with the given tags.
// If no tags are given, all tagged blocks are returned.
func (file *SourceFile) TagBlocks(tags ...string) []Block {
	// TODO
	return []Block{}
}
