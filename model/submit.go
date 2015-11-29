package model

import (
	"github.com/jmjoy/http-api-tester/bean"
)

type SubmitModel struct {
}

func NewSubmitModel() *SubmitModel {
	return new(SubmitModel)
}

func (this *SubmitModel) Submit(data bean.Data) error {
	return nil
}
