package main

import (
	"fmt"
	"os"
	"encoding/xml"
	"io/ioutil"
)

type Gpx struct {
	XMLName xml.Name `xml:"gpx"`
	Creator string `xml:"creator,attr"`
	Metadata Metadata `xml:"metadata"`
	Trk Trk `xml:"trk"`
}

type Metadata struct {
	XMLName xml.Name `xml:"metadata"`
	MdTime string `xml:"time"`
}

type Trk struct {
	XMLName xml.Name `xml:"trk"`
	TrkType string `xml:"type"`
	Extensions Extensions `xml:"extensions"`
	TrkSeg TrkSeg `xml:"trkseg"`
}

type Extensions struct {
	XMLName xml.Name `xml:"extensions"`
	TotalTime float32 `xml:"totalTime"`
	TotalDistance float32 `xml:"totalDistance"`
}

type TrkSeg struct {
	XMLName xml.Name `xml:"trkseg"`
	TrkPtList []Trkpt `xml:"trkpt"`
}

type Trkpt struct {
	XMLName xml.Name `xml:"trkpt"`
	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`
	PtTime string `xml:"time"`
}

func main() {
	xmlFile, err := os.Open("20220520.gpx")
	if err != nil {
		fmt.Println("File open error ", err)
		return
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var gpx Gpx

	xml.Unmarshal(byteValue, &gpx)

	fmt.Println(gpx.Trk.Extensions.TotalTime)
	fmt.Println(gpx.Trk.Extensions.TotalDistance)
	fmt.Println(gpx.Metadata.MdTime)
	fmt.Println("---------------------------")

	fmt.Println(gpx.Trk.TrkSeg.TrkPtList[0].Lat)
	fmt.Println(gpx.Trk.TrkSeg.TrkPtList[0].Lon)
	fmt.Println(gpx.Trk.TrkSeg.TrkPtList[0].PtTime)
	fmt.Println("---------------------------")

	for _, pkt := range gpx.Trk.TrkSeg.TrkPtList {
		fmt.Printf(" POINT ( %f, %f ) | %s \n", pkt.Lat, pkt.Lon, pkt.PtTime)
	}
}
