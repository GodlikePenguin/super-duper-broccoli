package utils

import "fmt"

type Box struct {
	Contents string
}

func init() {
	fmt.Println("Running utils init")
}

func (b *Box) SetContents(newContents string) {
	b.Contents = newContents
}

func (b Box) VerboseString() string {
	return fmt.Sprintf("My contents is %s", b.Contents)
}
