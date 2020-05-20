package main

import (
	"log"

	"github.com/fathandsomeboy/goav/avcodec"
	"github.com/fathandsomeboy/goav/avdevice"
	"github.com/fathandsomeboy/goav/avfilter"
	"github.com/fathandsomeboy/goav/avformat"
	"github.com/fathandsomeboy/goav/avutil"
	"github.com/fathandsomeboy/goav/swresample"
	"github.com/fathandsomeboy/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
