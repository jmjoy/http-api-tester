package model

import "github.com/jmjoy/http-api-tester/app"

var HistoryModel = &historyModel{
	Model: app.NewModel("history"),
	key:   "history",
}

type historyModel struct {
	*app.Model
	key string
}

func (this *historyModel) GetAll() (datas []Data, err error) {
	has, err := this.Model.Get(this.key, &datas)
	if err != nil {
		return
	}

	if !has {
		datas = []Data{}
		return
	}

	return
}

func (this *historyModel) Insert(data Data) (err error) {
	var datas []Data
	datas, err = this.GetAll()
	if err != nil {
		return
	}

	datas = append(datas, data)
	if len(datas) > 50 { // 最多保存50条历史记录
		datas = datas[1:]
	}

	err = this.Model.Put(this.key, datas)
	return
}
