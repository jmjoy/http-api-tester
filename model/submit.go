package model

type SubmitModel struct {
}

func NewSubmitModel() *SubmitModel {
	return new(SubmitModel)
}

func (this *SubmitModel) Submit(data Data) error {
	return nil
}
