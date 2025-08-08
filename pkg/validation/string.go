package validation

import (
	"fmt"
)

func CharacterLong(key string, length int) error {
	return fmt.Errorf("%s must be at least %d characters long", key, length)
}
