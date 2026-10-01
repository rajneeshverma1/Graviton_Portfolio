package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// PortfolioData holds all data rendered into the HTML template
type PortfolioData struct {
	Name       Name
	Bio        string
	Status     Status
	Experience []WorkItem
	Links      []FooterLink
	Notes      []StickyNote
	Year       int
}

type Name struct {
	First string
	Last  string
}

type Status struct {
	Label string
	Live  bool
}

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

type Badge struct {
	Label string
	Class string
}

type FooterLink struct {
	Label string
	URL   string
	ID    string
}

type StickyNote struct {
	ID        string
	Class     string
	Title     string
	PageNum   string
	Icon      string
	HasClip   bool
	SubText   string
}

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
				Company:   "Stealth Startup",
				Role:      "Software Developer · Remote",
				Location:  "US",
				Dates:     "SEP 2026 → NOW",
				IsCurrent: true,
				Description: "Architecting core backend microservices in Go from the ground up. 50+ PRs and 1,800+ commits in month one.",
			},
			{
				Company:  "Amboras",
				Role:     "Product Eng. Intern · Remote",
				Location: "SF BAY AREA",
				Dates:    "FEB → AUG 2026",
				BackedBy: "backed by YC & A*",
				Description: "Engineered high-throughput Go APIs, shipped nonstop: 50+ features, 2,000+ commits, 30k+ users.",
				Badges: []Badge{
					{Label: "YC", Class: "yc-badge"},
					{Label: "A*", Class: "a-badge"},
				},
			},
			{
				Company:  "Dodge AI",
				Role:     "Product Eng. Intern · Remote",
				Location: "BENGALURU",
				Dates:    "AUG → OCT 2025",
				BackedBy: "backed by Google, Accel & Antler",
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
			{ID: "note-notes",   Class: "pink-note",   Title: "notes",       PageNum: "p. 15", Icon: "✏️"},
			{ID: "note-sayhi",   Class: "white-note",  Title: "say hi",      PageNum: "p. 17", Icon: "✈️"},
		},
	}
}

func indexHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.Execute(w, portfolioData()); err != nil {
			http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
			log.Printf("template error: %v", err)
		}
	}
}

func main() {
	// Parse HTML templates
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("failed to parse template: %v", err)
	}

	// Serve static files (CSS, JS, images)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Main route
	http.HandleFunc("/", indexHandler(tmpl))

	addr := ":8080"
	log.Printf("🚀 gravitonportfolio server running at http://localhost%s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
// v2: HTTP server skeleton
// v3: structs refined
