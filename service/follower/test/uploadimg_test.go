package test

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/spf13/viper"
	"io/ioutil"
	"ops/pkg/model/azureStorage"
	"testing"
)

func TestConf(t *testing.T) {
	t.Run("viper", func(t *testing.T) {
		viper.SetConfigName("conf")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("../conf/")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		fmt.Println(viper.Get("tls_domain"))
	})
}

func TestImg(t *testing.T) {
	t.Run("img", func(t *testing.T) {
		fileName := "/Users/django/Desktop/Django/v2-ea924fbdd7d32b3cd562fd3129674e81_hd.jpg"
		// Read the entire file into a byte slice
		bytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}

		var base64Encoding string

		// Append the base64 encoded output
		base64Encoding += base64.StdEncoding.EncodeToString(bytes)

		// Print the full base64 representation of the image
		fmt.Println(base64Encoding)
		//azureStorage.UploadFile("test.png",data)
		decodeString, err := base64.StdEncoding.DecodeString(base64Encoding)
		if err != nil {
			return
		}
		fmt.Println(len(decodeString))
		azureStorage.UploadFile("zhaoran.png", decodeString)
	})
}

func TestListBlob(t *testing.T) {
	t.Run("img", func(t *testing.T) {
		//create, err := azureStorage.AzureStorage.Container.Create(context.Background(),azblob.Metadata{}, azblob.PublicAccessNone)

		for marker := (azblob.Marker{}); marker.NotDone(); {
			// Get a result segment starting with the blob indicated by the current Marker.
			listBlob, _ := azureStorage.AzureStorage.Container.ListBlobsFlatSegment(ctx, marker, azblob.ListBlobsSegmentOptions{})

			// ListBlobs returns the start of the next segment; you MUST use this to get
			// the next segment (after processing the current result segment).
			marker = listBlob.NextMarker

			// Process the blobs returned in this result segment (if the segment is empty, the loop body won't execute)
			for _, blobInfo := range listBlob.Segment.BlobItems {
				fmt.Print(" Blob name: " + blobInfo.Name + "\n")
			}
		}
	})
}

func TestDeleteBlob(t *testing.T) {
	t.Run("delete", func(t *testing.T) {
		// Cleaning up the quick start by deleting the container and the file created locally
		strings := []string{"agtJYDv7zQsnnSD7jaJ8.png"}
		for _, s := range strings {
			blobURL := azureStorage.AzureStorage.Container.NewBlockBlobURL(s)
			response, err := blobURL.Delete(context.Background(), azblob.DeleteSnapshotsOptionNone, azblob.BlobAccessConditions{})
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(response.Response())
		}
	})
}

func TestSliece(t *testing.T) {
	t.Run("qiepian", func(t *testing.T) {

	})
}
