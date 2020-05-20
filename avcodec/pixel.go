// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avcodec contains the codecs (decoders and encoders) provided by the libavcodec library
//Provides some generic global options, which can be set on all the encoders and decoders.
package avcodec

//#cgo CFLAGS: -I../../ffmpeg/include
//#cgo LDFLAGS: -L../../ffmpeg/lib/ -llibavformat  -llibavcodec -llibavutil -llibavdevice -llibavfilter -llibswresample -llibswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libswscale/swscale.h>
import "C"

const (
	AV_PIX_FMT_YUV        = 0
	AV_PIX_FMT_YUV420P9   = C.AV_PIX_FMT_YUV420P9
	AV_PIX_FMT_YUV422P9   = C.AV_PIX_FMT_YUV422P9
	AV_PIX_FMT_YUV444P9   = C.AV_PIX_FMT_YUV444P9
	AV_PIX_FMT_YUV420P10  = C.AV_PIX_FMT_YUV420P10
	AV_PIX_FMT_YUV422P10  = C.AV_PIX_FMT_YUV422P10
	AV_PIX_FMT_YUV440P10  = C.AV_PIX_FMT_YUV440P10
	AV_PIX_FMT_YUV444P10  = C.AV_PIX_FMT_YUV444P10
	AV_PIX_FMT_YUV420P12  = C.AV_PIX_FMT_YUV420P12
	AV_PIX_FMT_YUV422P12  = C.AV_PIX_FMT_YUV422P12
	AV_PIX_FMT_YUV440P12  = C.AV_PIX_FMT_YUV440P12
	AV_PIX_FMT_YUV444P12  = C.AV_PIX_FMT_YUV444P12
	AV_PIX_FMT_YUV420P14  = C.AV_PIX_FMT_YUV420P14
	AV_PIX_FMT_YUV422P14  = C.AV_PIX_FMT_YUV422P14
	AV_PIX_FMT_YUV444P14  = C.AV_PIX_FMT_YUV444P14
	AV_PIX_FMT_YUV420P16  = C.AV_PIX_FMT_YUV420P16
	AV_PIX_FMT_YUV422P16  = C.AV_PIX_FMT_YUV422P16
	AV_PIX_FMT_YUV444P16  = C.AV_PIX_FMT_YUV444P16
	AV_PIX_FMT_YUVA420P9  = C.AV_PIX_FMT_YUVA420P9
	AV_PIX_FMT_YUVA422P9  = C.AV_PIX_FMT_YUVA422P9
	AV_PIX_FMT_YUVA444P9  = C.AV_PIX_FMT_YUVA444P9
	AV_PIX_FMT_YUVA420P10 = C.AV_PIX_FMT_YUVA420P10
	AV_PIX_FMT_YUVA422P10 = C.AV_PIX_FMT_YUVA422P10
	AV_PIX_FMT_YUVA444P10 = C.AV_PIX_FMT_YUVA444P10
	AV_PIX_FMT_YUVA420P16 = C.AV_PIX_FMT_YUVA420P16
	AV_PIX_FMT_YUVA422P16 = C.AV_PIX_FMT_YUVA422P16
	AV_PIX_FMT_YUVA444P16 = C.AV_PIX_FMT_YUVA444P16
	AV_PIX_FMT_RGB24      = C.AV_PIX_FMT_RGB24
	AV_PIX_FMT_RGBA       = C.AV_PIX_FMT_RGBA

	SWS_FAST_BILINEAR        = C.SWS_FAST_BILINEAR
	SWS_BILINEAR             = C.SWS_BILINEAR
	SWS_BICUBIC              = C.SWS_BICUBIC
	SWS_X                    = C.SWS_X
	SWS_POINT                = C.SWS_POINT
	SWS_AREA                 = C.SWS_AREA
	SWS_BICUBLIN             = C.SWS_BICUBLIN
	SWS_GAUSS                = C.SWS_GAUSS
	SWS_SINC                 = C.SWS_SINC
	SWS_LANCZOS              = C.SWS_LANCZOS
	SWS_SPLINE               = C.SWS_SPLINE
	SWS_SRC_V_CHR_DROP_MASK  = C.SWS_SRC_V_CHR_DROP_MASK
	SWS_SRC_V_CHR_DROP_SHIFT = C.SWS_SRC_V_CHR_DROP_SHIFT
	SWS_PARAM_DEFAULT        = C.SWS_PARAM_DEFAULT
	SWS_PRINT_INFO           = C.SWS_PRINT_INFO
	SWS_FULL_CHR_H_INT       = C.SWS_FULL_CHR_H_INT
	SWS_FULL_CHR_H_INP       = C.SWS_FULL_CHR_H_INP
	SWS_DIRECT_BGR           = C.SWS_DIRECT_BGR
	SWS_ACCURATE_RND         = C.SWS_ACCURATE_RND
	SWS_BITEXACT             = C.SWS_BITEXACT
	SWS_ERROR_DIFFUSION      = C.SWS_ERROR_DIFFUSION
	SWS_MAX_REDUCE_CUTOFF    = C.SWS_MAX_REDUCE_CUTOFF
	SWS_CS_ITU709            = C.SWS_CS_ITU709
	SWS_CS_FCC               = C.SWS_CS_FCC
	SWS_CS_ITU601            = C.SWS_CS_ITU601
	SWS_CS_ITU624            = C.SWS_CS_ITU624
	SWS_CS_SMPTE170M         = C.SWS_CS_SMPTE170M
	SWS_CS_SMPTE240M         = C.SWS_CS_SMPTE240M
	SWS_CS_DEFAULT           = C.SWS_CS_DEFAULT
	SWS_CS_BT2020            = C.SWS_CS_BT2020
)

func (pf PixelFormat) String() string {
	switch int(pf) {
	case AV_PIX_FMT_YUV420P9:
		return "YUV420P9"

	case AV_PIX_FMT_YUV422P9:
		return "YUV422P9"

	case AV_PIX_FMT_YUV444P9:
		return "YUV444P9"

	case AV_PIX_FMT_YUV420P10:
		return "YUV420P10"

	case AV_PIX_FMT_YUV422P10:
		return "YUV422P10"

	case AV_PIX_FMT_YUV440P10:
		return "YUV440P10"

	case AV_PIX_FMT_YUV444P10:
		return "YUV444P10"

	case AV_PIX_FMT_YUV420P12:
		return "YUV420P12"

	case AV_PIX_FMT_YUV422P12:
		return "YUV422P12"

	case AV_PIX_FMT_YUV440P12:
		return "YUV440P12"

	case AV_PIX_FMT_YUV444P12:
		return "YUV444P12"

	case AV_PIX_FMT_YUV420P14:
		return "YUV420P14"

	case AV_PIX_FMT_YUV422P14:
		return "YUV422P14"

	case AV_PIX_FMT_YUV444P14:
		return "YUV444P14"

	case AV_PIX_FMT_YUV420P16:
		return "YUV420P16"

	case AV_PIX_FMT_YUV422P16:
		return "YUV422P16"

	case AV_PIX_FMT_YUV444P16:
		return "YUV444P16"

	case AV_PIX_FMT_YUVA420P9:
		return "YUVA420P9"

	case AV_PIX_FMT_YUVA422P9:
		return "YUVA422P9"

	case AV_PIX_FMT_YUVA444P9:
		return "YUVA444P9"

	case AV_PIX_FMT_YUVA420P10:
		return "YUVA420P10"

	case AV_PIX_FMT_YUVA422P10:
		return "YUVA422P10"

	case AV_PIX_FMT_YUVA444P10:
		return "YUVA444P10"

	case AV_PIX_FMT_YUVA420P16:
		return "YUVA420P16"

	case AV_PIX_FMT_YUVA422P16:
		return "YUVA422P16"

	case AV_PIX_FMT_YUVA444P16:
		return "YUVA444P16"

	case AV_PIX_FMT_RGB24:
		return "RGB24"

	case AV_PIX_FMT_RGBA:
		return "RGBA"
	}

	return "{UNKNOWN}"
}
