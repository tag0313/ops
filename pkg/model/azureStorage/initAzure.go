package azureStorage

import (
	"fmt"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"net/url"
	"ops/pkg/logger"
	"ops/pkg/utils"
	"os"
	"sync"
)

type Azure struct {
	Container azblob.ContainerURL
}

var (
	AzureStorage *Azure
	once sync.Once
)

func InitAzureStorage() error{
	var err error
	once.Do(func() {
		accountName, accountKey := utils.GetConfigStr("azure.storage.account"), utils.GetConfigStr("azure.storage.key")
		if len(accountName) == 0 || len(accountKey) == 0 {
			logger.Fatal("Either the AZURE_STORAGE_ACCOUNT or AZURE_STORAGE_ACCESS_KEY environment variable is not set")
			os.Exit(0)
		}
		var credential *azblob.SharedKeyCredential
		credential, err = azblob.NewSharedKeyCredential(accountName, accountKey)
		if err != nil {
			logger.Fatal("Invalid credentials with error: " + err.Error())
			return
		}
		p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

		containerName := utils.GetConfigStr("azure.container")
		URL, _ := url.Parse(
			fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName))
		AzureStorage = &Azure{
			Container: azblob.NewContainerURL(*URL, p),
		}
	})
	if err == nil{
		logger.Info("Initialize azure storage successfully")
	}
	return err
}