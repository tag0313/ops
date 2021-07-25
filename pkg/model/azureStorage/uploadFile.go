package azureStorage

import (
	"context"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/asim/go-micro/v3/util/log"
	"ops/pkg/utils"
)

func UploadFile(fileName string, data []byte) string {
	ctx := context.Background()
	blobURL := AzureStorage.Container.NewBlockBlobURL(fileName)
	_, err := azblob.UploadBufferToBlockBlob(ctx, data, blobURL, azblob.UploadToBlockBlobOptions{
		BlockSize:   4 * 1024 * 1024,
		Parallelism: 16})
	if err != nil {
		log.Error("File upload failed", err)
		return utils.RECODE_UPLOADERR
	}
	return utils.RECODE_OK
}
