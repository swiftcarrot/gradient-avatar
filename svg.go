package avatar

import (
	"bytes"
	"text/template"
)

type SVGData struct {
	Color1   string
	Color2   string
	Text     string
	Width    int
	Height   int
	FontSize float64
}

func CreateSVG(data SVGData) (string, error) {
	t := template.Must(template.New("svg").Parse(`<?xml version="1.0" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg width="{{.Width}}" height="{{.Height}}" viewBox="0 0 {{.Width}} {{.Height}}" version="1.1" xmlns="http://www.w3.org/2000/svg">
  <g>
    <defs>
      <linearGradient id="avatar" x1="0" y1="0" x2="1" y2="1">
        <stop offset="0%" stop-color="{{.Color1}}"/>
        <stop offset="100%" stop-color="{{.Color2}}"/>
      </linearGradient>
    </defs>
    <rect fill="url(#avatar)" x="0" y="0" width="{{.Width}}" height="{{.Height}}"/>
    <text x="50%" y="50%" alignment-baseline="central" dominant-baseline="middle" text-anchor="middle" fill="#fff" font-family="sans-serif" font-size="{{.FontSize}}">{{.Text}}</text>
  </g>
</svg>
`))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return "", err
	}

	result := tpl.String()
	return result, nil
}
