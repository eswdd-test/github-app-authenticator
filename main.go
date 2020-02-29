package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/bradleyfalzon/ghinstallation"
	"github.com/jessevdk/go-flags"
	"log"
	"net/http"
	"os"
)


var appOpts struct {
	InstallationId int64 `long:"installation-id" description:"GitHub App installation id"`
	ApplicationId  int64 `long:"application-id" description:"GitHub App id"`
	PrivateKeyPath string `long:"private-key" description:"Path to private key PEM file"`
}

var parser = flags.NewParser(&appOpts, flags.Default)

func main() {
	_, err := parser.Parse()

	if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
		os.Exit(1)
	}

	if err != nil {
		os.Exit(1)
	}

	if appOpts.InstallationId == 0 {
		argumentMissing("GitHub App installation id")
	}
	if appOpts.ApplicationId == 0 {
		argumentMissing("GitHub App id")
	}
	if appOpts.PrivateKeyPath == "" {
		argumentMissing("GitHub private key file")
	}

	// init github client, required in all modes
	itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, appOpts.ApplicationId, appOpts.InstallationId, appOpts.PrivateKeyPath)
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	token, err := itr.Token(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(token)
}

func argumentMissing(arg string) {
	log.Println(arg + " not specified")
	f := bufio.NewWriter(os.Stdout)
	//noinspection GoUnhandledErrorResult
	defer f.Flush()
	parser.WriteHelp(f)
	os.Exit(1)
}