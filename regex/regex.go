package regex

import (
	"fmt"
	"regexp"
)

func emailValidation(email string) error {
	re := regexp.MustCompile(`^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
	if !re.MatchString(email) {
		return fmt.Errorf("Please Check Email Format")
	}
	return nil
}
