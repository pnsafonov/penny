package unexported

import (
	"fmt"

	"github.com/pnsafonov/penny/generic"
)

type secret generic.Type

func secretInspect(s secret) string {
	return fmt.Sprintf("%#v", s)
}
