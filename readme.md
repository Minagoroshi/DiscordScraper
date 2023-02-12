# Tag Checker

A command-line program that scrapes a wordlist of usernames with a specified tag on Discord. It reads a list of words from a wordlist file and checks the availability of the specified tag for each word.

## Usage

```bash
go run tagchecker.go -token <bot_access_token> -wordlist <path_to_wordlist> -tag <tag_number>
```

## Flags

- `-token`: Discord token (Required) (Nitro required)
- `-wordlist`: Wordlist to use (default: wordlist.txt)
- `-tag`: Tag to check (default: 1)

## Output

The program outputs the words that have the specified tag available to a file named `output.txt`. If the file does not exist, it creates it, otherwise it appends the output to the existing file.

## Dependencies

- bufio
- flag
- log
- os
- time