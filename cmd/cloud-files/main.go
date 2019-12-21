package main

import (
	"context"
	"gocloud.dev/blob"
	"log"
	"os"

	// Import the blob packages we want to be able to open.
	_ "gocloud.dev/blob/azureblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
	"io/ioutil"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("usage: cloud-files BUCKET_URL FILE")
	}
	bucketURL := os.Args[1]
	file := os.Args[2]

	ctx := context.Background()
	b, err := blob.OpenBucket(ctx, bucketURL)
	if err != nil {
		log.Fatalf("Failed to setup bucket: %s", err)
	}
	defer b.Close()

	// Prepare the file for upload.
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	w, err := b.NewWriter(ctx, file, nil)
	if err != nil {
		log.Fatalf("Failed to obtain writer: %s", err)
	}

	_, err = w.Write(data)
	if err != nil {
		log.Fatalf("Failed to write bucket: %s", err)
	}

	if err = w.Close(); err != nil {
		log.Fatalf("Failed to close: %s", err)
	}
}