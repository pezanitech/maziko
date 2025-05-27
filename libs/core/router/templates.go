package router

// RouteTemplateData contains data needed to generate routesgen file
type RouteTemplateData struct {
	BuildPrefix   string
	Imports       []string
	RouteHandlers []RouteHandler
}

// MetaData contains metadata for the HTML template
type MetaData struct {
	Title       string
	Description string
	URL         string
	Image       string
	ThemeColor  string
	Type        string
}

// DefaultMetaData provides default metadata values
var DefaultMetaData = MetaData{
	Title:       "Maziko",
	Description: "A modern full-stack web framework designed for performance, developer experience, and scalability",
	URL:         "https://maziko.pezani.com/",
	Image:       "https://maziko.pezani.com/images/og-image.jpg",
	ThemeColor:  "#007daa",
	Type:        "website",
}

// RouteHandler represents a route's handler configuration
type RouteHandler struct {
	Path     string
	Method   string
	Package  string
	Function string
}

// RootHTMLTemplate is the base HTML structure for the app
var RootHTMLTemplate = `
<!DOCTYPE html>
<html lang="en">
	<head>
	    <meta charset="UTF-8">
	    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	    <link rel="icon" href="favicon.ico" type="image/x-icon" />

		<meta
            name="theme-color"
            content="{{ .meta.ThemeColor }}"
        />
        <meta
            name="description"
            content="{{ .meta.Description }}"
        />

        <!-- Open Graph / Facebook -->
        <meta
            property="og:type"
            content="{{ .meta.Type }}"
        />
        <meta
            property="og:url"
            content="{{ .meta.URL }}"
        />
        <meta
            property="og:title"
            content="{{ .meta.Title }}"
        />
        <meta
            property="og:description"
            content="{{ .meta.Description }}"
        />
        <meta
            property="og:image"
            content="{{ .meta.Image }}"
        />
        <meta
            property="og:image:secure_url"
            itemprop="image"
            content="{{ .meta.Image }}"
        />

        <!-- Twitter -->
        <meta
            property="twitter:card"
            content="summary_large_image"
        />
        <meta
            property="twitter:url"
            content="{{ .meta.URL }}"
        />
        <meta
            property="twitter:title"
            content="{{ .meta.Title }}"
        />
        <meta
            property="twitter:description"
            content="{{ .meta.Description }}"
        />
        <meta
            property="twitter:image"
            content="{{ .meta.Image }}"
        />

        <title>{{ .meta.Title }}</title>

	    <link href="{{ vite "app/global.css" }}" rel="stylesheet">
	    {{ .inertiaHead }}

	    {{ if .hmr }}
	        <script type="module">
	            import RefreshRuntime from '{{ vite "@react-refresh" }}'
	            RefreshRuntime.injectIntoGlobalHook(window)
	            window.$RefreshReg$ = () => { }
	            window.$RefreshSig$ = () => (type) => type
	            window.__vite_plugin_react_preamble_installed__ = true
	        </script>
	    {{ end }}
	    <style>
	        #loading-spinner {
	            position: fixed;
	            top: 0;
	            left: 0;
	            width: 100%;
	            height: 100%;
	            display: flex;
	            justify-content: center;
	            align-items: center;
	            background-color: #ffffff;
	            z-index: 9999;
	            transition: opacity 0.3s ease-out;
	        }
	        
	        @media (prefers-color-scheme: dark) {
	            #loading-spinner {
	                background-color: #000000;
	            }
	            .spinner {
	                border-color: rgba(255, 255, 255, 0.1);
	                border-top-color: {{ .meta.ThemeColor }};
	            }
	        }
	        
	        .spinner {
	            width: 60px;
	            height: 60px;
	            border: 6px solid rgba(0, 0, 0, 0.1);
	            border-radius: 50%;
	            border-top-color: {{ .meta.ThemeColor }};
	            animation: spin 1s ease-in-out infinite;
	        }
	        @keyframes spin {
	            to { transform: rotate(360deg); }
	        }
	        .spinner-hidden {
	            opacity: 0;
	            pointer-events: none;
	        }
	    </style>
	</head>

	<body class="font-sans antialiased">
	    <div id="loading-spinner">
	        <div class="spinner"></div>
	    </div>
	    {{ .inertia }}
	    <script type="module" src="{{ vite "app/app.jsx" }}"></script>
	    <script>
	        // Hide spinner when the app is loaded
	        window.addEventListener('DOMContentLoaded', function() {
	            // Small delay to ensure React has started rendering
	            setTimeout(function() {
	                const spinner = document.getElementById('loading-spinner');
	                if (spinner) {
	                    spinner.classList.add('spinner-hidden');
	                    // Remove from DOM after transition completes
	                    setTimeout(function() {
	                        spinner.parentNode.removeChild(spinner);
	                    }, 300);
	                }
	            }, 100);
	        });
	    </script>
	</body>
</html>`

// RoutesTemplate is the template for the routesgen file
var RoutesTemplate = `// GENERATED FILE - DO NOT EDIT
// auto-generated by Maziko
package gen

import (
{{- range .Imports}}
	{{- if contains . "/_"}}
	{{extractParent .}}__{{extractName .}} {{.}}
	{{- else}}
	{{.}}
	{{- end}}
{{- end}}
)

func Routes() {
{{- range .RouteHandlers}}
	{{- if contains .Package "."}}
	{{before .Package "."}}_{{after .Package "."}}.Route()
	{{- else}}
	{{.Package}}.Route()
	{{- end}}
{{- end}}
}`
