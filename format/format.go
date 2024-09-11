package format

import (
	"github.com/ducthanh98/vdk/av/avutil"
	"github.com/ducthanh98/vdk/format/aac"
	"github.com/ducthanh98/vdk/format/flv"
	"github.com/ducthanh98/vdk/format/mp4"
	"github.com/ducthanh98/vdk/format/rtmp"
	"github.com/ducthanh98/vdk/format/rtsp"
	"github.com/ducthanh98/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
