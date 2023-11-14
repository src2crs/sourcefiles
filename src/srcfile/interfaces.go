package srcfile

// This file defines a set of interfaces that may be used in the package.
// It is not clear if these interfaces will all be used.
// This package is work in progress and this file can be seen as a design document.

// Line is an interface for a single line of text.
type Line interface {
	// Text returns the text of the line.
	String() string

	// IsCode returns true if the line is code.
	IsCode() bool

	// IsComment returns true if the line is a comment.
	IsComment() bool

	// IsTag returns true if the line is a tag that marks a section of code.
	IsTag() bool

	// IsTagBegin returns true if the line is a tag that marks the beginning of a section of code.
	IsTagBegin() bool

	// IsTagEnd returns true if the line is a tag that marks the end of a section of code.
	IsTagEnd() bool

	// Tag returns the tag of the line if it is a tag.
	// Returns an empty string if the line is not a tag.
	Tag() string

	// IsBlank returns true if the line is blank.
	IsBlank() bool
}

// Block is an interface for a block of lines.
type Block interface {
	// Lines returns the lines of the block.
	Lines() []Line

	// String returns the text of the block.
	String() string

	// IsCode returns true if the block is code.
	IsCode() bool

	// IsComment returns true if the block is a comment.
	IsComment() bool

	// IsTag returns true if the block is a tag that marks a section of code.
	IsTag() bool

	// Tag returns the tag of the line if it is a tag.
	// Returns an empty string if the line is not a tag.
	Tag() string
}

// File is an interface for a file.
type File interface {
	// Name returns the name of the file.
	Name() string

	// Lines returns all lines of the file.
	Lines() []Line

	// String returns the text of the file.
	String() string

	// Blocks returns the blocks of the file.
	Blocks() []Block

	// CodeBlocks returns the blocks of the file that are code.
	CodeBlocks() []Block

	// CommentBlocks returns the blocks of the file that are comments.
	CommentBlocks() []Block

	// Tags returns the tags contained in the file.
	Tags() []string

	// TagBlocks returns the blocks of the file that are tagged with the given tags.
	// If no tags are given, all tagged blocks are returned.
	TagBlocks(tags ...string) []Block
}
