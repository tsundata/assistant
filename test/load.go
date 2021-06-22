package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	err = filepath.Walk("./configs", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if ext := filepath.Ext(path); ext != ".yml" {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		name := strings.ReplaceAll(path, ".yml", "")
		name = strings.ReplaceAll(name, "_", "/")
		p := &api.KVPair{Key: fmt.Sprintf("config/%s", name), Value: data}
		_, err = kv.Put(p, nil)
		if err != nil {
			panic(err)
		}

		return err
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
