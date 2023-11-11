package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/lspaccatrosi16/bupload/lib/download"
	"github.com/lspaccatrosi16/bupload/lib/upload"
	"github.com/lspaccatrosi16/bupload/lib/util"
	"github.com/lspaccatrosi16/go-cli-tools/credential"
	"github.com/lspaccatrosi16/go-cli-tools/input"
)

func main() {
	fmt.Println("Provider Credentials:")

	cred, err := credential.GetUserAuth(util.AppName)
	handle(err)

	buf := bytes.NewBuffer(nil)

	fmt.Fprintln(buf, "Key   : "+cred.Key)
	fmt.Fprintln(buf, "Secret: "+strings.Repeat("*", len(cred.Secret)))

	fmt.Println(buf.String())

	opts := []input.SelectOption{
		{
			Name:  "Upload",
			Value: "u",
		},
		{
			Name:  "Download",
			Value: "d",
		},
		{
			Name:  "New Credentials",
			Value: "n",
		},
	}

	v, err := input.GetSelection("Command", opts)

	handle(err)

	switch v {
	case "u":
		err = upload.Upload(cred)
	case "d":
		err = download.Download(cred)
	case "n":
		_, err = credential.RefreshUserCredentials(util.AppName)
	}

	handle(err)

	fmt.Println("Done")
}

func handle(e error) {
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}
}
