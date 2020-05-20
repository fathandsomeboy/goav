// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avfilter

/*
#cgo CFLAGS: -I../ffmpeg/include
#cgo LDFLAGS: -L../ffmpeg/lib/ -llibavformat  -llibavcodec -llibavutil -llibavdevice -llibavfilter -llibswresample -llibswscale
	#include <libavfilter/avfilter.h>
*/
import "C"

//Get a filter definition matching the given name.
func AvfilterGetByName(n string) *Filter {
	return (*Filter)(C.avfilter_get_by_name(C.CString(n)))
}

//Register a filter.
func (f *Filter) AvfilterRegister() int {
	return int(C.avfilter_register((*C.struct_AVFilter)(f)))
}

//Iterate over all registered filters.
func (f *Filter) AvfilterNext() *Filter {
	return (*Filter)(C.avfilter_next((*C.struct_AVFilter)(f)))
}
