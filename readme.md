## what is it

The goal of the project is to generate roller coasters via a genetic algorithm

### layout

There are a few different packages in here.

- `genetic` - implementation of a genetic algorithm to build coasters.
  hopefully this will be split out into useful parts, at the moment it's just
  useful to me.

- `exe_reader` - Reads the RCT2 exe and prints out track segment data

- `rle` - deals with run length encoding of RCT files (so they can be
read/written in the game)

- `td6` - encodes/decodes raw TD6 file data into Go structs

- `image/*` - turns track data into pretty PNG files.

- `bits` - convenience functions for dealing with bits.

- `wip/*` - various "discovery" scripts that read bytes from `openrct2.exe`
  dealing with track data.

## next steps

- verify the track data that we have read is accurate

- compute which track pieces are available for a given track piece

- complete overhead view of track in image/above.go

- image of track in isometric view, same as used in RCT2

- actually write the genetic algorithm/fitness function

## as of july 12

- the RLE encoding is good and roundtrip diffs between it and files generated
by the game might be off in parts due to the #s read but they can still be
loaded in the game.

- round tripping a valid ride between td6.Unmarshal and td6.Marshal can still
  generate a valid ride at the other end.
