//--------------------------------- load data table ------------------------------------------------|
let page = 1;
let orderAscDesc = "ASC";
let titleOrder = "id";
let filter;
const table = document.querySelector("#tablePermission");
const thead = table.querySelectorAll("thead td");
const tbody = table.querySelector("tbody");
function DeleteRows() {
    const max = tbody.rows.length;
    for (let i = 0; i < max; i++) {
        tbody.deleteRow(0);
    }
}
function Call() {
    // axios
    axios.get('/api/permission/', {
        params: {
            page,
            orderAscDesc,
            titleOrder,
            filter
        }
    }).then(response => {
        console.log(response);
    }).catch(error => {
        console.log(error.response.data);
    })
}
// load page
window.addEventListener('load',function () {
    DeleteRows();
    Call();
});