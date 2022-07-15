package main

import (
	"flag"
	"log"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main() {

	flag.Parse()

	if flag.NArg() != 2 {
		log.Println("\nUsage: \n\tali-oss <local file path> <ali oss file path>")
		return 
	}

	args := flag.Args()
	log.Println(args[1], args[0])

  endpoint := os.Getenv("ENDPOINT")
  accessKeyId := os.Getenv("ACCESS_KEY_ID")
  accessKeySecret := os.Getenv("ACCESS_KEY_SECRET")
  bucketName := os.Getenv("BUCKET_NAME")

  client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
  if err != nil {
  	log.Println("oss connect error: ", err)
  }
  
  bucket, err := client.Bucket(bucketName)
  if err != nil {
  	log.Println("oss get bucket error: ", err)
  }
  
  err = bucket.PutObjectFromFile(args[1], args[0])
  if err != nil {
  	log.Println("oss put object file error: ", err)
  }
}


