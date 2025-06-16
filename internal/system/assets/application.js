window.copy = function(selector, message) {
    let content = document.querySelector(selector).innerHTML;
    let element = document.querySelector(selector);
    if (element.tagName.toLowerCase() === 'textarea') {
      content = element.value;
    }

    navigator.clipboard.writeText(content);
    htmx.trigger("#flash", "htmx:notify",  {'message': message})
}


window.flash = function(elem, message) {
    elem.classList.remove('hidden');
    elem.innerHTML = message;
    setTimeout(() => {
        elem.classList.add('hidden');
    }, 3000);
}

document.addEventListener('DOMContentLoaded', () => {
  let el = document.querySelector("#MarkdownContent")
  if( el == undefined) {
    return
  }

  el.addEventListener("keydown", (e) => {
    if (e.key !== 'Tab') {
        return
    }

    e.preventDefault();
    const textarea = e.target;
    textarea.setRangeText(
        '\t',
        textarea.selectionStart,
        textarea.selectionStart,
        'end'
    );
  })


  el.addEventListener("scroll", (e) => {
    document.querySelector("#htmlcontainer").scrollTop = el.scrollTop;
  })
});


document.addEventListener("DOMContentLoaded", ()  => {
  let el = document.querySelector("#htmlcontainer")
  el.addEventListener("DOMSubtreeModified", (e) => {
    hljs.highlightAll();
  })

  el.addEventListener("scroll", (e) => {
    document.querySelector("#MarkdownContent").scrollTop = el.scrollTop;
  })
})
