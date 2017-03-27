package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"

	"github.com/Financial-Times/publish-carousel/native"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "sha224"
	app.Usage = "Golang sha224"

	app.Action = func(ctx *cli.Context) error {
		uuid := ctx.Args().Get(0)
		collection := ctx.Args().Get(1)

		mongo := native.NewMongoDatabase("localhost:9000", 10000)
		reader := native.NewMongoNativeReader(mongo)

		content, hash, err := reader.Get(collection, uuid)

		if err != nil {
			return err
		}

		c, _ := json.Marshal(content)
		os.Stdout.Write([]byte(hash))
		return nil
	}

	app.Run(os.Args)
}

// Hash hashes the given payload in SHA224 + Hex
func Hash(payload []byte) (string, error) {
	hash := sha256.New224()
	_, err := hash.Write(payload)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
