# png2rle

This is a go command line app that converts a .png formatted image file into a RLE datastructure that I use in Watchy-Screen. It takes a single png file name on the command line and writes the rle datastructure definition to standard out.

It's not very robust, for example the output contains the path info of the input file if it isn't in the local directory.

Only tested on BW images.
