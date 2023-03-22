// Package crdfmt inspired by https://github.com/nerik/formatcoords
package crdfmt

import (
	"fmt"
	"math"
	"strings"
)

type coords struct {
	lon, lat float64
	o        Options

	north, east          bool
	latValues, lonValues values
}

type values struct {
	origValue    float64
	degrees      float64
	degreesInt   float64
	degreesFrac  float64
	secondsTotal float64
	minutes      float64
	minutesInt   float64
	seconds      float64
	secondsInt   float64
}

// float64ToString rounds a number n up to the given precision
func float64ToString(n float64, precision int) string {
	s := math.Pow(10, float64(precision))
	if precision < 0 {
		precision = 0
	}
	return fmt.Sprintf(fmt.Sprintf("%%.%df", precision), math.Round(n*s)/s)
}

// float64ToIntString truncates a float64 to string
func float64ToIntString(n float64) string { return fmt.Sprintf("%d", int64(n)) }

func (c *coords) computeFor(origValue float64) values {
	v := values{
		origValue: origValue,
		degrees:   math.Abs(origValue),
	}

	v.degreesInt = math.Floor(v.degrees)
	v.degreesFrac = v.degrees - v.degreesInt
	v.secondsTotal = 3600 * v.degreesFrac
	v.minutes = v.secondsTotal / 60
	v.minutesInt = math.Floor(v.minutes)
	v.seconds = v.secondsTotal - (v.minutesInt * 60)
	v.secondsInt = math.Floor(v.seconds)
	return v
}

func (c *coords) compute() {
	c.east = c.lon > 0
	c.north = c.lat > 0
	c.lonValues = c.computeFor(c.lon)
	c.latValues = c.computeFor(c.lat)
}

// formatting layouts
const (
	LayoutDeg     = "XDD"
	LayoutMin     = "XDDMM"
	LayoutSec     = "XDDMMSS"
	LayoutSecFrac = "XDDMMss"
)

const (
	unitDegrees = `°`
	unitMinutes = `'`
	unitSeconds = `"`
)

func (c *coords) formatFor(v values, layout string, symbol string) string {
	m := map[bool]string{true: "-"}
	return strings.NewReplacer(
		"DD", float64ToIntString(v.degreesInt)+unitDegrees,
		"dd", float64ToString(v.degrees, c.o.Precision)+unitDegrees,
		"D", float64ToIntString(v.degreesInt),
		"d", float64ToString(v.degrees, c.o.Precision),
		"MM", float64ToIntString(v.minutesInt)+unitMinutes,
		"mm", float64ToString(v.minutes, c.o.Precision)+unitMinutes,
		"M", float64ToIntString(v.minutesInt),
		"m", float64ToString(v.minutes, c.o.Precision),
		"SS", float64ToIntString(v.secondsInt)+unitSeconds,
		"ss", float64ToString(v.seconds, c.o.Precision)+unitSeconds,
		"S", float64ToIntString(v.secondsInt),
		"s", float64ToString(v.seconds, c.o.Precision),
		"-", m[v.origValue < 0],
		"X", symbol,
	).Replace(layout)
}

func (c *coords) format(layout string) (lon, lat string) {
	lonSymbol, latSymbol := "W", "S"
	if c.east {
		lonSymbol = "E"
	}
	if c.north {
		latSymbol = "N"
	}

	lon = c.formatFor(c.lonValues, layout, lonSymbol)
	lat = c.formatFor(c.latValues, layout, latSymbol)
	return
}

// Options is a geo formatter options
type Options struct {
	Precision int
}

// LonTo180 converts longitudes from 0..360 interval to -180..180 interval
func LonTo180(lon float64) float64 { return math.Mod(lon+180, 360) - 180 }

// LonTo360 converts longitudes from -180..180 interval to 0..360 interval
func LonTo360(lon float64) float64 { return math.Mod(lon+360, 360) }

// GeoFormatCoords formats geographical coordinates into geographical form.
// Input values: lon is in 0..360, lat is in -90..90
func GeoFormatCoords(lon, lat float64, layout string, options Options) (slon, slat string) {
	c := coords{
		lon: LonTo180(lon), // from 0..360 to -180..180
		lat: lat,
		o:   options,
	}
	c.compute()
	return c.format(layout)
}
