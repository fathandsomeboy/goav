// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo CFLAGS: -I../ffmpeg/include
//#cgo LDFLAGS: -L../ffmpeg/lib/ -llibavformat  -llibavcodec -llibavutil -llibavdevice -llibavfilter -llibswresample -llibswscale
//#include <libavutil/avutil.h>
import "C"
import "github.com/fathandsomeboy/goav/avcodec"

func newRational(r C.struct_AVRational) avcodec.Rational {
	return avcodec.NewRational(int(r.num), int(r.den))
}
