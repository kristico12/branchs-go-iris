//--------------------------------- load data table ------------------------------------------------|
let page = 1;
let orderAscDesc = "ASC";
let titleOrder = "id";
let filter;
const table = document.querySelector("#tablePermission");
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
    axios.get('/api/permission/', {
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
                row.insertCell(j).innerHTML = value[keyDate[j]];
            }
            row.insertCell(thead.length - 1).innerHTML = '<div class="uk-flex uk-flex-around">' +
                '<i uk-toggle="target: #modal-edit-branch-Office" class="icon" uk-icon="icon: pencil; ratio: 1.3" onclick="OpenEdit(\''+i+'\')"></i>' +
                '<i class="icon" uk-icon="icon: minus-circle; ratio: 1.3" onclick="Delete(\''+i+'\')"></i>' +
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
//------------------------------------------------- Save Permission ------------------------------------------------------|
document.querySelector("#savePermission").addEventListener('click', function () {
    const permission = {
        name: document.querySelector("#inputName").value
    }
    // remove info errors
    Object.keys(permission).forEach(value => {
        if (value !== 'id') {
            document.querySelector(`#error${capitalize(value)}`).textContent = null;
        }
    });
    // show Loadin
    document.querySelector("#loading").classList.remove("uk-hidden");
    axios({
        method: "POST",
        url: '/api/permission/',
        headers: {
            'Content-Type': 'application/json',
        },
        data: permission
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
//---------------------------------------------- Delete Permission ---------------------------------------------------|
function GetDataTable(i) {
    const row = tbody.rows.item(i);
    let permission = {};
    for (const j in row.cells) {
        const cell = row.cells[j];
        if (cell.nodeName === "TD" && cell.textContent.length > 0) {
            switch (j) {
                case "0":
                    permission.id = cell.textContent;
                    break;
                case "1":
                    permission.name = cell.textContent;
                    break;
                default:
                    break;
            }
        }
    }
    return permission;
}
function Delete(i) {
    let deletePermission = Object.assign({},GetDataTable(i));
    deletePermission.id = parseInt(deletePermission.id);
    // clear message error
    ClearError();
    UIkit.modal.confirm("Esta seguro que desea eliminar a "+deletePermission.name)
        .then(function () {
            axios({
                method: "DELETE",
                url: '/api/permission/',
                headers: {
                    'Content-Type': 'application/json',
                },
                data: deletePermission
            })
                .then(() => {
                    page = 1;
                    DeleteRows();
                    ClearError();
                    Call();
                })
                .catch(error => {
                    if (error.response.data.message) {
                        document.querySelector("#errorList").innerHTML = error.response.data.message;
                    } else {
                        document.querySelector("#errorList").innerHTML = "A ocurrido un error Intente mas tarde";
                    }
                })
        },function () {});
}
// load page
window.addEventListener('load',function () {
    DeleteRows();
    Call();
});