package main

import (
	"bufio"
	"flag"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

/*
                                                                                                                                        .                *
     ##            :.                                                                 ::::::::::::::+#@                                       +@*               *
     @@           W@*                                                                #@@@@@@@@@@@@@@@@+                                       +@@:              *
     #@.          @@.                                                                ##W@*::    .::::.                                        .@@.              *
     +@:          W@            :    :+    :.          ::        :    :+               +@:                :          :.     .     :::.         @@               *
     +@+          #@        .*@@@@W+ *@:  *@+         #@W    .*@@@@W+ *@:              +@.               :@#         @@+   #@+ .#@@@@@W.       @W               *
     :@+          #W       :@@*:::*@W#@+  W@*         @@.   :@@*:::*@W#@+              +@.               :@W         W@+   #@**@#:   +@@.      @W               *
     :@+          ##      .@@.      W@@:  +@#        :@*   .@@.      W@@:              +@                .@*         #@:   *@W@.      :@W      @W               *
    .#@#::::::::::WW*#:   W@.        @@.   #@        #@    W@.        @@.              +@                .@+         #@.   +@@.        #@.     @W               *
   .@@@@@@@@@@@@@@@@@@    @#         *@.    @+       @*    @#         *@.              +@*:######@@+      @:         *@    :@*         :@:     @W               *
   +@W@#:.       :#@*.   :@:         +@.    *W      +@    :@:         +@.              +@@@@@@@@@@@*      @:         *@    .@+         :@+     @W               *
     #@+          *W     +@.         +@.     @:     W*    +@.         +@.              *@:                @:         *@    .@+         :@+     @@               *
     @@:          #@     +@.         :@:     +W    :@     +@.         :@:              *@:                @:         *@    :@+         :@+    .@@               *
    .@@.          #@     :@:         :@+      @+   W*     :@:         :@+              *@:                @+         #@    :@+         +@*    .@@               *
    *@@           W@:     @#         +@*      *@  .@       @#         +@*              *@+                @*        .@@    :@*         +@*     :.               *
    @@#           W@+     #@+       +@@W       @# W+       #@+       +@@W              #@+                W@        W@@    +@*         +@*     ::               *
   *@@+           W@+      W@*.   :W@*@@       +@W@         W@*.   :W@*@@              #@*                :@W:    :WW#@.   +@#         *@*    :@@*              *
   W@W            W@+       #@@@@@@W..@#        @@#          #@@@@@@W..@#              #@+                 +@@@@@@@W W@:   +@#         *@*    *@@*              *
   +#.            :#          :+#:.             .+             :+#:.                   :*                    ::##:.  :*    :+          :+      #*
*/

var (
	Quality int
	path    string // Pic path.
	toFile  string
)

var gs = []rune{
	'@', 'W', '#', '*', '+', ':', '.', ' ',
}

func init() {
	flag.IntVar(&Quality, "q", 1, "-q=1")
	flag.StringVar(&path, "path", "", "-path=hi.txt")
	flag.StringVar(&toFile, "to", "", "-to=result.txt")
	flag.Parse()
}

func main() {
	// Opena image.
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	gray := image.NewGray(img.Bounds())

	for x := 0; x < gray.Rect.Max.X; x++ {
		for y := 0; y < gray.Rect.Max.Y; y++ {
			oldColor := img.At(x, y)
			grayColor := color.GrayModel.Convert(oldColor)
			// g, _, _, _ := grayColor.RGBA()
			gray.Set(x, y, grayColor)
		}
	}

	getGray := func(x, y int) int {
		g, _, _, _ := gray.At(x, y).RGBA()
		// buff.WriteRune(gs[g>>8/4])
		return int(g >> 8)
	}

	// Writer
	var writer *bufio.Writer
	if toFile != "" {
		file, err := os.Create(toFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		writer = bufio.NewWriter(file)
	} else {
		writer = bufio.NewWriter(os.Stdout)
	}

	for y := 0; y < gray.Rect.Max.Y; y += Quality {
		for x := 0; x < gray.Rect.Max.X; x += Quality {
			// Get average gray.
			total := 0
			n := 0
			for i := 0; i < Quality; i++ {
				for j := 0; j < Quality; j++ {
					temp_x, temp_y := x+i, y+j
					if temp_x > gray.Rect.Max.X || temp_y > gray.Rect.Max.Y {
						continue
					}
					total += getGray(temp_x, temp_y)
					n++
				}
			}
			// Put.
			index := total / (n * 256 / len(gs))
			writer.WriteRune(gs[index])
			recod(index)
		}
		writer.WriteRune('\n')
	}
	writer.Flush()
	println("Done!")
}

var Map = make(map[int]int)

func recod(i int) {
	if n := Map[i]; n == 0 {
		Map[i] = 0
	}
	Map[i]++
}
