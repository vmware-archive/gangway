(function(){
    var clickCopy = function(e) {
        e.preventDefault();
        console.log(e.target.parentNode.parentNode);
        var el = e.target.parentNode.parentNode.querySelector(".active");
        var textarea = document.createElement("textarea");
        textarea.classList.add("visually-hidden");
        textarea.textContent = el.textContent;
        document.body.appendChild(textarea);
        textarea.select();
        document.execCommand("copy");
        textarea.remove();
        M.toast({html: "Code has been copied to clipboard."});
    };
    
    var btnCopy = document.querySelectorAll(".btn-copy");
    for (var i = 0; i < btnCopy.length; i++) {
        btnCopy[i].addEventListener("click", clickCopy)
    }
    
    var tabs = document.querySelectorAll(".tabs");
    for (var i = 0; i < tabs.length; i++) {
        M.Tabs.init(tabs[i]);
    }
})();