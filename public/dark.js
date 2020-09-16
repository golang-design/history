/* 
  Copyright 2020 The golang.design Initiative.
  All rights reserved. Use of this source code
  is governed by a MIT license that can be found
  in the LICENSE file.
 */

const darkSwitch = document.getElementById('darkSwitch')
window.addEventListener('load', () => {
  if (!darkSwitch) { return }

  // init
  const darkThemeSelected =
    localStorage.getItem('darkSwitch') !== null &&
    localStorage.getItem('darkSwitch') === 'dark';
  darkSwitch.checked = darkThemeSelected;
  darkThemeSelected ? document.body.setAttribute('data-theme', 'dark') :
    document.body.removeAttribute('data-theme');

  // reset
  darkSwitch.addEventListener('change', () => {
    if (darkSwitch.checked) {
      document.body.setAttribute('data-theme', 'dark');
      localStorage.setItem('darkSwitch', 'dark');
    } else {
      document.body.removeAttribute('data-theme');
      localStorage.removeItem('darkSwitch');
    }
  })
})
