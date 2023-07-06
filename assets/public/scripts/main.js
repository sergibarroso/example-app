function ready(fn) {
  if (document.attachEvent ? document.readyState === 'complete' : document.readyState !== 'loading') {
    fn();
  } else {
    document.addEventListener('DOMContentLoaded', fn);
  }
}

ready(function () {
  let jsContainer = document.getElementsByClassName('js-container')[0];
  let p = document.createElement('p');
  p.classList.add('slogan');
  p.classList.add('slogan__bg');
  p.classList.add('slogan__subtitle');
  p.appendChild(text);
  jsContainer.appendChild(p);
});
