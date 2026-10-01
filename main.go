// Package main is the Graviton Portfolio server.
//
// A pixel-perfect interactive portfolio designed to look like an open
// physical binder/notebook. Built entirely in Go using net/http and
// html/template — zero frameworks, zero JavaScript build steps.
//
// Architecture:
//   - All HTML is stored as a Go raw-string constant in tmpl.go
//   - All CSS  is stored as a Go raw-string constant in css.go
//   - All JS   is stored as a Go raw-string constant in js.go
//   - Images   are served from the embedded static/images/ directory
//
// Run:  go run . (or go build -o gravitonportfolio && ./gravitonportfolio)
// Open: http://localhost:8080
package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

// ── Embedded image assets ────────────────────────────────────────────────────

//go:embed static/images
var imageFS embed.FS

// ── Portfolio data model ─────────────────────────────────────────────────────

// PortfolioData is the root data structure passed to the HTML template renderer.
type PortfolioData struct {
	Name       Name
	Bio        string
	Status     Status
	Experience []WorkItem
	Links      []FooterLink
	Notes      []StickyNote
	Year       int
}

// Name holds the split first/last name so the template can style them differently.
type Name struct {
	First string
	Last  string
}

// Status represents the "currently working" indicator with live pulse animation.
type Status struct {
	Label string
	Live  bool
}

// WorkItem represents a single work-experience entry on the left page.
type WorkItem struct {
	Company     string
	Role        string
	Location    string
	Dates       string
	BackedBy    string
	Description string
	Badges      []Badge
	IsCurrent   bool
}

// Badge is a small coloured label (e.g. "YC", "A*", "G").
type Badge struct {
	Label string
	Class string // CSS class controlling badge colour
}

// FooterLink is one of the external links shown at the bottom of page 01.
type FooterLink struct {
	Label string
	URL   string
	ID    string
}

// StickyNote represents one of the six coloured notes on the right page.
type StickyNote struct {
	ID      string
	Class   string // CSS colour class
	Title   string
	PageNum string
	Icon    string
	HasClip bool   // render metallic paperclip SVG above note
	SubText string // optional small italic detail line
}

// ── Data ─────────────────────────────────────────────────────────────────────

// portfolioData returns the complete, hard-coded portfolio content.
// Update this function to change what is displayed on the site.
func portfolioData() PortfolioData {
	return PortfolioData{
		Year: time.Now().Year(),
		Name: Name{First: "Sanatan", Last: "Sharma"},
		Bio:  "I build high-performance backend systems & products: Golang, APIs, microservices, deploy.",
		Status: Status{
			Label: "NOW • 8 → PRODUCTION • STEALTH (US)",
			Live:  true,
		},
		Experience: []WorkItem{
			{
				Company:     "Stealth Startup",
				Role:        "Software Developer · Remote",
				Location:    "US",
				Dates:       "SEP 2026 → NOW",
				IsCurrent:   true,
				Description: "Architecting core backend microservices in Go from the ground up. 50+ PRs and 1,800+ commits in month one.",
			},
			{
				Company:     "Amboras",
				Role:        "Product Eng. Intern · Remote",
				Location:    "SF BAY AREA",
				Dates:       "FEB → AUG 2026",
				BackedBy:    "backed by YC & A*",
				Description: "Engineered high-throughput Go APIs, shipped nonstop: 50+ features, 2,000+ commits, 30k+ users.",
				Badges: []Badge{
					{Label: "YC", Class: "yc-badge"},
					{Label: "A*", Class: "a-badge"},
				},
			},
			{
				Company:     "Dodge AI",
				Role:        "Product Eng. Intern · Remote",
				Location:    "BENGALURU",
				Dates:       "AUG → OCT 2025",
				BackedBy:    "backed by Google, Accel & Antler",
				Description: "Built their whole backend platform in Go, the one they demo to land enterprise customers.",
				Badges: []Badge{
					{Label: "G", Class: "g-badge"},
					{Label: "⚡", Class: "accel-badge"},
					{Label: "🦋", Class: "ant-badge"},
				},
			},
		},
		Links: []FooterLink{
			{Label: "resume",   URL: "#", ID: "link-resume"},
			{Label: "github",   URL: "#", ID: "link-github"},
			{Label: "linkedin", URL: "#", ID: "link-linkedin"},
			{Label: "x",        URL: "#", ID: "link-x"},
			{Label: "email",    URL: "#", ID: "link-email"},
		},
		Notes: []StickyNote{
			{ID: "note-work",     Class: "yellow-note", Title: "work",        PageNum: "p. 03", Icon: "💼", HasClip: true},
			{ID: "note-projects", Class: "peach-note",  Title: "projects",    PageNum: "p. 05", Icon: "📁"},
			{ID: "note-skills",   Class: "purple-note", Title: "what I know", PageNum: "p. 07", Icon: "💡", SubText: "Go concurrency · gRPC · microservices"},
			{ID: "note-tech",     Class: "cyan-note",   Title: "tech I use",  PageNum: "p. 11", Icon: "💻", SubText: "Go · Docker · Postgres · Redis · K8s"},
			{ID: "note-notes",    Class: "pink-note",   Title: "notes",       PageNum: "p. 15", Icon: "✏️"},
			{ID: "note-sayhi",    Class: "white-note",  Title: "say hi",      PageNum: "p. 17", Icon: "✈️"},
		},
	}
}

// ── HTTP Handlers ─────────────────────────────────────────────────────────────

// indexHandler renders the main portfolio page using the embedded HTML template
// and the CSS/JS constants — no disk reads happen at request time.
func indexHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-cache")
		if err := tmpl.Execute(w, portfolioData()); err != nil {
			http.Error(w, "render error: "+err.Error(), http.StatusInternalServerError)
			log.Printf("template render error: %v", err)
		}
	}
}

// cssHandler serves the stylesheet stored as a Go string constant in css.go.
// No external .css file is read — the CSS lives entirely in Go source.
func cssHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=3600")
	_, _ = w.Write([]byte(strings.TrimSpace(CSS)))
}

// jsHandler serves the client-side script stored as a Go string constant in js.go.
// No external .js file is read — the JavaScript lives entirely in Go source.
func jsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=3600")
	_, _ = w.Write([]byte(strings.TrimSpace(JS)))
}

// ── Entry point ───────────────────────────────────────────────────────────────

func main() {
	// Parse the HTML template from the Go string constant (tmpl.go).
	// html/template is used for safe HTML escaping of portfolio data fields.
	tmpl, err := template.New("portfolio").Parse(HTMLTemplate)
	if err != nil {
		log.Fatalf("failed to parse portfolio template: %v", err)
	}

	mux := http.NewServeMux()

	// Main portfolio page
	mux.HandleFunc("/", indexHandler(tmpl))

	// Inline assets — CSS and JS served from Go string constants
	mux.HandleFunc("/static/style.css", cssHandler)
	mux.HandleFunc("/static/script.js", jsHandler)

	// Image assets — served from the embedded static/images/ directory
	// Images are the only binary files; all text assets are pure Go.
	imgServer := http.FileServer(http.FS(imageFS))
	mux.Handle("/static/images/", imgServer)

	addr := ":8080"
	log.Printf("🚀  Graviton Portfolio  →  http://localhost%s", addr)
	log.Printf("📦  HTML template : tmpl.go  (Go raw string constant)")
	log.Printf("🎨  Stylesheet    : css.go   (Go raw string constant)")
	log.Printf("⚡  JavaScript    : js.go    (Go raw string constant)")
	log.Printf("🖼️   Images        : static/images/ (go:embed)")

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
