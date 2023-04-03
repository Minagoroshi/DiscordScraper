# Disclaimer: I made this public because it's too simple of a script to "gatekeep" to myself. This could be made better with proxies, but I want to spend my time working on more important projects, if you wish to improve it please make a pull request and i'll review & merge it.


# Tag Checker

A command-line program that scrapes a wordlist of usernames with a specified tag on Discord. It reads a list of words from a wordlist file and checks the availability of the specified tag for each word.

## Usage
Using a terminal/command prompt run:
```bash
cd <path_to_discordscraper>
```
Then run:
```bash
DiscordOGUScraper [options]
```

## Examples
Vanity URL example:
```bash
DiscordOGUScraper -type vanity
```
User tag example:
```bash
DiscordOGUScraper -token <token> 
```
Custom wordlist example:
```bash
DiscordOGUScraper -token <token> -wordlist <path_to_wordlist>
```
Custom tag example:
```bash
DiscordOGUScraper -token <token> -tag <tag>
```

## Flags

- `-token`: Discord token (Used for tag scraper) (Nitro required)
- `-wordlist`: Wordlist to use (Optional. default: wordlist.txt)
- `-tag`: Tag to check (Optional. default: 1)
- `-type`: Type of scrape (Optional. default: "user" Options: (user, vanity)

## Output

The program outputs the words that have the specified tag available to a file named `output.txt`. If the file does not exist, it creates it, otherwise it appends the output to the existing file.

## Dependencies

- bufio
- flag
- log
- os
- time
