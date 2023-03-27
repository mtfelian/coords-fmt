# Go Coordinates Formatter

This library implements simple geographical coordinates formatter in Go. 

This were inspired by the appropriate JS library: https://github.com/nerik/formatcoords

## Import

`import github.com/mtfelian/go-coords-fmt`

## Usage example 

```

import github.com/mtfelian/go-coords-fmt

// ...

sLon, sLat := crdfmt.GeoFormatCoords(149.128684, -35.282, "XDDMMSS", crdfmt.Options{Precision: 0})
```
gives 	
```
    sLon: `E149°7'43"`
	sLat: `S35°16'55"`
```	
with degree symbol, minutes and seconds.
