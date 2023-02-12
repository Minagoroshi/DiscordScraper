package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"time"
)

// Application Flags
var (
	Token    = flag.String("token", "", "Bot access token")
	Wordlist = flag.String("wordlist", "wordlist.txt", "Wordlist to use")
	Tag      = flag.Int("tag", 1, "Tag to check")
)

func init() {
	flag.Parse()
}

func main() {

	f, err := os.Open(*Wordlist)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer f.Close()

	var words []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file:", err)
	}

	for _, word := range words {
		tag, err := checkTag(word, *Tag, *Token)
		if err != nil {
			log.Fatal("Error checking tag:", err)
		}
		if tag {
			log.Println("Tag available:", word, "#", *Tag)
			// check if output file exists. If not, create it and append to it, otherwise, append to it.
			if _, err := os.Stat("output.txt"); os.IsNotExist(err) {
				err := os.WriteFile("output.txt", []byte(word+"\n"), 0644)
				if err != nil {
					log.Fatal("Error creating file:", err)
				}
			} else {
				f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0600)
				if err != nil {
					log.Fatal("Error opening file:", err)
				}
				if _, err = f.WriteString(word + "\n"); err != nil {
					log.Fatal("Error writing to file:", err)
				}
				f.Close()
			}
		} else {
			log.Println("Tag not available:", word, "#", *Tag)
		}
		time.Sleep(1 * time.Second)

	}
}
