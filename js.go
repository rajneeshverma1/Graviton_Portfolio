package main

// JS holds the client-side interactivity as a Go string constant.
const JS = `
/* Graviton Portfolio — script.js */
(function () {
  'use strict';

  const NOTE_DATA = {
    'note-work': {
      title: 'work', pg: 'p. 03',
      content: ` + "`" + `<p>Three startups, <strong>all remote</strong>, all real products.</p>
        <ul>
          <li><strong>Stealth Startup</strong> (Sep 2026 → Now) — Architecting Go microservices. 1,800+ commits in month one.</li>
          <li><strong>Amboras</strong> (Feb → Aug 2026, SF) — High-throughput Go APIs, 30k+ users, YC & A* backed.</li>
          <li><strong>Dodge AI</strong> (Aug → Oct 2025, Bengaluru) — Full Go backend platform. Google, Accel & Antler backed.</li>
        </ul>` + "`" + `
    },
    'note-projects': {
      title: 'projects', pg: 'p. 05',
      content: ` + "`" + `<p>Things I built because they didn't exist yet.</p>
        <ul>
          <li><strong>graviton</strong> — Lightweight Go HTTP framework with observability built-in.</li>
          <li><strong>pgstream</strong> — Real-time PostgreSQL → Redis sync daemon in Go.</li>
          <li><strong>dropzone-cli</strong> — Zero-config Go deploy pipeline for indie devs.</li>
        </ul>` + "`" + `
    },
    'note-skills': {
      title: 'what I know', pg: 'p. 07',
      content: ` + "`" + `<ul>
          <li><strong>Go concurrency</strong> — goroutines, channels, sync, race detection</li>
          <li><strong>Distributed systems</strong> — consensus, eventual consistency, CAP</li>
          <li><strong>REST & gRPC APIs</strong> — protobuf, interceptors, versioning</li>
          <li><strong>Databases</strong> — PostgreSQL, Redis, SQLite, migrations</li>
          <li><strong>Infrastructure</strong> — Docker, Kubernetes, Terraform, CI/CD</li>
          <li><strong>Observability</strong> — OpenTelemetry, Prometheus, Grafana</li>
        </ul>` + "`" + `
    },
    'note-tech': {
      title: 'tech I use', pg: 'p. 11',
      content: ` + "`" + `<ul>
          <li>🐹 <strong>Go (Golang)</strong> — primary language, daily driver</li>
          <li>🐘 <strong>PostgreSQL</strong> — relational DB of choice</li>
          <li>⚡ <strong>Redis</strong> — caching, pub/sub, distributed locks</li>
          <li>🐳 <strong>Docker</strong> — containerising everything</li>
          <li>☸️ <strong>Kubernetes</strong> — orchestration at scale</li>
          <li>🔧 <strong>gRPC + protobuf</strong> — service comms</li>
          <li>📡 <strong>OpenTelemetry</strong> — distributed tracing</li>
        </ul>` + "`" + `
    },
    'note-notes': {
      title: 'notes', pg: 'p. 15',
      content: ` + "`" + `<ul>
          <li>Most startups ship monoliths disguised as microservices.</li>
          <li>The 80/20 of Kubernetes — you probably need 20% of it.</li>
          <li>Go's simplicity is its superpower. Don't abstract everything.</li>
          <li>Writing is thinking. Write more docs, even for yourself.</li>
        </ul>` + "`" + `
    },
    'note-sayhi': {
      title: 'say hi', pg: 'p. 17',
      content: ` + "`" + `<p>I read every message. Drop me a line:</p>
        <ul>
          <li>📧 <strong>Email</strong> — sanatan@example.com</li>
          <li>🐙 <strong>GitHub</strong> — github.com/sanatan-sharma</li>
          <li>💼 <strong>LinkedIn</strong> — linkedin.com/in/sanatan-sharma</li>
          <li>🐦 <strong>X</strong> — @sanatan_dev</li>
        </ul>
        <p style="margin-top:10px;font-style:italic;color:#999;font-size:11px;">Best for: Go/backend collabs, consulting, or geeking about distributed systems.</p>` + "`" + `
    }
  };

  const modal    = document.getElementById('snModal');
  const backdrop = document.getElementById('snBackdrop');
  const closeBtn = document.getElementById('snClose');
  const pgEl     = document.getElementById('snPg');
  const titleEl  = document.getElementById('snTitle');
  const bodyEl   = document.getElementById('snBody');
  const card     = modal.querySelector('.sn-card');

  const NOTE_COLORS = {
    'note-work':     '#fef9c3',
    'note-projects': '#fed7aa',
    'note-skills':   '#ede9fe',
    'note-tech':     '#cffafe',
    'note-notes':    '#fce7f3',
    'note-sayhi':    '#f5f5ef',
  };

  function openModal(el) {
    const id   = el.dataset.id;
    const data = NOTE_DATA[id];
    if (!data) return;
    pgEl.textContent    = data.pg;
    titleEl.textContent = data.title;
    bodyEl.innerHTML    = data.content;
    card.style.background = NOTE_COLORS[id] || '#fffef9';
    modal.setAttribute('aria-hidden', 'false');
    modal.classList.add('open');
    document.body.style.overflow = 'hidden';
    closeBtn.focus();
  }

  function closeModal() {
    modal.classList.remove('open');
    modal.setAttribute('aria-hidden', 'true');
    document.body.style.overflow = '';
  }

  document.querySelectorAll('.snote').forEach(function (n) {
    n.addEventListener('click', function () { openModal(n); });
    n.addEventListener('keydown', function (e) {
      if (e.key === 'Enter' || e.key === ' ') { e.preventDefault(); openModal(n); }
    });
  });

  closeBtn.addEventListener('click', closeModal);
  backdrop.addEventListener('click', closeModal);
  document.addEventListener('keydown', function (e) { if (e.key === 'Escape') closeModal(); });

  /* Subtle parallax tilt */
  const binder  = document.getElementById('binder');
  let   ticking = false;
  document.addEventListener('mousemove', function (e) {
    if (ticking) return;
    ticking = true;
    requestAnimationFrame(function () {
      const rx = ((e.clientY / window.innerHeight) - 0.5) * 2.5;
      const ry = -((e.clientX / window.innerWidth)  - 0.5) * 2.5;
      binder.style.transform = ` + "`" + `perspective(2000px) rotateX(${rx}deg) rotateY(${ry}deg)` + "`" + `;
      ticking = false;
    });
  });
  document.addEventListener('mouseleave', function () {
    binder.style.transform = '';
  });

})();

`
