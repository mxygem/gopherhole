package support

import (
	"fmt"
	"strings"

	"github.com/jaysonesmith/gopherhole/board"
)

func CheckBoardDimensions(x, y int, b board.Board) error {
	var errs []string

	xl := len(b)
	if xl != x {
		errs = append(errs, fmt.Sprintf("X dimension is incorrect. expected %d found %d", x, xl))
	}

	if xl > 0 {
		yl := len(b[0])
		if yl != y {
			errs = append(errs, fmt.Sprintf("Y dimension is incorrect. expected %d found %d", y, yl))
		}
	} else {
		errs = append(errs, "Y dimension not checked as X is 0")
	}

	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}

	return nil
}
