// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo CFLAGS: -I../../ffmpeg/include
//#cgo LDFLAGS: -L../../ffmpeg/lib/ -llibavformat  -llibavcodec -llibavutil -llibavdevice -llibavfilter -llibswresample -llibswscale
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"

	"github.com/fathandsomeboy/goav/avcodec"
)

func toCPacket(pkt *avcodec.Packet) *C.struct_AVPacket {
	return (*C.struct_AVPacket)(unsafe.Pointer(pkt))
}

func fromCPacket(pkt *C.struct_AVPacket) *avcodec.Packet {
	return (*avcodec.Packet)(unsafe.Pointer(pkt))
}
