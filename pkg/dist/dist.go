package dist

import (
	"html/template"
	"io"
)

var (
	StaticFiles = map[string][]byte{
		"favicon-16x16.png":               favicon16x16Png,
		"favicon-32x32.png":               favicon32x32Png,
		"oauth2-redirect.html":            oAuth2RedirectHtml,
		"swagger-ui-bundle.js":            swaggerUiBundleJs,
		"swagger-ui.css":                  swaggerUiCss,
		"swagger-ui-standalone-preset.js": swaggerUiStandalonePresetJs,
	}

	indexHtmlTemplate = template.Must(
		template.New("index.html").Parse(indexHtml))

	pluginDownloadUrl = "SwaggerUIBundle.plugins.DownloadUrl"
	pluginHideTopbar  = "HideTopbarPlugin"
)

func ExecuteIndexHtml(w io.Writer, docUrl string, hideTopbar bool) error {
	plugins := pluginDownloadUrl
	if hideTopbar {
		plugins = plugins + ", " + pluginHideTopbar
	}

	config := struct {
		DocUrl  string
		Plugins string
	}{
		DocUrl:  docUrl,
		Plugins: plugins,
	}

	return indexHtmlTemplate.Execute(w, config)
}
