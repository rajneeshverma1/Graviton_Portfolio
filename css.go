package main

// CSS holds the complete stylesheet embedded as a Go string constant.
const CSS = `
/* ====================================================
   GRAVITON PORTFOLIO — Pixel-perfect binder notebook
   ==================================================== */

@import url('https://fonts.googleapis.com/css2?family=Caveat:wght@400;600;700&family=Inter:wght@300;400;500;600;700;800;900&family=DM+Serif+Display:ital@0;1&family=Space+Mono:wght@400;700&family=Permanent+Marker&display=swap');

*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

:root {
  --desk:        #d4c89a;
  --desk-warm:   #c8bc8e;
  --cover-teal:  #5bc4be;
  --cover-dark:  #3da8a2;
  --paper:       #f7f3eb;
  --paper-l:     #f9f5ee;
  --grid:        rgba(0,170,160,0.13);
  --ink:         #1c1c1c;
  --red-name:    #c1272d;
  --green-live:  #22c55e;
  --note-yellow: #fde68a;
  --note-peach:  #fdba74;
  --note-purple: #ddd6fe;
  --note-cyan:   #a5f3fc;
  --note-pink:   #fca5a5;
  --note-white:  #f5f5ef;
  --font-hand:   'Caveat', cursive;
  --font-marker: 'Permanent Marker', cursive;
  --font-body:   'Inter', sans-serif;
  --font-mono:   'Space Mono', monospace;
  --font-serif:  'DM Serif Display', serif;
}

html, body {
  height: 100%; width: 100%;
  overflow-x: hidden;
  font-family: var(--font-body);
}

/* ── DESK BACKGROUND ─────────────────────────────── */
.desk-bg {
  position: fixed; inset: 0; z-index: 0;
  background:
    radial-gradient(ellipse at 15% 25%, #e8d8a0 0%, #d4c48a 35%, #c6b67a 70%),
    #c8bc8e;
  overflow: hidden;
}
.plant-shadow {
  position: absolute;
  top: -40px; left: -40px;
  width: 380px; height: 580px;
  pointer-events: none;
  opacity: 0.9;
}

/* ── BOOK SCENE & OPEN ANIMATION ─────────────────── */
.book-scene {
  position: relative; z-index: 1;
  min-height: 100vh;
  display: flex; align-items: center; justify-content: center;
  padding: 28px 12px;
  perspective: 2000px;
}

.binder-wrapper {
  display: flex;
  align-items: stretch;
  width: 100%; max-width: 1380px;
  min-height: 820px;
  filter:
    drop-shadow(0 30px 70px rgba(0,0,0,0.30))
    drop-shadow(0 6px 16px rgba(0,0,0,0.18));
  transform-origin: center center;
}

/* ── COVER EDGES (teal with grid texture) ────────── */
.binder-cover-edge {
  flex: 0 0 40px;
  background: var(--cover-teal);
  position: relative;
  overflow: hidden;
}
.left-edge  { border-radius: 10px 0 0 10px; }
.right-edge { border-radius: 0 10px 10px 0; }

/* Diamond grid pattern on cover */
.binder-cover-edge::before {
  content: '';
  position: absolute; inset: 0;
  background-image:
    linear-gradient(45deg, rgba(255,255,255,0.1) 25%, transparent 25%),
    linear-gradient(-45deg, rgba(255,255,255,0.1) 25%, transparent 25%),
    linear-gradient(45deg, transparent 75%, rgba(255,255,255,0.1) 75%),
    linear-gradient(-45deg, transparent 75%, rgba(255,255,255,0.1) 75%);
  background-size: 8px 8px;
  background-position: 0 0, 0 4px, 4px -4px, -4px 0px;
}
/* Light/dark edge shading */
.binder-cover-edge::after {
  content: '';
  position: absolute; inset: 0;
  background: linear-gradient(90deg,
    rgba(0,0,0,0.12) 0%, transparent 40%, transparent 60%, rgba(0,0,0,0.08) 100%
  );
}

/* ── BINDER SPINE (metallic rings) ───────────────── */
.binder-spine {
  flex: 0 0 48px;
  background: linear-gradient(90deg,
    #c8d0d8 0%, #dde4ea 20%, #eaeff3 45%, #d8e0e8 70%, #bec8d0 100%
  );
  box-shadow: inset -3px 0 8px rgba(0,0,0,0.12), inset 2px 0 6px rgba(0,0,0,0.08);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 48px;
  position: relative; z-index: 4;
}

.ring-group {
  display: flex; flex-direction: column; align-items: center; gap: 54px;
}

.ring {
  width: 38px; height: 38px;
  border-radius: 50%;
  background: linear-gradient(140deg, #e8ecf0 0%, #b8c4cc 35%, #d8e0e6 55%, #8898a8 80%, #c0ccd4 100%);
  box-shadow:
    0 3px 10px rgba(0,0,0,0.28),
    inset 0 1px 4px rgba(255,255,255,0.65),
    inset 0 -1px 3px rgba(0,0,0,0.22);
  display: flex; align-items: center; justify-content: center;
}

.ring-inner {
  width: 20px; height: 20px;
  border-radius: 50%;
  background: linear-gradient(140deg, #6a7a88 0%, #b0bcc8 50%, #7a8a98 100%);
  box-shadow: inset 0 2px 5px rgba(0,0,0,0.35), 0 1px 3px rgba(255,255,255,0.5);
}

/* ── PAGES ───────────────────────────────────────── */
.page {
  flex: 1;
  background: var(--paper);
  position: relative;
  padding: 28px 26px 16px;
  display: flex;
  flex-direction: column;
  min-height: 820px;
  overflow: hidden;
}

.page-grid-overlay {
  position: absolute; inset: 0; pointer-events: none;
  background-image:
    linear-gradient(var(--grid) 1px, transparent 1px),
    linear-gradient(90deg, var(--grid) 1px, transparent 1px);
  background-size: 22px 22px;
}

.page-left  { border-right: 1px solid rgba(0,0,0,0.05); background: var(--paper-l); }
.page-right { background: var(--paper); }

/* ── BOOK OPEN ANIMATION ─────────────────────────── */
.page-left {
  transform-origin: right center;
  animation: flipLeft 0.8s cubic-bezier(0.4,0,0.2,1) 0.2s both;
}
.page-right {
  transform-origin: left center;
  animation: flipRight 0.8s cubic-bezier(0.4,0,0.2,1) 0.2s both;
}
.binder-cover-edge.left-edge {
  animation: flipLeft 0.8s cubic-bezier(0.4,0,0.2,1) 0.2s both;
}
.binder-cover-edge.right-edge {
  animation: flipRight 0.8s cubic-bezier(0.4,0,0.2,1) 0.2s both;
}

@keyframes flipLeft {
  from { transform: perspective(1400px) rotateY(90deg); opacity: 0.3; }
  to   { transform: perspective(1400px) rotateY(0deg);  opacity: 1; }
}
@keyframes flipRight {
  from { transform: perspective(1400px) rotateY(-90deg); opacity: 0.3; }
  to   { transform: perspective(1400px) rotateY(0deg);   opacity: 1; }
}

/* ── POLAROID CLUSTER ────────────────────────────── */
.polaroid-cluster {
  position: relative;
  width: 100%;
  height: 195px;
  margin-bottom: 8px;
}

.polaroid {
  background: #fff;
  position: absolute;
  box-shadow: 0 6px 22px rgba(0,0,0,0.2), 0 2px 6px rgba(0,0,0,0.1);
}

.main-polaroid {
  left: 0; top: 0;
  width: 210px;
  padding: 8px 8px 32px;
  transform: rotate(-1.5deg);
  z-index: 2;
}

.mini-polaroid {
  left: 158px; top: 22px;
  width: 100px;
  padding: 5px 5px 22px;
  transform: rotate(4.5deg);
  z-index: 3;
}

.pol-img {
  width: 100%; height: 136px;
  object-fit: cover; display: block;
}

.mini-dark {
  width: 100%; height: 68px;
  background: linear-gradient(135deg, #0a0a18 0%, #161630 100%);
  display: flex; align-items: center; justify-content: center;
}
.mini-moon { font-size: 26px; }

.pol-label {
  font-family: var(--font-hand);
  font-size: 13px; font-weight: 600;
  color: #444; text-align: center;
  margin-top: 5px;
}

/* Tape strips */
.tape {
  position: absolute;
  background: rgba(215,205,155,0.6);
  border: 1px solid rgba(195,185,135,0.35);
  border-radius: 1px;
  z-index: 10;
}
.tape-l { width: 34px; height: 14px; top: -6px; left: 14px; transform: rotate(-7deg); }
.tape-r { width: 34px; height: 14px; top: -6px; right: 14px; transform: rotate(7deg); }
.tape-m { width: 24px; height: 11px; top: -5px; left: 50%; transform: translateX(-50%) rotate(-4deg); }

/* ── SKYLINE CAPTION ─────────────────────────────── */
.skyline-caption {
  font-family: var(--font-hand);
  font-size: 13px;
  color: #6a6358;
  line-height: 1.5;
  margin-bottom: 12px;
  padding-left: 2px;
}

.blink-today {
  color: var(--red-name);
  font-weight: 700;
  animation: blink 1.4s step-start infinite;
}
@keyframes blink {
  0%, 100% { opacity: 1; }
  50%       { opacity: 0; }
}

/* ── NAME BLOCK ──────────────────────────────────── */
.name-block { margin-bottom: 14px; }

.name-heading {
  font-size: clamp(30px, 3.8vw, 46px);
  line-height: 1.05;
  margin-bottom: 6px;
  letter-spacing: -0.5px;
}

.name-first {
  font-family: var(--font-body);
  font-weight: 900;
  color: var(--ink);
  letter-spacing: -1.5px;
}

.name-last {
  font-family: var(--font-serif);
  font-style: italic;
  color: var(--red-name);
  letter-spacing: 0px;
}

.bio-text {
  font-family: var(--font-body);
  font-size: 12px;
  color: #555;
  line-height: 1.6;
  margin-bottom: 8px;
}

.strike {
  text-decoration: line-through;
  color: #b0a090;
}

.bio-highlight {
  color: var(--ink);
  font-weight: 500;
}

.status-row {
  display: flex; align-items: center; gap: 7px;
  margin-top: 4px;
}

.live-dot {
  width: 9px; height: 9px; border-radius: 50%;
  background: var(--green-live);
  box-shadow: 0 0 0 3px rgba(34,197,94,0.22), 0 0 10px rgba(34,197,94,0.55);
  flex-shrink: 0;
  animation: livePulse 2s ease-in-out infinite;
}
@keyframes livePulse {
  0%, 100% { box-shadow: 0 0 0 3px rgba(34,197,94,0.2), 0 0 8px rgba(34,197,94,0.4); }
  50%       { box-shadow: 0 0 0 5px rgba(34,197,94,0.15), 0 0 16px rgba(34,197,94,0.65); }
}

.status-label {
  font-family: var(--font-mono);
  font-size: 9px;
  color: #444;
  letter-spacing: 0.4px;
}
.status-label strong { color: var(--ink); }

/* ── WORK SECTION ────────────────────────────────── */
.work-section { margin-bottom: 10px; flex: 1; }

.work-header {
  display: flex; align-items: center; gap: 9px;
  margin-bottom: 10px;
}

.work-title {
  font-family: var(--font-body);
  font-size: 22px; font-weight: 800;
  color: var(--ink);
  letter-spacing: -0.5px;
}

.work-count {
  font-family: var(--font-mono);
  font-size: 8.5px; color: #999;
  letter-spacing: 1.5px;
  border: 1px solid #d4d0c8;
  border-radius: 3px;
  padding: 1px 6px;
}

.classic-tag {
  font-family: var(--font-mono);
  font-size: 7.5px; font-weight: 700;
  color: #fff;
  background: #1c1c1c;
  border-radius: 3px;
  padding: 2px 6px;
  letter-spacing: 0.5px;
  line-height: 1.3;
  text-align: center;
  transform: rotate(-1.5deg);
  box-shadow: 0 2px 6px rgba(0,0,0,0.25);
}

.work-items { display: flex; flex-direction: column; gap: 12px; }

.work-entry {
  border-left: 2px solid #e2ddd6;
  padding-left: 11px;
  transition: border-color 0.2s;
}
.work-entry:hover { border-color: var(--cover-teal); }

.work-entry-top {
  display: flex; align-items: flex-start; gap: 7px;
  margin-bottom: 3px;
}

.work-dot {
  width: 7px; height: 7px; border-radius: 50%;
  background: #ccc; flex-shrink: 0; margin-top: 5px;
}
.dot-live {
  background: var(--green-live);
  box-shadow: 0 0 0 2px rgba(34,197,94,0.2);
  animation: livePulse 2s ease-in-out infinite;
}

.work-meta { flex: 1; }

.work-title-row {
  display: flex; align-items: baseline;
  flex-wrap: wrap; gap: 4px;
  margin-bottom: 2px;
}

.work-co {
  font-family: var(--font-body);
  font-size: 13px; font-weight: 700;
  color: var(--ink);
}

.work-role-txt {
  font-family: var(--font-body);
  font-size: 11px; color: #777;
}

.work-dates-row {
  display: flex; align-items: center;
  flex-wrap: wrap; gap: 4px;
}

.work-when {
  font-family: var(--font-mono);
  font-size: 8.5px; color: #aaa;
  letter-spacing: 0.2px;
}

.co-badge {
  font-family: var(--font-mono);
  font-size: 8px; font-weight: 700;
  border-radius: 3px; padding: 1px 5px;
}
.yc-badge    { background: #ff6600; color: #fff; }
.a-badge     { background: #1c1c1c; color: #fff; }
.g-badge     { background: #4285f4; color: #fff; }
.accel-badge { background: #6d28d9; color: #fff; }
.ant-badge   { background: #0f766e; color: #fff; }

.backed-note {
  font-size: 8px; color: #bbb;
  font-style: italic;
}

.work-blurb {
  font-family: var(--font-body);
  font-size: 11px; color: #666;
  line-height: 1.55;
  margin-left: 14px;
}

/* ── FOOTER LINKS ────────────────────────────────── */
.footer-links {
  display: flex; gap: 12px;
  flex-wrap: wrap;
  margin-top: 12px;
  padding-top: 6px;
}

.f-link {
  font-family: var(--font-hand);
  font-size: 15px; color: #555;
  text-decoration: none;
  transition: color 0.15s, transform 0.15s;
}
.f-link:hover { color: var(--cover-dark); transform: translateY(-1px); }

/* ── PAGE FOOT ───────────────────────────────────── */
.page-foot {
  display: flex; justify-content: space-between; align-items: center;
  padding-top: 8px; margin-top: 8px;
  border-top: 1px dashed #d8d2c8;
}
.right-foot { justify-content: flex-end; }

.ag-stamp {
  font-family: var(--font-hand);
  font-size: 11px; font-weight: 700;
  color: #fff; background: #1c1c1c;
  border-radius: 3px; padding: 2px 9px;
  letter-spacing: 1px;
}

.pg-num {
  font-family: var(--font-mono);
  font-size: 9.5px; color: #bbb;
  letter-spacing: 1.5px;
}

/* ── STICKY NOTES GRID ───────────────────────────── */
.notes-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px 12px;
  padding: 8px 4px 4px;
  flex: 1;
}

.snote {
  position: relative;
  padding: 12px 12px 14px;
  border-radius: 2px;
  cursor: pointer;
  user-select: none;
  outline: none;
  min-height: 118px;
  transition:
    transform 0.24s cubic-bezier(0.34,1.56,0.64,1),
    box-shadow 0.22s ease,
    filter 0.2s ease;
}

/* Individual note rotations & base shadows */
.yellow-n { background: var(--note-yellow); transform: rotate(-2deg);    box-shadow: 2px 4px 16px rgba(0,0,0,0.14), 0 1px 4px rgba(0,0,0,0.08); }
.peach-n  { background: var(--note-peach);  transform: rotate(1.5deg);   box-shadow: 2px 4px 16px rgba(0,0,0,0.12); }
.purple-n { background: var(--note-purple); transform: rotate(-0.8deg);  box-shadow: 2px 4px 16px rgba(0,0,0,0.14); }
.cyan-n   { background: var(--note-cyan);   transform: rotate(2deg);     box-shadow: 2px 4px 16px rgba(0,0,0,0.12); }
.pink-n   { background: var(--note-pink);   transform: rotate(-2.2deg);  box-shadow: 2px 4px 16px rgba(0,0,0,0.14); }
.white-n  { background: var(--note-white);  transform: rotate(1deg);     box-shadow: 2px 4px 16px rgba(0,0,0,0.10); border: 1px solid #e0ddd5; }

/* Top peel shadow */
.snote::after {
  content: ''; position: absolute;
  top: 0; left: 0; right: 0; height: 4px;
  background: rgba(0,0,0,0.05);
  border-radius: 2px 2px 0 0;
}

.snote:hover, .snote:focus-visible {
  transform: translateY(-7px) scale(1.05) rotate(0deg) !important;
  box-shadow: 0 18px 44px rgba(0,0,0,0.22), 0 4px 12px rgba(0,0,0,0.12) !important;
  filter: brightness(1.04); z-index: 12;
}

/* Paperclip on yellow note */
.sn-clip {
  position: absolute; top: -20px; left: 20px;
  z-index: 20;
  filter: drop-shadow(0 1px 3px rgba(0,0,0,0.2));
}

.sn-pg {
  font-family: var(--font-mono);
  font-size: 8px; color: rgba(0,0,0,0.32);
  letter-spacing: 1px; margin-bottom: 7px;
}

.sn-icon {
  width: 38px; height: 38px;
  margin-bottom: 4px;
}
.sn-icon svg { width: 100%; height: 100%; }

.sn-title {
  font-family: var(--font-hand);
  font-size: 21px; font-weight: 700;
  color: rgba(0,0,0,0.65);
  line-height: 1.1;
}

/* ── BOTTOM ROW (bust + wave) ────────────────────── */
.page-bottom-row {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  margin-top: 10px;
  padding: 0 4px;
}

.bust-area {
  display: flex; align-items: center; gap: 9px;
}

.bust-img {
  width: 66px; height: 66px;
  border-radius: 50%;
  object-fit: cover;
  border: 3px solid #fff;
  box-shadow: 0 4px 14px rgba(0,0,0,0.2);
  transform: rotate(-3deg);
  transition: transform 0.2s;
}
.bust-img:hover { transform: rotate(0deg) scale(1.06); }

.pick-note-label {
  font-family: var(--font-hand);
  font-size: 17px; font-weight: 600;
  color: var(--red-name);
  transform: rotate(-4deg);
  white-space: nowrap;
}

.wave-pol {
  position: relative;
  width: 124px;
  background: #fff;
  padding: 5px 5px 20px;
  transform: rotate(3deg);
  box-shadow: 0 6px 18px rgba(0,0,0,0.18);
}

.wave-tape-strip {
  position: absolute; top: -6px; left: 50%;
  transform: translateX(-50%) rotate(-2deg);
  width: 30px; height: 11px;
  background: rgba(215,205,155,0.6);
  border: 1px solid rgba(195,185,135,0.3);
}

.wave-pol-img { width: 100%; height: 88px; object-fit: cover; display: block; }

/* ── DECORATIVE PEN ──────────────────────────────── */
.side-pen {
  position: absolute;
  right: -16px; top: 50%;
  transform: translateY(-50%);
  display: flex; flex-direction: column;
  align-items: center; z-index: 8;
}

.pen-top {
  width: 13px; height: 32px;
  background: linear-gradient(90deg, #1a3a5c, #2563a8, #163055);
  border-radius: 3px 3px 0 0;
  box-shadow: inset -1px 0 3px rgba(0,0,0,0.3);
}

.pen-barrel {
  width: 12px; min-height: 110px;
  background: linear-gradient(90deg, #1a3a5c, #2563a8 35%, #3b82f6 55%, #163055);
  position: relative;
}

.pen-clip-tab {
  position: absolute; left: -5px; top: 6px;
  width: 5px; height: 65px;
  background: linear-gradient(90deg, #b0bcc8, #e8ecf0, #9aacb8);
  border-radius: 2px;
  box-shadow: -1px 0 4px rgba(0,0,0,0.2);
}

.pen-text {
  position: absolute; top: 50%; left: 50%;
  transform: translate(-50%,-50%) rotate(90deg);
  font-family: var(--font-hand); font-size: 7px;
  color: rgba(255,255,255,0.38); letter-spacing: 2px;
  white-space: nowrap;
}

.pen-grip-zone {
  width: 12px; height: 20px;
  background: linear-gradient(90deg, #0d1f33, #1a3a5c, #0d1f33);
  border-radius: 0 0 2px 2px;
}

.pen-nib {
  width: 0; height: 0;
  border-left: 6px solid transparent;
  border-right: 6px solid transparent;
  border-top: 16px solid #c8a84b;
  filter: drop-shadow(0 2px 2px rgba(0,0,0,0.2));
}

/* ── MODAL ───────────────────────────────────────── */
.sn-modal {
  position: fixed; inset: 0; z-index: 200;
  display: flex; align-items: center; justify-content: center;
  pointer-events: none; opacity: 0;
  transition: opacity 0.22s ease;
}
.sn-modal.open { pointer-events: all; opacity: 1; }

.sn-backdrop {
  position: absolute; inset: 0;
  background: rgba(15,25,15,0.48);
  backdrop-filter: blur(4px);
}

.sn-card {
  position: relative; z-index: 1;
  background: #fffef9;
  border-radius: 3px;
  padding: 34px 38px;
  min-width: 310px; max-width: 500px;
  box-shadow: 0 28px 70px rgba(0,0,0,0.28), 0 4px 14px rgba(0,0,0,0.14);
  transform: scale(0.86) rotate(-1.5deg);
  transition: transform 0.28s cubic-bezier(0.34,1.56,0.64,1);
}
.sn-modal.open .sn-card { transform: scale(1) rotate(0deg); }

.sn-close {
  position: absolute; top: 12px; right: 14px;
  background: none; border: none; cursor: pointer;
  font-size: 22px; color: #aaa; line-height: 1;
  transition: color 0.15s, transform 0.15s;
}
.sn-close:hover { color: #1c1c1c; transform: scale(1.2); }

.sn-card-pg { font-family: var(--font-mono); font-size: 9px; color: #bbb; letter-spacing: 1.5px; margin-bottom: 8px; }
.sn-card-title { font-family: var(--font-hand); font-size: 32px; font-weight: 700; color: #1c1c1c; margin-bottom: 14px; }
.sn-card-body { font-family: var(--font-body); font-size: 13px; color: #555; line-height: 1.65; }
.sn-card-body ul { margin-left: 18px; margin-top: 6px; }
.sn-card-body li { margin-bottom: 5px; }
.sn-card-body strong { color: #1c1c1c; }

/* ── SCROLLBAR ───────────────────────────────────── */
::-webkit-scrollbar { width: 5px; }
::-webkit-scrollbar-thumb { background: rgba(0,0,0,0.13); border-radius: 3px; }

`
// book-open anim v2
// warm sandy desk v2
// teal cover grid v2
// metallic rings v2
// polaroid cluster v2
// name typography v2
// bio strikethrough v2
// live pulse dot v2
// work section v2
// snote grid v2
// snote hover v2
// pick note label v2
// wave polaroid v2
// modal spring v2
// blink today v2
/* book-open flip v2 */
/* desk radial gradient v2 */
/* diamond grid texture v2 */
/* ring gradient v2 */
