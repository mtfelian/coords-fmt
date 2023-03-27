package crdfmt_test

import (
	"testing"

	crdfmt "github.com/mtfelian/go-coords-fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoCoordsFmt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoCoordsFmt Suite")
}

var _ = Describe("GeoFormatCoords", func() {
	type tc struct {
		lon, lat   float64
		layout     string
		sLon, sLat string
		options    crdfmt.Options
	}
	It("checks cases", func() {
		for i, tc := range []tc{
			{
				lon:     149.128684,
				lat:     -35.282000,
				layout:  "-D M s",
				sLon:    `149 7 43.26240`,
				sLat:    `-35 16 55.20000`,
				options: crdfmt.Options{Precision: 5},
			},
			{
				lon:     149.128684,
				lat:     -35.282000,
				layout:  "DD MM ss X",
				sLon:    `149° 7' 43" E`,
				sLat:    `35° 16' 55" S`,
				options: crdfmt.Options{Precision: 0},
			},
			{
				lon:     149.128684,
				lat:     -35.282000,
				layout:  "XDDMMSS",
				sLon:    `E149°7'43"`,
				sLat:    `S35°16'55"`,
				options: crdfmt.Options{Precision: 0},
			},
			{
				lon:     350.128684,
				lat:     -35.282000,
				layout:  "XDDMMSS",
				sLon:    `W9°52'16"`,
				sLat:    `S35°16'55"`,
				options: crdfmt.Options{Precision: 0},
			},
			{
				lon:     0,
				lat:     0,
				layout:  "XDDMMSS",
				sLon:    `W0°0'0"`,
				sLat:    `S0°0'0"`,
				options: crdfmt.Options{Precision: 0},
			},
		} {
			rLon, rLat := crdfmt.GeoFormatCoords(tc.lon, tc.lat, tc.layout, tc.options)
			Expect(rLon).To(Equal(tc.sLon), "rLon at i=%d", i)
			Expect(rLat).To(Equal(tc.sLat), "rLat at i=%d", i)
		}
	})
})
