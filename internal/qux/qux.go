package qux

import (
    "fmt"
	"newgroupproject/internal/corge"
)

func Qux() string {
    return fmt.Sprintf("Qux %s", corge.Corge())
}