//---------------------------------- Global Data------------------------------------|
let branchOffice;
let page = 1;
let orderAscDesc = "ASC";
let titleOrder = "id";

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
        const table = document.querySelector("#tableBranchOffice");
        const thead = table.querySelectorAll("thead td");
        const tbody = table.querySelector("tbody")
        data.Data.forEach((value, i) => {
            const row = tbody.insertRow(i);
            const keyDate = Object.keys(value)
            for (let j = 0; j < thead.length - 1; j++) {
                row.insertCell(j).innerHTML = value[keyDate[j]];
            }
        })
    }).catch(error => {
        const showError = document.querySelector("#errorBranchOffice")
        showError.innerHTML = error;
    })
}

// load page
window.onload = function () {
    branchOffice = location.pathname.split("/").filter(value => value.length > 0)[0];
    Call();
}