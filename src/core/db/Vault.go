package db

import (
	"log"
	"settings"
	"sync"

	"github.com/hashicorp/vault/api"
)

var once sync.Once
var vault *Vault

type Vault struct {
	*api.Client
}

func VaultConnect() *Vault {
	once.Do(func() {
		log.Print("Connecting to HashiCorp Vault (" + settings.C.Vault.VaultAddr + ")")
		client, err := api.NewClient(&api.Config{
			Address: settings.C.Vault.VaultAddr,
		})

		if err != nil {
			log.Fatal(err)
		}

		client.SetToken(settings.C.Vault.Token)
		vault = &Vault{client}
	})

	return vault
}

func Read(key string) map[string]interface{} {
	c := VaultConnect()
	val, err := c.Logical().Read(key)
	if err != nil {
		log.Panic(err)
		return nil
	}

	if val == nil {
		return nil
	}

	return val.Data
}

func Write(key string, val map[string]interface{}) {
	c := VaultConnect()
	_, err := c.Logical().Write(key, val)
	if err != nil {
		log.Fatal(err)
	}
}

func List(key string) map[string]interface{} {
	c := VaultConnect()
	val, err := c.Logical().List(key)

	if err != nil {
		log.Panic(err)
		return nil
	}

	if val == nil {
		return nil
	}

	return val.Data
}

func Delete(key string) {
	c := VaultConnect()
	_, err := c.Logical().Delete(key)

	if err != nil {
		log.Panic(err)
	}
}
