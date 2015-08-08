package main

type PluginController struct {
	*Controller
}

func NewPluginController() interface{} {
	return &PluginController{
		&Controller{},
	}
}

func (this *PluginController) Get() {
	this.RenderJson(200, "", GetConfigJson())
}
