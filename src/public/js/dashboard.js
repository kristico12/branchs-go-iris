window.onload = function () {
    // assing year
    const year = document.querySelector("#year");
    year.innerHTML = new Date().getFullYear().toString();
}