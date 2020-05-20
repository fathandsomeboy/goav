// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo CFLAGS: -I../../ffmpeg/include
//#cgo LDFLAGS: -L../../ffmpeg/lib/ -llibavformat  -llibavcodec -llibavutil -llibavdevice -llibavfilter -llibswresample -llibswscale
//#include <libavcodec/avcodec.h>
import "C"

const (
	AV_CODEC_FLAG_UNALIGNED            = int(C.AV_CODEC_FLAG_UNALIGNED)
	AV_CODEC_FLAG_QSCALE               = int(C.AV_CODEC_FLAG_QSCALE)
	AV_CODEC_FLAG_4MV                  = int(C.AV_CODEC_FLAG_4MV)
	AV_CODEC_FLAG_OUTPUT_CORRUPT       = int(C.AV_CODEC_FLAG_OUTPUT_CORRUPT)
	AV_CODEC_FLAG_QPEL                 = int(C.AV_CODEC_FLAG_QPEL)
	AV_CODEC_FLAG_PASS1                = int(C.AV_CODEC_FLAG_PASS1)
	AV_CODEC_FLAG_PASS2                = int(C.AV_CODEC_FLAG_PASS2)
	AV_CODEC_FLAG_LOOP_FILTER          = int(C.AV_CODEC_FLAG_LOOP_FILTER)
	AV_CODEC_FLAG_GRAY                 = int(C.AV_CODEC_FLAG_GRAY)
	AV_CODEC_FLAG_PSNR                 = int(C.AV_CODEC_FLAG_PSNR)
	AV_CODEC_FLAG_TRUNCATED            = int(C.AV_CODEC_FLAG_TRUNCATED)
	AV_CODEC_FLAG_INTERLACED_DCT       = int(C.AV_CODEC_FLAG_INTERLACED_DCT)
	AV_CODEC_FLAG_LOW_DELAY            = int(C.AV_CODEC_FLAG_LOW_DELAY)
	AV_CODEC_FLAG_GLOBAL_HEADER        = int(C.AV_CODEC_FLAG_GLOBAL_HEADER)
	AV_CODEC_FLAG_BITEXACT             = int(C.AV_CODEC_FLAG_BITEXACT)
	AV_CODEC_FLAG_AC_PRED              = int(C.AV_CODEC_FLAG_AC_PRED)
	AV_CODEC_FLAG_INTERLACED_ME        = int(C.AV_CODEC_FLAG_INTERLACED_ME)
	AV_CODEC_FLAG_CLOSED_GOP           = int(C.AV_CODEC_FLAG_CLOSED_GOP)
	AV_CODEC_FLAG2_FAST                = int(C.AV_CODEC_FLAG2_FAST)
	AV_CODEC_FLAG2_NO_OUTPUT           = int(C.AV_CODEC_FLAG2_NO_OUTPUT)
	AV_CODEC_FLAG2_LOCAL_HEADER        = int(C.AV_CODEC_FLAG2_LOCAL_HEADER)
	AV_CODEC_FLAG2_DROP_FRAME_TIMECODE = int(C.AV_CODEC_FLAG2_DROP_FRAME_TIMECODE)
	AV_CODEC_FLAG2_CHUNKS              = int(C.AV_CODEC_FLAG2_CHUNKS)
	AV_CODEC_FLAG2_IGNORE_CROP         = int(C.AV_CODEC_FLAG2_IGNORE_CROP)
	AV_CODEC_FLAG2_SHOW_ALL            = int(C.AV_CODEC_FLAG2_SHOW_ALL)
	AV_CODEC_FLAG2_EXPORT_MVS          = int(C.AV_CODEC_FLAG2_EXPORT_MVS)
	AV_CODEC_FLAG2_SKIP_MANUAL         = int(C.AV_CODEC_FLAG2_SKIP_MANUAL)
	AV_CODEC_FLAG2_RO_FLUSH_NOOP       = int(C.AV_CODEC_FLAG2_RO_FLUSH_NOOP)
)
