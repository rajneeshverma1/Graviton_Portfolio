package main

// HTMLTemplate is the portfolio page template, rendered server-side by Go html/template.
const HTMLTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>{{.Name.First}} {{.Name.Last}} — Portfolio</title>
  <meta name="description" content="{{.Name.First}} {{.Name.Last}} — {{.Bio}}"/>
  <link rel="preconnect" href="https://fonts.googleapis.com"/>
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin/>
  <link href="https://fonts.googleapis.com/css2?family=Caveat:wght@400;600;700&family=Inter:wght@300;400;500;600;700;800;900&family=DM+Serif+Display:ital@0;1&family=Space+Mono:wght@400;700&family=Permanent+Marker&display=swap" rel="stylesheet"/>
  <link rel="stylesheet" href="/static/style.css"/>
</head>
<body>

<!-- Warm sandy desk background -->
<div class="desk-bg">
  <!-- Plant shadow SVG overlay -->
  <svg class="plant-shadow" viewBox="0 0 400 600" xmlns="http://www.w3.org/2000/svg">
    <ellipse cx="80" cy="120" rx="60" ry="100" fill="rgba(120,100,50,0.08)" transform="rotate(-20,80,120)"/>
    <ellipse cx="60" cy="200" rx="40" ry="80" fill="rgba(100,80,40,0.07)" transform="rotate(10,60,200)"/>
    <ellipse cx="100" cy="80" rx="50" ry="90" fill="rgba(110,90,45,0.06)" transform="rotate(-35,100,80)"/>
  </svg>
</div>

<!-- Book wrapper — animates open on load -->
<div class="book-scene">
  <div class="binder-wrapper" id="binder">

    <!-- Left teal cover edge -->
    <div class="binder-cover-edge left-edge"></div>

    <!-- Metallic ring spine -->
    <div class="binder-spine">
      <div class="ring-group">
        <div class="ring"><div class="ring-inner"></div></div>
        <div class="ring"><div class="ring-inner"></div></div>
        <div class="ring"><div class="ring-inner"></div></div>
        <div class="ring"><div class="ring-inner"></div></div>
      </div>
    </div>

    <!-- =============== LEFT PAGE p.01 =============== -->
    <div class="page page-left" id="page-01">
      <div class="page-grid-overlay"></div>

      <!-- Polaroid cluster -->
      <div class="polaroid-cluster">
        <!-- Main large polaroid -->
        <div class="polaroid main-polaroid">
          <div class="tape tape-l"></div>
          <div class="tape tape-r"></div>
          <img src="/static/images/polaroid_skyline.png" alt="Night city skyline" class="pol-img"/>
        </div>
        <!-- Mini "me" polaroid overlapping -->
        <div class="polaroid mini-polaroid">
          <div class="tape tape-m"></div>
          <div class="mini-dark">
            <span class="mini-moon">🌙</span>
          </div>
          <div class="pol-label">me</div>
        </div>
      </div>

      <!-- Handwritten caption -->
      <p class="skyline-caption"><em>every window is a day of my last year. the blinking one is <span class="blink-today">today.</span></em></p>

      <!-- Name -->
      <div class="name-block">
        <h1 class="name-heading">
          <span class="name-first">{{.Name.First}}</span> <span class="name-last">{{.Name.Last}}</span>
        </h1>
        <p class="bio-text">I build <span class="strike">websites</span> products end to end:<br/><span class="bio-highlight">screen, API, database, deploy.</span></p>
        <div class="status-row">
          <span class="live-dot"></span>
          <span class="status-label">NOW &nbsp;·&nbsp; <strong>G &rarr; PRODUCTION</strong> &nbsp;·&nbsp; STEALTH (US)</span>
        </div>
      </div>

      <!-- Work section -->
      <div class="work-section">
        <div class="work-header">
          <span class="work-title">Work</span>
          <span class="work-count">3 STARTUPS</span>
          <span class="classic-tag">CLASSIC<br/>FREE</span>
        </div>

        <div class="work-items">
          {{range .Experience}}
          <div class="work-entry">
            <div class="work-entry-top">
              <div class="work-dot{{if .IsCurrent}} dot-live{{end}}"></div>
              <div class="work-meta">
                <div class="work-title-row">
                  <span class="work-co">{{.Company}}</span>
                  <span class="work-role-txt">{{.Role}}</span>
                </div>
                <div class="work-dates-row">
                  <span class="work-when">{{.Dates}} &bull; {{.Location}}</span>
                  {{range .Badges}}<span class="co-badge {{.Class}}">{{.Label}}</span>{{end}}
                  {{if .BackedBy}}<span class="backed-note">{{.BackedBy}}</span>{{end}}
                </div>
              </div>
            </div>
            <p class="work-blurb">{{.Description}}</p>
          </div>
          {{end}}
        </div>
      </div>

      <!-- Footer links -->
      <div class="footer-links">
        {{range .Links}}
        <a href="{{.URL}}" class="f-link" id="{{.ID}}">&#8599; {{.Label}}</a>
        {{end}}
      </div>

      <div class="page-foot">
        <div class="ag-stamp">antigravity</div>
        <span class="pg-num">P. 01</span>
      </div>
    </div><!-- /page-left -->

    <!-- =============== RIGHT PAGE p.02 =============== -->
    <div class="page page-right" id="page-02">
      <div class="page-grid-overlay"></div>

      <!-- 2×3 Sticky notes grid -->
      <div class="notes-grid">

        <!-- Yellow: work -->
        <div class="snote yellow-n" data-id="note-work" tabindex="0" role="button" aria-label="work section p.03">
          <div class="sn-clip" aria-hidden="true">
            <svg width="22" height="44" viewBox="0 0 22 44" fill="none">
              <path d="M11 2C11 2 3 2 3 12 L3 32 C3 38 7 42 11 42 C15 42 19 38 19 32 L19 11 C19 8 17 6 15 6 C13 6 11 8 11 11 L11 31 C11 33 12 34 11 34" stroke="#9ca3af" stroke-width="2" stroke-linecap="round" fill="none"/>
            </svg>
          </div>
          <div class="sn-pg">p.03</div>
          <div class="sn-icon">
            <svg viewBox="0 0 40 40" fill="none" stroke="rgba(0,0,0,0.45)" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
              <rect x="8" y="14" width="24" height="18" rx="2"/>
              <path d="M14 14V11a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v3"/>
              <line x1="20" y1="20" x2="20" y2="26"/>
              <line x1="14" y1="23" x2="26" y2="23"/>
            </svg>
          </div>
          <div class="sn-title">work</div>
        </div>

        <!-- Peach: projects -->
        <div class="snote peach-n" data-id="note-projects" tabindex="0" role="button" aria-label="projects section p.05">
          <div class="sn-pg">p.05</div>
          <div class="sn-icon">
            <svg viewBox="0 0 40 40" fill="none" stroke="rgba(0,0,0,0.45)" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M6 12 L6 32 Q6 34 8 34 L32 34 Q34 34 34 32 L34 16 Q34 14 32 14 L20 14 L17 11 Q16 10 14 10 L8 10 Q6 10 6 12Z"/>
            </svg>
          </div>
          <div class="sn-title">projects</div>
        </div>

        <!-- Purple: what I know -->
        <div class="snote purple-n" data-id="note-skills" tabindex="0" role="button" aria-label="skills section p.07">
          <div class="sn-pg">p.07</div>
          <div class="sn-icon">
            <svg viewBox="0 0 40 40" fill="none" stroke="rgba(0,0,0,0.45)" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M20 6 C12 6 8 12 8 17 C8 22 11 25 14 27 L14 31 Q14 33 16 33 L24 33 Q26 33 26 31 L26 27 C29 25 32 22 32 17 C32 12 28 6 20 6Z"/>
              <line x1="16" y1="33" x2="24" y2="33"/>
              <line x1="17" y1="36" x2="23" y2="36"/>
              <line x1="20" y1="6" x2="20" y2="3"/>
            </svg>
          </div>
          <div class="sn-title">what I know</div>
        </div>

        <!-- Cyan: tech I use -->
        <div class="snote cyan-n" data-id="note-tech" tabindex="0" role="button" aria-label="tech section p.11">
          <div class="sn-pg">p.11</div>
          <div class="sn-icon">
            <svg viewBox="0 0 40 40" fill="none" stroke="rgba(0,0,0,0.45)" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
              <rect x="5" y="8" width="30" height="20" rx="2"/>
              <path d="M14 32 L26 32"/>
              <path d="M20 28 L20 32"/>
              <polyline points="11,16 16,20 11,24"/>
              <line x1="19" y1="24" x2="27" y2="24"/>
            </svg>
          </div>
          <div class="sn-title">tech I use</div>
        </div>

        <!-- Pink: notes -->
        <div class="snote pink-n" data-id="note-notes" tabindex="0" role="button" aria-label="notes section p.15">
          <div class="sn-pg">p.15</div>
          <div class="sn-icon">
            <svg viewBox="0 0 40 40" fill="none" stroke="rgba(0,0,0,0.45)" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
              <line x1="8" y1="14" x2="32" y2="14"/>
              <line x1="8" y1="20" x2="32" y2="20"/>
              <line x1="8" y1="26" x2="22" y2="26"/>
              <path d="M28 28 L34 22 L30 18 L24 24 L28 28Z"/>
              <path d="M24 24 L22 32 L28 28Z" fill="rgba(0,0,0,0.3)"/>
            </svg>
          </div>
          <div class="sn-title">notes</div>
        </div>

        <!-- White: say hi -->
        <div class="snote white-n" data-id="note-sayhi" tabindex="0" role="button" aria-label="say hi section p.17">
          <div class="sn-pg">p.17</div>
          <div class="sn-icon">
            <svg viewBox="0 0 40 40" fill="none" stroke="rgba(0,0,0,0.4)" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M8 32 L34 18 L22 34 L20 26 Z"/>
              <path d="M20 26 L22 34"/>
            </svg>
          </div>
          <div class="sn-title">say hi</div>
        </div>

      </div><!-- /notes-grid -->

      <!-- Bottom area: bust + annotation + wave -->
      <div class="page-bottom-row">
        <div class="bust-area">
          <img src="/static/images/marble_bust.png" alt="Classical marble bust" class="bust-img"/>
          <div class="pick-note-label">pick a note.</div>
        </div>
        <div class="wave-pol">
          <div class="wave-tape-strip"></div>
          <img src="/static/images/great_wave.png" alt="The Great Wave" class="wave-pol-img"/>
        </div>
      </div>

      <!-- Decorative pen clipped to right edge -->
      <div class="side-pen" aria-hidden="true">
        <div class="pen-top"></div>
        <div class="pen-barrel">
          <div class="pen-clip-tab"></div>
          <div class="pen-text">antigravity</div>
        </div>
        <div class="pen-grip-zone"></div>
        <div class="pen-nib"></div>
      </div>

      <div class="page-foot right-foot">
        <span class="pg-num">P. 02</span>
      </div>
    </div><!-- /page-right -->

    <!-- Right teal cover edge -->
    <div class="binder-cover-edge right-edge"></div>

  </div><!-- /binder-wrapper -->
</div><!-- /book-scene -->

<!-- Sticky note expand modal -->
<div class="sn-modal" id="snModal" role="dialog" aria-modal="true" aria-hidden="true">
  <div class="sn-backdrop" id="snBackdrop"></div>
  <div class="sn-card">
    <button class="sn-close" id="snClose" aria-label="Close">&times;</button>
    <div class="sn-card-pg" id="snPg"></div>
    <div class="sn-card-title" id="snTitle"></div>
    <div class="sn-card-body" id="snBody"></div>
  </div>
</div>

<script src="/static/script.js"></script>
</body>
</html>
`
