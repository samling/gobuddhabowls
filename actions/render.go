package actions

import (
	"buddhabowls/helpers"
	"fmt"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
	"math"
	"time"
)

var r *render.Engine
var assetsBox = packr.NewBox("../public")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			"format_money": func(val float64) string {
				return fmt.Sprintf("$%.2f", math.Round(val*100)/100)
			},
			"format_date": func(d time.Time) string {
				// if !d.Valid {
				// 	return ""
				// }
				return helpers.FormatDate(d)
			},
			"parseable_date": func(d time.Time) string {
				return helpers.RFC3339Date(d)
			},
			"get_percentage": func(num float64, total float64) float64 {
				return math.Round(100 * num / total)
			},
			"today": func() string {
				t := time.Now()
				return helpers.FormatDate(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()))
			},
		},
	})
}
