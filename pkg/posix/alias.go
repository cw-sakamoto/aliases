package posix

import (
	"fmt"
	"strings"
)

// see: http://pubs.opengroup.org/onlinepubs/9699919799/utilities/alias.html
func Alias(name string, str string) *Cmd {
	cmd := Command("alias")

	if name != "" && str != "" {
		cmd.Args = append(cmd.Args, fmt.Sprintf("%s='%s'", name, strings.Replace(str, "'", "\\'", -1)))
	} else if name != "" {
		cmd.Args = append(cmd.Args, name)
	}

	return cmd
}
