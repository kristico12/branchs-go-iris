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
function GetDataTable(i) {
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
    return editBranchOffice;
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
                '<i class="icon" uk-icon="icon: minus-circle; ratio: 1.3" onclick="DeleteBranch(\''+i+'\')"></i>' +
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
const dialogTimeEditEntranceHour = new mdDateTimePicker.default({
    type: 'time',
    trigger: document.querySelector("#inputEditEntranceHour")
});
const dialogTimeEditExitHour = new mdDateTimePicker.default({
    type: 'time',
    trigger: document.querySelector("#inputEditExitHour")
});
function OpenEdit(i) {
    let editBranchOffice = Object.assign({},GetDataTable(i));
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
document.querySelector("#dialogEditOpenEntrance").addEventListener('click', function () {
    dialogTimeEditEntranceHour.toggle();
});
document.querySelector("#inputEditEntranceHour").addEventListener('onOk', function () {
    this.value = dialogTimeEditEntranceHour.time.set({second:0}).format('HH:mm:ss');
});
document.querySelector("#dialogEditOpenExit").addEventListener('click', function () {
    dialogTimeEditExitHour.toggle();
});
document.querySelector("#inputEditExitHour").addEventListener('onOk', function () {
    this.value = dialogTimeEditExitHour.time.set({second:0}).format('HH:mm:ss');
});
document.querySelector("#editBranchOffice").addEventListener('click', function () {
    const editBranchOffice = {
        id: document.querySelector("#inputEditId").value,
        city: selectEditCity.selected(),
        province: selectEditProvince.selected(),
        address: document.querySelector("#inputEditAddress").value,
        checkInTime: document.querySelector("#inputEditEntranceHour").value,
        exitTime: document.querySelector("#inputEditExitHour").value
    }
    // remove error geenral
    document.querySelector(`#messageEdit`).textContent = "";
    // remove info errors
    Object.keys(editBranchOffice).forEach(value => {
        if (value !== 'id') {
            document.querySelector(`#errorEdit${capitalize(value)}`).textContent = null;
        }
    });
    // show Loadin
    document.querySelector("#loadingEdit").classList.remove("uk-hidden");
    axios({
        method: "PUT",
        url: '/api/branch_office/',
        headers: {
            'Content-Type': 'application/json',
        },
        data: editBranchOffice
    })
        .then(() => {
            DeleteRows();
            ClearError();
            Call();
            // quit loading
            document.querySelector("#loadingEdit").classList.add("uk-hidden");
            // close modal
            UIkit.modal("#modal-edit-branch-Office").hide();
        })
        .catch(error => {
            const infoError = error.response;
            if (infoError.status === 400) {
                const errorData = infoError.data;
                if (Object.keys(errorData).includes("errors")) {
                    for (const customError of errorData.errors) {
                        const setError = document.querySelector(`#errorEdit${capitalize(customError.Key)}`);
                        setError.textContent = customError.Value;
                    }
                }
                if (Object.keys(errorData).includes("message")) {
                    const setError = document.querySelector(`#messageEdit`);
                    setError.textContent = errorData.message;
                }
            }
            // quit loading
            document.querySelector("#loadingEdit").classList.add("uk-hidden");
        })
});
//------------------------------------------ code for Delete branch ----------------------------------------------------|
function DeleteBranch(i) {
    let deleteBranchOffice = Object.assign({},GetDataTable(i));
    deleteBranchOffice.id = parseInt(deleteBranchOffice.id);
    // clear message error
    showError.innerHTML = "";
    UIkit.modal.confirm("Esta seguro que desea eliminar a "+deleteBranchOffice.address+" en "+deleteBranchOffice.city)
        .then(function () {
            axios({
                method: "DELETE",
                url: '/api/branch_office/',
                headers: {
                    'Content-Type': 'application/json',
                },
                data: deleteBranchOffice
            })
                .then(() => {
                    page = 1;
                    DeleteRows();
                    ClearError();
                    Call();
                })
                .catch(error => {
                    if (error.response.data.message) {
                        showError.innerHTML = error.response.data.message;
                    } else {
                        showError.innerHTML = "A ocurrido un error Intente mas tarde";
                    }
                })
        },function () {});
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