# Tag Checker

A command-line program that scrapes a wordlist of usernames with a specified tag on Discord. It reads a list of words from a wordlist file and checks the availability of the specified tag for each word.

## Usage
Using a terminal/command prompt run:
```bash
cd <path_to_discordscraper>
```
Then run:
```
```bash
./DiscordOGUScraper -token <discord_token> -wordlist <path_to_wordlist> -tag <tag_number>
```

## Flags

- `-token`: Discord token (Required) (Nitro required)
- `-wordlist`: Wordlist to use (Optional. default: wordlist.txt)
- `-tag`: Tag to check (Optional. default: 1)
- `-type`: Type of scrape (Optional. default: "user" (user, vanity)

## Output

The program outputs the words that have the specified tag available to a file named `output.txt`. If the file does not exist, it creates it, otherwise it appends the output to the existing file.

## Dependencies

- bufio
- flag
- log
- os
- time
