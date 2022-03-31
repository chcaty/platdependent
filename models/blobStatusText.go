package main

type BlobStatusText struct {
	hash         string
	key          string
	blobBucket   string `json:ignore`
	blobFileType string `json:ignore`
}
