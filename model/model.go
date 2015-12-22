package model

type UpsertType string

const (
	UPSERT_ADD    UpsertType = "ADD"
	UPSERT_UPDATE UpsertType = "UPDATE"
)

const (
	BOOKMARK_DEFAULT_NAME = "Default"
	PLUGIN_DEFAULT_NAME   = "default"
)
