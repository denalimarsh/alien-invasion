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
2019/11/26 00:34:16 Game completed on turn 10000.
2019/11/26 00:34:16 
2019/11/26 00:34:16                     The World:
2019/11/26 00:34:16 -------------------------------------------------
2019/11/26 00:34:16 Beijing south=Mumbai 
2019/11/26 00:34:16 Mumbai north=Istanbul west=Cairo east=Beijing 
2019/11/26 00:34:16 Istanbul east=Beijing 
2019/11/26 00:34:16 Cairo north=Mumbai 
2019/11/26 00:34:16 -------------------------------------------------
```

As aliens destroy cities, they have a habit of cutting their own transporation lines and getting trapped until the game terminates. If you'd like to aid in the planetary conquest of Earth, you can improve alien technology with the `--advancedTech=true` flag. With their new technology, aliens are able to teleport out of trapped cities whenever they get stuck.

You can try it out with `invasion start --file="./assets/big_world.txt" --numAliens=10 --advancedTech=true`:
```
2019/11/26 04:22:56 The game has started!
2019/11/26 04:22:56 
2019/11/26 04:22:56 Turn 1: London has been destroyed by aliens 2 and 7 
2019/11/26 04:22:56 Turn 2: alien 4 has teleported to Stockton 
2019/11/26 04:22:56 Turn 2: Melbourne has been destroyed by aliens 5 and 6 
2019/11/26 04:22:56 Turn 5: Berlin has been destroyed by aliens 8 and 4 
2019/11/26 04:22:56 Turn 8: alien 3 has teleported to Orlando 
2019/11/26 04:22:56 Turn 10: alien 1 has teleported to Sacramento 
2019/11/26 04:22:56 Turn 11: alien 3 has teleported to Philadelphia 
2019/11/26 04:22:56 Turn 12: alien 1 has teleported to Anchorage 
2019/11/26 04:22:56 Turn 15: alien 3 has teleported to Paris 
2019/11/26 04:22:56 Turn 17: alien 3 has teleported to Tokyonorth=Berlin 
2019/11/26 04:22:56 Turn 19: Delhi has been destroyed by aliens 10 and 3 
2019/11/26 04:22:56 Turn 22: alien 1 has teleported to Singapore 
2019/11/26 04:22:56 Turn 24: alien 1 has teleported to Sydney 
2019/11/26 04:22:56 Turn 26: alien 1 has teleported to Bakersfield 
2019/11/26 04:22:56 Turn 28: alien 1 has teleported to Houston 
2019/11/26 04:22:56 Turn 33: alien 1 has teleported to Hanoi 
2019/11/26 04:22:56 Turn 37: alien 1 has teleported to Orlando 
2019/11/26 04:22:56 Turn 40: alien 1 has teleported to Balitmore 
2019/11/26 04:22:56 Turn 47: alien 1 has teleported to Paris 
2019/11/26 04:22:56 Turn 55: alien 1 has teleported to Stockton 
2019/11/26 04:22:56 Turn 57: Cairo has been destroyed by aliens 9 and 1 
2019/11/26 04:22:56 
2019/11/26 04:22:56 Game completed on turn 58.

```
## Testing
Run tests for types and game logic:
```
go test ./...
```

## Considerations
The aliens are still pretty weak. Even with advanced technology, it's common for at least two aliens to end up in mutually exclusive cycles which leave them effectively isolated. Some potential technological improvements that could help our alien friends might include mass teleporting every 100 turns without a city destruction, as well as repositioning themselves on turn 0 should they be initially placed in an already populated city (and therefore instantly killed). It is also likely that the static world.txt files currently used are poorly suited for alien transportation, a hypothesis that we could test in next iteration by adding dynamic world.txt file generation and running tests against them to uncover the optimal number of cities and paths. Another simple, interesting additional feature could be the introduction of numerous alien lives or team allegiances. For example, if an alien has three lives it could teleport away from city destruction twice before being killed in the final encounter. Teams could potentially make the game much more complex, depending on how interactions between same-team and opposite-team alien encounters are implemented.