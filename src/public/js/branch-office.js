//---------------------------------- Global Data------------------------------------|
let branchOffice;
let page = 1;
let orderAscDesc = "ASC";
let titleOrder = "id";
let pagination;
const table = document.querySelector("#tableBranchOffice");
const thead = table.querySelectorAll("thead td");
const tbody = table.querySelector("tbody");
const showError = document.querySelector("#errorList");
//--------------------------------- functions ------------------------------------------------|
function CreatePaginate(info) {
    pagination = new tui.Pagination(document.querySelector("#pagination"), {
        totalItems: info.Filtered,
        itemsPerPage: info.NumberForPage,
        page: info.Page,
        visiblePages: 3,
        centerAlign: true
    });
}
function DeleteRows() {
    const max = tbody.rows.length;
    for (let i = 0; i < max; i++) {
        tbody.deleteRow(0);
    }
}
function ClearError() {
    showError.innerHTML = "";
}

function Call() {
    // axios
    axios.get('/api/branch_office/', {
        params: {
            branchOffice,
            page,
            orderAscDesc,
            titleOrder
        }
    }).then(response => {
        const data = response.data;
        data.Data.forEach((value, i) => {
            const row = tbody.insertRow(i);
            const keyDate = Object.keys(value)
            for (let j = 0; j < thead.length - 1; j++) {
                row.insertCell(j).innerHTML = value[keyDate[j]];
            }
        })
        CreatePaginate(data);
    }).catch(error => {
        showError.innerHTML = error;
    })
}
// order by
document.querySelector("#tableBranchOffice thead tr").addEventListener('click', function (e) {
    if (e.target.id !== "") {
        page = 1;
        titleOrder = e.target.id;
        orderAscDesc = orderAscDesc === "ASC" ? "DESC" : "ASC";
        DeleteRows();
        ClearError();
        Call();
    }
});
//-------------------------------------------------------- CODE FOR EDIT BRANCH --------------------------------------------------------|
// load page
window.addEventListener('load',function () {
    // init for table
    branchOffice = location.pathname.split("/").filter(value => value.length > 0)[0];
    DeleteRows();
    Call();
});