//go:generate go-enum -f=$GOFILE --marshal

package game

// Status indicates the state of the current board
/*
ENUM(
complete
incomplete
error
)
*/
type Status int
