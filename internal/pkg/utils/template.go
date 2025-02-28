package utils

import (
	"html/template"

	"github.com/gogf/gf/v2/os/gtime"
)

func GetOutputPathTemplate() *template.Template {
	return template.
		Must(template.
			New("outputPathTemplate").
			Funcs(getFuncsMap()).
			Parse(`{{ outputPath }}/{{ .Platform }}/{{ .Anchor }}/{{ currentMonth }}/`))
}

func GetFilenameTemplate(outputPath, format string) *template.Template {
	return template.
		Must(template.
			New("filenameTemplate").
			Funcs(getFuncsMap()).
			Parse(outputPath + `[{{ currentTime }}][{{ .Anchor }}][{{ .RoomName }}].` + format))
}

func getFuncsMap() template.FuncMap {
	return template.FuncMap{
		"currentTime": func() string {
			return gtime.Datetime()
		},
		"currentDate": func() string {
			return gtime.Date()
		},
		"currentMonth": func() string {
			return gtime.Now().Format("Y-m")
		},
		"outputPath": GetOutputPath,
	}
}
