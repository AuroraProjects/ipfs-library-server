package core

import (
	"bytes"
	shell "github.com/ipfs/go-ipfs-api"
	"io/ioutil"
	"ipfs-library-server/global"
)

func IPFSUpload(s string) (string, error) {
	sh := shell.NewShell(global.CONFIG.IPFS.Api)
	cid, err := sh.Add(bytes.NewBufferString(s))
	if err != nil {
		return "", err
	}
	return cid, nil
}

func IPFSCat(cid string) (string, error) {
	sh := shell.NewShell(global.CONFIG.IPFS.Api)
	read, err := sh.Cat(cid)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(read)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func IPFSDownload(cid string) error {
	sh := shell.NewShell(global.CONFIG.IPFS.Api)
	if err := sh.Get(cid, global.CONFIG.IPFS.OutDir); err != nil {
		return err
	}
	return nil
}
