package bar

import (
	"fmt"
	"newgroupproject/internal/baz"
	"newgroupproject/internal/qux"
)

func Bar() string{
	return fmt.Sprintf("Bar %s %s",baz.Baz(),qux.Qux())
}