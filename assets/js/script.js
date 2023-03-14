window.addEventListener("scroll", function() {
    var scrollEd = window.scrollY;

    if (scrollEd > 100) {
        document.getElementById("nav-vertical").style.display = "none";
        document.getElementById("small-nav").style.display = "flex";

    } else {
        document.getElementById("nav-vertical").style.display = "flex";
        document.getElementById("small-nav").style.display = "none";
    }
});

