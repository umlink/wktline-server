package model

type UploadFileInput struct {
	FilePath string
	FileName string
	Size     int64
	Hash     string
}

type UploadFileOutput struct {
	Name       string
	Type       string
	Url        string
	BucketHash string
}
