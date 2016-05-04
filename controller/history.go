package controller

import (
	"github.com/jmjoy/http-api-tester/app"
	"github.com/jmjoy/http-api-tester/model"
)

type HistoryController struct {
	*app.Controller
}

// Get: get all history
func (this *HistoryController) Get() error {
	datas, err := model.HistoryModel.GetAll()
	if err != nil {
		return err
	}

	return this.JsonSuccess(datas)
}
