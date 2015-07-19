# What is it?

The goal of the project is to generate cool looking roller coasters via a
genetic algorithm. Something like this:

<img
src="https://kev.inburke.com/slides/coasters/images/double-dare-better.png"
alt="The coolest coaster ever" />

## What's been done so far?

Here's what we have:

- A list of all of the track pieces used in the game and metadata about them
  (elevation change, left/right etc)

- A way to read rides from the game, and serialize track data generated outside
  the game to be used in the game.

- Some library functions to make vector math easier (e.g. after this track
piece, what is the new location and orientation of the car?)

- A basic genetic algorithm with mutation, selection, crossover, and a fitness
  function.

- Results of every experiment are recorded to disk. There's a server you can
use to view the results of experiments, and which has an easy "convert to TD6"
button.

## What needs to get done?

- [The fitness function][fitness] does not have enough inputs. Specifically, it
currently checks whether a track is a complete loop and doesn't collide with
itself. It *could* check:

    - Excitement. The easiest heuristic for excitement is [what's
    used to compute excitement in the game Roller Coaster Tycoon
    2][openrct2-excitement]; would take into account the number of drops, the
    speed, G forces etc.

    - Physics - whether the car can make it around the track.

    - G forces

    - More ! Scenery, etc, you name it.

- The fitness function is currently **very slow**. Need to run the code with
`pprof` to figure out what exactly is slow, and then improve its performance.
The existing code is extremely dumb - it runs sequentially in one thread.
Likely there are some improvements to be made by using more than one CPU at a
time.

- The server can display track data but the camera is not right, so sometimes
the track is not visible. It would be really neat to be able to better see
track data outside of the game.

- The genetic algorithm will probably need some tweaking once the fitness
function has more parameters. This is hard to evaluate right now.

[fitness]: https://github.com/kevinburke/rct/blob/master/genetic/rct2.go#L436
[openrct2-excitement]: https://github.com/OpenRCT2/OpenRCT2/blob/develop/src/ride/ride_ratings.c#L2106

## Installation

You need to have Go installed (I use Go 1.4, but it should work with any
version). Once you do that, run:

```bash
go get github.com/kevinburke/rct
```

You may or may not need to run `make install` after this.

Check you have a valid install by running: `make test`. You should get output
like this:

```bash
go test -timeout 1s \
		./bits/... \
		./genetic/... \
		./geo/... \
		./image/... \
		./rle/... \
		./server/... \
		./td6/... \
		./tracks/...
?   	github.com/kevinburke/rct/bits	[no test files]
ok  	github.com/kevinburke/rct/genetic	0.011s
?   	github.com/kevinburke/rct/genetic/get_latest_track	[no test files]
?   	github.com/kevinburke/rct/genetic/get_old_experiments	[no test files]
?   	github.com/kevinburke/rct/genetic/run_experiment	[no test files]
ok  	github.com/kevinburke/rct/geo	0.010s
?   	github.com/kevinburke/rct/image	[no test files]
?   	github.com/kevinburke/rct/image/above_runner	[no test files]
ok  	github.com/kevinburke/rct/rle	0.011s
?   	github.com/kevinburke/rct/rle/decode_td6	[no test files]
?   	github.com/kevinburke/rct/rle/encode_td6	[no test files]
?   	github.com/kevinburke/rct/server	[no test files]
ok  	github.com/kevinburke/rct/td6	0.010s
ok  	github.com/kevinburke/rct/tracks	0.008s
?   	github.com/kevinburke/rct/tracks/branch_factor	[no test files]
```

(Some of the packages don't have tests unfortunately)

## Running experiments

You should be able to run new experiments by typing `make experiment`. This
will place new experiments in subdirectories in `/usr/local/rct`. Start the
server by running `make serve` and you should be able to view your experiment
by browsing at `localhost:8080`.

### Package Layout

There are a few different packages in here.

- `genetic` - implementation of a genetic algorithm to build coasters.
  Hopefully this can become a better implementation than just RCT specific but
  there are other genetic algorithm implementations for Go at the moment, so
  this is not a huge priority.

- `exe_reader` - Reads the RCT2 exe and prints out metadata about track
segments.

- `rle` - deals with run length encoding of RCT files (so they can be
read/written in the game). More info in the specific package.

- `td6` - a td6 file is a RCT ride file. This package encodes/decodes raw TD6
file data into Go structs.

- `image/*` - turns track data into pretty PNG files. Hasn't been touched in
  a while and only ever implemented a 2d viewer. Decided to try using three.js
  instead, for greater compatibility and access to a better feature set.

- `bits` - convenience functions for dealing with bits.

- `wip/*` - various "discovery" scripts that read bytes from `openrct2.exe`
dealing with track data. Shouldn't need to touch this, we might need to clean
it up at some point because Go complains about the number of `main` files in
this folder.
