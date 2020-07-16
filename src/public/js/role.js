//--------------------------------- load data table ------------------------------------------------|
let page = 1;
let orderAscDesc = "ASC";
let titleOrder = "id";
let filter;
const table = document.querySelector("#tableRole");
const thead = table.querySelectorAll("thead td");
const tbody = table.querySelector("tbody");
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
    document.querySelector("#errorList").innerHTML = "";
}
function Call() {
    // axios
    axios.get('/api/role/', {
        params: {
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
                if (typeof value[keyDate[j]] === "object") {
                    row.insertCell(j).innerHTML = value[keyDate[j]].String
                } else {
                    row.insertCell(j).innerHTML = value[keyDate[j]];
                }
            }
            row.insertCell(thead.length - 1).innerHTML = '<div class="uk-flex uk-flex-around">' +
                '<i uk-toggle="target: #modal-edit-role" class="icon" uk-icon="icon: pencil; ratio: 1.3" onclick="OpenEdit(\''+i+'\')"></i>' +
                '<i class="icon" uk-icon="icon: minus-circle; ratio: 1.3"></i>' +
                '</div>';
        });
        CreatePaginate(data);
    }).catch(error => {
        if (error.response.data.message) {
            document.querySelector("#errorList").innerHTML = error.response.data.message;
        } else {
            document.querySelector("#errorList").innerHTML = "A ocurrido un error Intente mas tarde";
        }
    })
}
// order by
table.querySelector("thead tr").addEventListener('click', function (e) {
    if (e.target.id !== "") {
        page = 1;
        titleOrder = e.target.id;
        orderAscDesc = orderAscDesc === "ASC" ? "DESC" : "ASC";
        DeleteRows();
        ClearError();
        Call();
    }
});
//------------------------------------------- Save Role --------------------------------------------------------------|
document.querySelector("#saveRole").addEventListener('click', function () {
    const role = {
        name: document.querySelector("#inputName").value,
        description: document.querySelector("#inputDescription").value
    }
    // remove info errors
    Object.keys(role).forEach(value => {
        document.querySelector(`#error${capitalize(value)}`).textContent = null;
    });
    // show Loadin
    document.querySelector("#loading").classList.remove("uk-hidden");
    axios({
        method: "POST",
        url: '/api/role/',
        headers: {
            'Content-Type': 'application/json',
        },
        data: role
    })
        .then(() => location.reload())
        .catch(error => {
            const infoErrorPermission = error.response;
            if (infoErrorPermission.status === 400) {
                const errorData = infoErrorPermission.data;
                if (Object.keys(errorData).includes("errors")) {
                    for (const customError of errorData.errors) {
                        const setError = document.querySelector(`#error${capitalize(customError.Key)}`);
                        setError.textContent = customError.Value;
                    }
                }
                if (Object.keys(errorData).includes("message")) {
                    const setError = document.querySelector(`#message`);
                    setError.textContent = errorData.message;
                }
            }
            document.querySelector("#loading").classList.add("uk-hidden");
        })
})
//--------------------------------------------- update role ---------------------------------------------------------------|
function GetDataTable(i) {
    const row = tbody.rows.item(i);
    let role = {};
    for (const j in row.cells) {
        const cell = row.cells[j];
        if (cell.nodeName === "TD") {
            switch (j) {
                case "0":
                    role.id = cell.textContent;
                    break;
                case "1":
                    role.name = cell.textContent;
                    break;
                case "2":
                    role.description = cell.textContent;
                    break;
                default:
                    break;
            }
        }
    }
    return role;
}
function OpenEdit(i) {
    let editRole = Object.assign({},GetDataTable(i));
    document.querySelector("#inputEditId").value = editRole.id;
    document.querySelector("#inputEditName").value = editRole.name;
    document.querySelector("#inputEditDescription").value = editRole.description;
}
document.querySelector("#editRole").addEventListener('click', function () {
    const editRole = {
        id: parseInt(document.querySelector("#inputEditId").value),
        name: document.querySelector("#inputEditName").value,
        description: document.querySelector("#inputEditDescription").value,
    }
    // remove error general
    document.querySelector(`#messageEdit`).textContent = "";
    // remove info errors
    Object.keys(editRole).forEach(value => {
        if (value !== 'id') {
            document.querySelector(`#errorEdit${capitalize(value)}`).textContent = null;
        }
    });
    // show Loadin
    document.querySelector("#loadingEdit").classList.remove("uk-hidden");
    axios({
        method: "PUT",
        url: '/api/role/',
        headers: {
            'Content-Type': 'application/json',
        },
        data: editRole
    })
        .then(() => {
            DeleteRows();
            ClearError();
            Call();
            // quit loading
            document.querySelector("#loadingEdit").classList.add("uk-hidden");
            // close modal
            UIkit.modal("#modal-edit-role").hide();
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
})
// load page
window.addEventListener('load',function () {
    DeleteRows();
    Call();
});