package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/common-nighthawk/go-figure"

	"github.com/spf13/cobra"
)

// Quotes Struct for quote data
type Quotes []struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}

func getJSON() {
	fmt.Print("Loading...\n\n")

	url := fmt.Sprint("https://benswift404.github.io/stoicly-cli/quotes.json")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()

	var quotes Quotes

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	rand.Seed(time.Now().Unix())
	randomInt := rand.Intn(len(record))
	fmt.Print("#####  #####\n\n")
	fmt.Printf("\"%s\" - %s\n\n", quotes[randomInt].Content, quotes[randomInt].Author)
	fmt.Print("#####  #####\n\n")
	fmt.Print("Stoicly CLI was written in Go by Ben Swift\n\n")
}

func getQuote() *cobra.Command {
	return &cobra.Command{
		Use: "quote",
		RunE: func(cmd *cobra.Command, args []string) error {
			getJSON()
			return nil
		},
	}
}

func main() {
	stoiclyBanner := figure.NewFigure("Stoicly CLI", "", true)
	stoiclyBanner.Print()
	cmd := &cobra.Command{
		Use:          "stoicly",
		Short:        "Welcome to the Stoicly CLI!",
		SilenceUsage: true,
	}

	cmd.AddCommand(getQuote())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
