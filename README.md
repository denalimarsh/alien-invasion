# Invasion
Mad aliens are about to invade earth and we have been tasked with simulating the impending invasion.

## Assumptions
- Input
    - City names are single, capitalized words such as "Seattle", "Berlin", and "Honolulu"
    - Path directions are all correctly spelled and formatted such as "north", "east", and "west"
    - inFile contains clean, digestible data and is always in .txt format
    - valid file path is given as command line argument
- Paths are not automatically bidirectional, such that `boston east=Berlin west=Honolulu` allows an alien to move from Boston to Berlin, but not from Berlin to Boston
- Once destroyed, cities are removed from the world but have their name retained for printing at the end of the game

## Installation
If you have already downloaded go and set your GOPATH, installing invasion is easy
- 1. Clone the repo
- 2. Run `go install` from the root directory

If not, getting set up with go is painless
- [Download go](https://golang.org/dl/)
- [Using GOPATH](https://github.com/golang/go/wiki/GOPATH)

## Usage

The root cmd for invasion is `invasion`:
```
$ invasion
invasion unleashes an army of aliens on an unsuspecting planet

Usage:
  invasion [command]

Available Commands:
  help        Help about any command
  start       starts the extraterrestrial invasion of a world

Flags:
  -h, --help   help for invasion

Use "invasion [command] --help" for more information about a command.
```

Start the game with `invasion start --file="./assets/small_world.txt" --numAliens=10`:
```
$ invasion start --file="./assets/small_world.txt" --numAliens=10
2019/11/26 00:34:16 
2019/11/26 00:34:16                     The World:
2019/11/26 00:34:16 -------------------------------------------------
2019/11/26 00:34:16 Beijing south=Mumbai west=Tokyo 
2019/11/26 00:34:16 Mumbai north=Istanbul west=Cairo east=Beijing 
2019/11/26 00:34:16 Tokyo north=Istanbul south=Delhi 
2019/11/26 00:34:16 Istanbul east=Beijing west=Paris south=Tokyo 
2019/11/26 00:34:16 Cairo north=Mumbai east=Seoul west=Paris 
2019/11/26 00:34:16 Paris south=Cairo 
2019/11/26 00:34:16 Delhi east=Seoul 
2019/11/26 00:34:16 Seoul south=Cairo west=Chicago 
2019/11/26 00:34:16 Chicago east=Beijing 
2019/11/26 00:34:16 Hanoi north=Chicago east=Seoul 
2019/11/26 00:34:16 -------------------------------------------------
2019/11/26 00:34:16 
2019/11/26 00:34:16 The game has started!
2019/11/26 00:34:16 
2019/11/26 00:34:16 Turn 1: Seoul has been destroyed by aliens 1 and 3 
2019/11/26 00:34:16 Turn 1: Chicago has been destroyed by aliens 6 and 8 
2019/11/26 00:34:16 Turn 2: Tokyo has been destroyed by aliens 7 and 5 
2019/11/26 00:34:16 Turn 4: Paris has been destroyed by aliens 10 and 2 
2019/11/26 00:34:16 
2019/11/26 00:34:16 Game completed on turn 10001.
2019/11/26 00:34:16 
2019/11/26 00:34:16                     The World:
2019/11/26 00:34:16 -------------------------------------------------
2019/11/26 00:34:16 Beijing south=Mumbai 
2019/11/26 00:34:16 Mumbai north=Istanbul west=Cairo east=Beijing 
2019/11/26 00:34:16 Istanbul east=Beijing 
2019/11/26 00:34:16 Cairo north=Mumbai 
2019/11/26 00:34:16 -------------------------------------------------
```
## Testing
Run tests for types and game logic:
```
go test ./...
```