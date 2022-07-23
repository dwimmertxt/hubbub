# hubbub

Simple module that generates 2D [improved Perlin noise](https://mrl.cs.nyu.edu/~perlin/paper445.pdf).
More or less a Frankenstein's monster of [this](https://adrianb.io/2014/08/09/perlinnoise.html) and [this](https://rtouti.github.io/graphics/perlin-noise-algorithm) implemented in Go. 

Usage is simple:

```go
import (
    "github.com/dwimmertxt/hubbub"
)

var (
    scale       float64 = 0.05
    persistence float64 = 0.5
    octaves     float64 = 4
)
// initialise hubbub struct
hb := hubbub.New()

// get a Perlin noise value in range [0, 1]
nz := hb.Get2D(2.15 * scale, 5.423 * scale)

// get an octaved Perlin noise value in range [0, 1] 
nz := hb.Octave2D(2.15 * scale, 5.423 * scale, persistence, octaves)
        
```
