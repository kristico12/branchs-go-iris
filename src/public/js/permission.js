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
    }).catch(error => {
        if (error.response.data.message) {
            document.querySelector("#errorList").innerHTML = error.response.data.message;
        } else {
            document.querySelector("#errorList").innerHTML = "A ocurrido un error Intente mas tarde";
        }
    })
}
// load page
window.addEventListener('load',function () {
    DeleteRows();
    Call();
});