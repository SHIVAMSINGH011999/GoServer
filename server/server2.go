package server

import (
	"context"
	"fmt"

	protos "github.com/shivam/translate2/protos/translation2"
)

type Translation struct {
	protos.UnimplementedTranslation2Server
}

func NewTranslation() *Translation {
	return &Translation{}
}

func (t *Translation) Translate2(
	ctx context.Context,
	i *protos.Translation2Input,
) (*protos.Translation2Output, error) {

	tra := &protos.Translation2Output{
		Text: i.Text + "from server 2",
	}
	fmt.Println(i.Text)
	return tra, nil
}
