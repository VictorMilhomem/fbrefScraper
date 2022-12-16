# FBREF Data Scraper


Scrape all players and teams from fbref.com and export it to
a json file.

# Build

For building this tool you'll need to have [Go](https://go.dev/) installed in your machine.
Recomended: [latest version](https://go.dev/dl/) (development version: go1.19.4)
```
cd fbrefScraper/
cd cmd/
go build
```

# Run

```
First option (this option don't need building the Cli app):
    go run main.go <options>

Second option (build):
    (Windows)
    ./cmd.exe <options>
    (Linux)
    ./cmd <options>
```

# Usage

```
Usage: go run main.go <options>

    options:
        -t     Set the team avaible at fbref.com
        -id    Set the id of the team avaible at fbref.com
        -year  Set the year avaible at fbref.com 
```

# Output

The file output format is Json, and it will appear 
inside the output folder with the team name.
Take a look at the [examples folder](./cmd/examples/)!



# TODO
- [X] Players and statistics structures
- [X] Commandline tool
- [ ] Teams statistics
- [ ] Competition statistics


