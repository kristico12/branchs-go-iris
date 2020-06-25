//---------------------------------- Global Data------------------------------------|
let branchOffice;
let page = 1;
let orderAscDesc = "ASC";
let titleOrder = "id";
let filter;
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
    // event pagination
    pagination.on('beforeMove', function (e) {
        page = e.page;
        DeleteRows();
        ClearError();
        Call();
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
            titleOrder,
            filter
        }
    }).then(response => {
        const data = response.data;
        data.Data.forEach((value, i) => {
            const row = tbody.insertRow(i);
            const keyDate = Object.keys(value);
            for (let j = 0; j < thead.length - 1; j++) {
                row.insertCell(j).innerHTML = value[keyDate[j]];
            }
            row.insertCell(thead.length - 1).innerHTML = '<div class="uk-flex uk-flex-around">' +
                '<i uk-toggle="target: #modal-edit-branch-Office" class="icon" uk-icon="icon: pencil; ratio: 1.3" onclick="OpenEdit(\''+i+'\')"></i>' +
                '<i class="icon" uk-icon="icon: minus-circle; ratio: 1.3"></i>' +
                '</div>';
        });
        CreatePaginate(data);
    }).catch(error => {
        if (error.response.data.message) {
            showError.innerHTML = error.response.data.message;
        } else {
            showError.innerHTML = "A ocurrido un error Intente mas tarde";
        }
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
// search
document.querySelector("#btnSearch").addEventListener('click', function () {
    const infoSearch = document.querySelector("#search");
    page = 1;
    filter = infoSearch.value;
    DeleteRows();
    ClearError();
    Call();
});
//-------------------------------------------------------- CODE FOR EDIT BRANCH --------------------------------------------------------|
let selectEditProvince;
let selectEditCity;
function OpenEdit(i) {
    const row = tbody.rows.item(i);
    let editBranchOffice = {};
    for (const j in row.cells) {
        const cell = row.cells[j];
        if (cell.nodeName === "TD" && cell.textContent.length > 0) {
            switch (j) {
                case "0":
                    editBranchOffice.id = cell.textContent;
                    break;
                case "1":
                    editBranchOffice.city = cell.textContent;
                    break;
                case "2":
                    editBranchOffice.province = cell.textContent;
                    break;
                case "3":
                    editBranchOffice.address = cell.textContent;
                    break;
                case "4":
                    editBranchOffice.checkInTime = cell.textContent;
                    break;
                case "5":
                    editBranchOffice.exitTime = cell.textContent;
                    break;
                default:
                    break;
            }
        }
    }
    // send info to panel edit
    document.querySelector("#inputEditId").value = editBranchOffice.id;
    SetDataSelectProvince("#edit-select-province");
    document.querySelector("#inputEditAddress").value = editBranchOffice.address;
    document.querySelector("#inputEditEntranceHour").value = moment(editBranchOffice.checkInTime,["h:mm:ss A"]).format('HH:mm:ss');
    document.querySelector("#inputEditExitHour").value = moment(editBranchOffice.exitTime, ["h:mm:ss A"]).format('HH:mm:ss');
}
function HandleOptionEditSelectProvince() {
    const selected = selectEditProvince.selected();
    if (selected !== "") {
        ClearDataCities("#edit-select-city option");
        SetDataSelectCity(selected, "#edit-select-city")
    }
}
// load page
window.addEventListener('load',function () {
    branchOffice = location.pathname.split("/").filter(value => value.length > 0)[0];
    DeleteRows();
    Call();
    // init select
    selectEditProvince = new SlimSelect({
        select: '#edit-select-province'
    });
    selectEditCity = new SlimSelect({
        select: '#edit-select-city'
    });
});