package base

type Model struct {
	db     dbHelper
	bucket string
}

func NewModel(bucket string) *Model {
	return &Model{
		db:     Db,
		bucket: bucket,
	}
}

func (this *Model) Db() dbHelper {
	return this.db
}

func (this *Model) Bucket() string {
	return this.bucket
}
