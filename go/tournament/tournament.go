package tournament

import (
	"io"
)

func Tally(reader io.Reader, buffer io.Writer) error {
	buffer.Write([]byte("Team                           | MP |  W |  D |  L |  P"))

	return nil
}
