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
