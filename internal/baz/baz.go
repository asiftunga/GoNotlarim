package baz

import (
    "fmt"

    "newgroupproject/internal/corge"
)

func Baz() string {
    return fmt.Sprintf("Baz %s", corge.Corge())
}