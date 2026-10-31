/* Copyright 2020 The golang.design Initiative.
   All rights reserved. Use of this source code
   is governed by a MIT license that can be found
   in the LICENSE file.

   Theme, menu and search for golang.design/history. */

(function () {
  'use strict';

  var root = document.documentElement;
  var body = document.body;

  // Theme: follow the system (auto), or force light or dark.
  var modes = ['auto', 'light', 'dark'];
  document.querySelectorAll('[data-theme-toggle]').forEach(function (b) {
    b.addEventListener('click', function () {
      var next = modes[(modes.indexOf(root.dataset.theme || 'auto') + 1) % modes.length];
      try {
        if (next === 'auto') localStorage.removeItem('theme');
        else localStorage.setItem('theme', next);
        localStorage.removeItem('darkSwitch'); // setting of the old site
      } catch (e) {}
      if (next === 'auto') delete root.dataset.theme;
      else root.dataset.theme = next;
      b.setAttribute('aria-label', b.dataset.label + ': ' + next);
    });
  });

  // Sidebar as a drawer on narrow screens.
  var side = document.querySelector('.side');
  document.querySelectorAll('[data-menu]').forEach(function (b) {
    b.addEventListener('click', function () {
      b.setAttribute('aria-expanded', body.classList.toggle('nav-open'));
    });
  });
  document.addEventListener('click', function (e) {
    if (!body.classList.contains('nav-open')) return;
    if (e.target.closest('[data-menu]')) return;
    if (!e.target.closest('.side') || e.target.closest('a')) body.classList.remove('nav-open');
  });

  // Keep the current page in view in the sidebar.
  var cur = side && side.querySelector('[aria-current="page"]');
  if (cur) side.scrollTop = Math.max(0, cur.offsetTop - side.clientHeight / 3);

  // Search over every entry of the document, loaded on first use.
  var q = document.getElementById('q');
  var results = document.getElementById('results');
  var toc = document.getElementById('toc');
  if (!q || !results || !toc) return;
  var base = body.dataset.base || '';
  var index = null;
  var loading = null;
  var max = 50;

  function load() {
    if (!loading) {
      loading = fetch(q.dataset.index)
        .then(function (r) { return r.json(); })
        .then(function (d) { index = d; });
    }
    return loading;
  }

  function el(tag, cls, text) {
    var e = document.createElement(tag);
    if (cls) e.className = cls;
    if (text) e.textContent = text;
    return e;
  }

  function run() {
    var terms = q.value.toLowerCase().split(/\s+/).filter(Boolean);
    if (!terms.length || !index) {
      results.hidden = true;
      toc.hidden = false;
      return;
    }
    var found = [];
    var total = 0;
    index.items.forEach(function (it) {
      var sec = index.sections[it[3]];
      var hay = (it[0] + ' ' + it[1] + ' ' + sec[2] + ' ' + index.pages[sec[0]][1]).toLowerCase();
      for (var i = 0; i < terms.length; i++) {
        if (hay.indexOf(terms[i]) < 0) return;
      }
      total++;
      if (found.length < max) found.push(it);
    });
    results.textContent = '';
    var count = el('li', 'count', total > max
      ? q.dataset.more.replace('%d', max).replace('%n', total)
      : q.dataset.count.replace('%d', total));
    results.appendChild(count);
    found.forEach(function (it) {
      var sec = index.sections[it[3]];
      var page = index.pages[sec[0]];
      var a = el('a');
      a.href = base + page[0] + (sec[1] ? '#' + sec[1] : '');
      if (it[0]) a.appendChild(el('span', 'ref', it[0]));
      a.appendChild(document.createTextNode(it[1]));
      a.appendChild(el('small', '', sec[2] && sec[2] !== page[1] ? page[1] + ' › ' + sec[2] : page[1]));
      var li = el('li');
      li.appendChild(a);
      results.appendChild(li);
    });
    results.hidden = false;
    toc.hidden = true;
  }

  q.addEventListener('focus', load);
  q.addEventListener('input', function () { load().then(run); });
  q.addEventListener('keydown', function (e) {
    if (e.key === 'Escape') {
      q.value = '';
      run();
      q.blur();
    } else if (e.key === 'Enter') {
      var first = results.querySelector('a');
      if (first) first.click();
    }
  });
  document.addEventListener('keydown', function (e) {
    var t = e.target.tagName;
    if (e.key !== '/' || e.metaKey || e.ctrlKey || t === 'INPUT' || t === 'TEXTAREA' || t === 'SELECT') return;
    e.preventDefault();
    if (getComputedStyle(side).transform !== 'none' && window.innerWidth <= 900) body.classList.add('nav-open');
    q.focus();
  });
})();
