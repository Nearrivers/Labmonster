package dirhandler

import "fmt"

type GetLabDirsError struct {
	err error
}

func (g *GetLabDirsError) Error() string {
	return fmt.Sprintf("couldn't get every directory inside the lab: %v", g.err)
}

func (g *GetLabDirsError) Unwrap() error {
	return g.err
}
