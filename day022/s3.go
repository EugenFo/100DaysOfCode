package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// func to list all buckets within the accessrange of this account
func loadBucketList(sess *session.Session) {
	// creates new session
	sv3 := s3.New(sess)
	result, err := sv3.ListBuckets(nil)
	if err != nil {
		fmt.Println("error!: ", err)
	}
	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}

// creates a new bucket
func createBucket(bucketName string, sess *session.Session) {
	sv3 := s3.New(sess)

	_, err := sv3.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Println("error:", err)
	}

	// waits until bucket exists else throw an error
	err = sv3.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})

	if err != nil {
		fmt.Println("Error while creating bucket:", err)
	}

}

// list all items of a given bucket
func listAllItems(sess *session.Session, bucketName string) {
	sv3 := s3.New(sess)

	resp, err := sv3.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucketName)})
	if err != nil {
		fmt.Println("error:", err)
	}

	for _, items := range resp.Contents {
		fmt.Println("ItemName: ", *items.Key)
		fmt.Println("Size: ", *items.Size)
		fmt.Println("Storage Class: ", *items.StorageClass)
	}
}

func uploadToBucket(sess *session.Session, bucketName string, fileName string, file *os.File) {
	// creates an manager
	uploader := s3manager.NewUploader(sess)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})

	if err != nil {
		fmt.Println("error while uploading:", err)
	}
	fmt.Println("Upload successful!")
}

func downloadingFromBucket(sess *session.Session, bucketName, fileName string, file *os.File) {
	// create an manager
	downloader := s3manager.NewDownloader(sess)
	numBytes, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})

	if err != nil {
		fmt.Println("error while downloading file:", err)
	}

	fmt.Println("Download successful!", file.Name(), numBytes, "Bytes")
}

func main() {
	// reades the credentials file
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: credentials.NewSharedCredentials("", "apiuser2"),
	})

	// throw err if an error while reading the cred occurs
	if err != nil {
		fmt.Println("error: ", err)
	}

	loadBucketList(sess)

	// reading arguments and create the bucket
	/* if len(os.Args) != 2 {
		fmt.Println("Missing Bucketname: ", os.Args[0])
	}
	bucket := os.Args[1]

	createBucket(bucket, sess) */

	// List all Items from a Bucket
	listAllItems(sess, "go-bkt-test")

	// uploading a file to Bucket
	/* 	bucket := os.Args[1]
	   	fileName := os.Args[2]
	   	if len(os.Args) != 3 {
	   		fmt.Println("not enough or too many arguments")
	   	}

	   	file, err := os.Open(fileName)
	   	if err != nil {
	   		fmt.Println("Error while opening the file:", err)
	   	}

	   	// make sure the file will be closed
	   	defer file.Close()

	   	uploadToBucket(sess, bucket, fileName, file) */

	// downloading an item from a bucket

	if len(os.Args) != 3 {
		fmt.Println("not enough or too many arguments")
	}
	bucket := os.Args[1]
	fileName := os.Args[2]

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("error while creating while on fs:", err)
	}

	// make sure the file will be closed
	defer file.Close()

	downloadingFromBucket(sess, bucket, fileName, file)
}
