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

## Layout characters

| character | meaning                                               |
|-----------|-------------------------------------------------------|
| DD        | degrees, integer value with `°` symbol                |
| dd        | degrees, floating point rounded value with `°` symbol |
| D         | degrees, integer value                                |
| d         | degrees, floating point rounded value                 |
| MM        | minutes, integer value with `'` symbol                |
| mm        | minutes, floating point rounded value with `'` symbol |
| M         | minutes, integer value                                |
| m         | minutes, floating point rounded value                 |
| SS        | seconds, integer value with `"` symbol                |
| ss        | seconds, floating point rounded value with `"` symbol |
| S         | seconds, integer value                                |
| s         | seconds, floating point rounded value                 |
| -         | `-` sign if passed value were less than 0             |
| X         | symbol, W or E for longitude, S or N for latitude     |

Floating point rounding precision may be specified in options, default value is `0`. 