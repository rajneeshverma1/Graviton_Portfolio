/* ============================================================
   GRAVITON PORTFOLIO — script.js
   Sticky note interactivity & modal system
   ============================================================ */

(function () {
  'use strict';

  // ── Modal content data keyed by note ID ──────────────────
  const NOTE_DATA = {
    'note-work': {
      title: 'work',
      icon: '💼',
      content: `
        <p>Three startups, <strong>all remote</strong>, all building real products that ship.</p>
        <ul>
          <li><strong>Stealth Startup</strong> (Sep 2026 → Now) — Architecting Go microservices from scratch. 1,800+ commits in month one.</li>
          <li><strong>Amboras</strong> (Feb → Aug 2026, SF) — High-throughput Go APIs, 30k+ users, backed by YC & A*.</li>
          <li><strong>Dodge AI</strong> (Aug → Oct 2025, Bengaluru) — Full backend platform in Go for enterprise demos. Google, Accel & Antler backed.</li>
        </ul>
      `,
    },
    'note-projects': {
      title: 'projects',
      icon: '📁',
      content: `
        <p>Open-source tools, side bets, and things I built because they didn't exist yet.</p>
        <ul>
          <li><strong>graviton</strong> — A lightweight Go HTTP framework with built-in observability.</li>
          <li><strong>pgstream</strong> — Real-time PostgreSQL → Redis sync daemon in Go.</li>
          <li><strong>dropzone-cli</strong> — Zero-config Go-based deploy pipeline for indie devs.</li>
        </ul>
        <p style="margin-top:10px; color:#888; font-size:11px;">More on GitHub →</p>
      `,
    },
    'note-skills': {
      title: 'what I know',
      icon: '💡',
      content: `
        <ul>
          <li><strong>Go concurrency</strong> — goroutines, channels, sync primitives, race detection.</li>
          <li><strong>Distributed systems</strong> — consensus, eventual consistency, CAP trade-offs.</li>
          <li><strong>REST & gRPC APIs</strong> — schema design, versioning, protobuf, interceptors.</li>
          <li><strong>Databases</strong> — PostgreSQL (advanced), Redis, SQLite, schema migrations.</li>
          <li><strong>Infrastructure</strong> — Docker, Kubernetes, Terraform, CI/CD pipelines.</li>
          <li><strong>Observability</strong> — OpenTelemetry, Prometheus, Grafana, structured logging.</li>
        </ul>
      `,
    },
    'note-tech': {
      title: 'tech I use',
      icon: '💻',
      content: `
        <ul>
          <li>🐹 <strong>Go (Golang)</strong> — primary language, daily driver</li>
          <li>🐘 <strong>PostgreSQL</strong> — relational DB of choice</li>
          <li>⚡ <strong>Redis</strong> — caching, pub/sub, distributed locks</li>
          <li>🐳 <strong>Docker</strong> — containerising everything</li>
          <li>☸️  <strong>Kubernetes</strong> — orchestration at scale</li>
          <li>🔧 <strong>gRPC + protobuf</strong> — service-to-service comms</li>
          <li>🌐 <strong>Next.js</strong> — frontend when I have to</li>
          <li>📡 <strong>OpenTelemetry</strong> — distributed tracing</li>
        </ul>
      `,
    },
    'note-notes': {
      title: 'notes',
      icon: '✏️',
      content: `
        <p>Things I'm thinking about:</p>
        <ul>
          <li>Why most startups still ship monoliths masquerading as microservices.</li>
          <li>The 80/20 of Kubernetes: you probably only need 20% of it.</li>
          <li>Go's simplicity is its superpower — resist the urge to abstract everything.</li>
          <li>Writing is thinking. Write more docs, even for yourself.</li>
        </ul>
      `,
    },
    'note-sayhi': {
      title: 'say hi',
      icon: '✈️',
      content: `
        <p>I read every message. Drop me a line:</p>
        <ul>
          <li>📧 <strong>Email</strong> — sanatan@example.com</li>
          <li>🐙 <strong>GitHub</strong> — github.com/sanatan-sharma</li>
          <li>💼 <strong>LinkedIn</strong> — linkedin.com/in/sanatan-sharma</li>
          <li>🐦 <strong>X / Twitter</strong> — @sanatan_dev</li>
        </ul>
        <p style="margin-top:12px; font-style:italic; color:#888;">Best for: collaborations, Go/backend consulting, or just geeking out about distributed systems.</p>
      `,
    },
  };

  // ── DOM refs ──────────────────────────────────────────────
  const modal        = document.getElementById('noteModal');
  const backdrop     = document.getElementById('modalBackdrop');
  const closeBtn     = document.getElementById('modalClose');
  const modalPageNum = document.getElementById('modalPageNum');
  const modalIcon    = document.getElementById('modalIcon');
  const modalTitle   = document.getElementById('modalTitle');
  const modalContent = document.getElementById('modalContent');

  // ── Open modal ────────────────────────────────────────────
  function openModal(noteEl) {
    const id   = noteEl.id;
    const page = noteEl.dataset.page;
    const data = NOTE_DATA[id];
    if (!data) return;

    modalPageNum.textContent = page;
    modalIcon.textContent    = data.icon;
    modalTitle.textContent   = data.title;
    modalContent.innerHTML   = data.content;

    // Tint card to match note colour
    const card = modal.querySelector('.note-modal-card');
    const colours = {
      'note-work':     '#fef9c3',
      'note-projects': '#fed7aa',
      'note-skills':   '#ede9fe',
      'note-tech':     '#cffafe',
      'note-notes':    '#fce7f3',
      'note-sayhi':    '#f5f5f0',
    };
    card.style.background = colours[id] || '#fffef8';

    modal.setAttribute('aria-hidden', 'false');
    modal.classList.add('open');
    document.body.style.overflow = 'hidden';
    closeBtn.focus();
  }

  // ── Close modal ───────────────────────────────────────────
  function closeModal() {
    modal.classList.remove('open');
    modal.setAttribute('aria-hidden', 'true');
    document.body.style.overflow = '';
  }

  // ── Bind sticky notes ─────────────────────────────────────
  document.querySelectorAll('.sticky-note').forEach(function (note) {
    note.addEventListener('click', function () { openModal(note); });
    note.addEventListener('keydown', function (e) {
      if (e.key === 'Enter' || e.key === ' ') {
        e.preventDefault();
        openModal(note);
      }
    });
  });

  // ── Close triggers ────────────────────────────────────────
  closeBtn.addEventListener('click', closeModal);
  backdrop.addEventListener('click', closeModal);
  document.addEventListener('keydown', function (e) {
    if (e.key === 'Escape') closeModal();
  });

  // ── Subtle cursor-parallax tilt on the binder ─────────────
  const binder = document.querySelector('.binder-wrapper');
  let ticking  = false;

  document.addEventListener('mousemove', function (e) {
    if (ticking) return;
    ticking = true;
    requestAnimationFrame(function () {
      const cx  = window.innerWidth  / 2;
      const cy  = window.innerHeight / 2;
      const dx  = (e.clientX - cx) / cx;   // -1 … 1
      const dy  = (e.clientY - cy) / cy;
      const rx  =  dy * 1.8;
      const ry  = -dx * 1.8;
      binder.style.transform = `perspective(1800px) rotateX(${rx}deg) rotateY(${ry}deg)`;
      ticking = false;
    });
  });

  document.addEventListener('mouseleave', function () {
    binder.style.transform = 'perspective(1800px) rotateX(0deg) rotateY(0deg)';
  });

  // ── Blink animation for "the blinking one is today" ───────
  (function blinkWindow() {
    const quote = document.querySelector('.skyline-quote');
    if (!quote) return;
    const text  = quote.innerHTML;
    // Wrap "today" with a blinking span
    quote.innerHTML = text.replace('today.', '<span class="blink-word">today.</span>');
    const style = document.createElement('style');
    style.textContent = `
      .blink-word {
        animation: blink-anim 1.4s step-start infinite;
        color: #c0392b;
        font-weight: 700;
      }
      @keyframes blink-anim {
        0%, 100% { opacity: 1; }
        50%       { opacity: 0; }
      }
    `;
    document.head.appendChild(style);
  })();

})();
// v2: modal system
// v3: parallax
// v3: blink anim
