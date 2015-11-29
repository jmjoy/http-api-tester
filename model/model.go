package model

type UpsertType string

const (
	UPSERT_ADD    UpsertType = "ADD"
	UPSERT_UPDATE UpsertType = "UPDATE"
)
