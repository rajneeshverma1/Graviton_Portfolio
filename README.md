# 📓 Graviton Portfolio

> A pixel-perfect, interactive portfolio web app designed to look like an open physical binder/notebook — built with **Go (net/http)** on the backend and vanilla HTML/CSS/JS on the frontend.

[![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat-square&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)

---

## ✨ Features

- 🗒️ **Binder/Notebook UI** — Two-page spread with leather cover edges, metallic ring spine, and paper textures
- 📌 **Interactive Sticky Notes** — 6 colour-coded notes with hover lift effect and modal expansion
- 🌙 **Polaroid Photography** — Night skyline + Hokusai Great Wave prints with tape overlays
- 🖊️ **Decorative Pen** — Metallic pen clipped to the binder edge with "antigravity" engraving
- 🎨 **van Gogh Sticker** — Circular portrait sticker with handwritten annotation
- ⚡ **Cursor Parallax** — Subtle 3D tilt as you move your mouse
- 🔠 **Mixed Typography** — Caveat (handwritten), Inter (body), DM Serif Display (cursive name), Space Mono (dates/badges)
- 🐹 **Pure Go backend** — `net/http` + `html/template`, zero dependencies

## 🛠️ Tech Stack

| Layer     | Technology          |
|-----------|---------------------|
| Backend   | Go (`net/http`)     |
| Templates | Go `html/template`  |
| Frontend  | HTML5, CSS3, Vanilla JS |
| Fonts     | Google Fonts (Caveat, Inter, DM Serif Display, Space Mono, Permanent Marker) |
| Images    | AI-generated (skyline, Great Wave, Van Gogh sticker) |

## 🚀 Run Locally

```bash
git clone https://github.com/rajneeshverma1/Graviton_Portfolio.git
cd Graviton_Portfolio
go run main.go
# → open http://localhost:8080
```

Or build the binary:

```bash
go build -o gravitonportfolio .
./gravitonportfolio
```

## 📁 Project Structure

```
gravitonportfolio/
├── main.go                  # Go HTTP server + portfolio data structs
├── go.mod                   # Go module definition
├── templates/
│   └── index.html           # Go HTML template (binder layout)
├── static/
│   ├── style.css            # Full CSS — grid desk, binder, notes, modal
│   ├── script.js            # Sticky note modal, parallax, blink animation
│   └── images/
│       ├── polaroid_skyline.png
│       ├── great_wave.png
│       └── van_gogh_sticker.png
└── README.md
```

## 📄 License

MIT © Sanatan Sharma

