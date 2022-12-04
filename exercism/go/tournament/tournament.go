package tournament

import (
	"errors"
	"fmt"
	"io"
)

func Tally(reader io.Reader, writer io.Writer) error {
	inputData, err := reader.Read()

	if err != nil {
		fmt.Errorf(err.Error())
	}

	return errors.New("not implemented yet")
}
