# png2rle

This is a go command line app that converts a .png formatted image file into a RLE datastructure that I use in Watchy-Screen. It takes a single png file name on the command line and writes the rle datastructure definition to standard out.

It's not very robust, for example the output contains the path info of the input file if it isn't in the local directory.

Only tested on BW images.

## Installation

1. Install Go on your system (e.g. follow the instructions at https://golang.org/doc/install or `apt install golang`)
2. Clone or download this repo
3. `go build`
4. ./png2rle <pngfile> 
