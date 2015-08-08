package main

type SubmitController struct {
	*Controller
}

func NewSubmitController() interface{} {
	return &SubmitController{
		&Controller{},
	}
}

func (this *SubmitController) Post() {
	this.RenderJson(400, "fuckyou", nil)
}
