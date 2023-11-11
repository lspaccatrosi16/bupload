package download

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	gio "io"

	"github.com/lspaccatrosi16/bupload/lib/io"
	"github.com/lspaccatrosi16/bupload/lib/provider"
	"github.com/lspaccatrosi16/go-cli-tools/credential"
)

func Download(cred credential.Credential) error {
	fmt.Println("Download File")
	bucket := io.GetBucket()

	provider, err := provider.GetProvider(cred, bucket)
	if err != nil {
		return err
	}

	object := io.GetObject()

	file, err := provider.GetFile(object)

	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(file)

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	fullPath := filepath.Join(wd, object)

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	defer f.Close()

	gio.Copy(f, buf)

	return nil
}
