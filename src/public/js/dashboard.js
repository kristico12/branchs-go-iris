window.onload = function () {
    // assing year
    const year = document.querySelector("#year");
    year.innerHTML = new Date().getFullYear().toString();
}
function capitalize(word) {
    return word[0].toUpperCase() + word.slice(1);
}