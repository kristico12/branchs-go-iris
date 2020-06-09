//--------------------- CONSTANT ------------------------------|
class CitysProvince {
    constructor(region, c_digo_dane_del_departamento, departamento, c_digo_dane_del_municipio, municipio) {
        this.Region = region
        this.CodiDepartment = c_digo_dane_del_departamento
        this.DepartmentName = departamento
        this.CodiCity = c_digo_dane_del_municipio
        this.City = municipio
    }
}

let AllCities = [];
let filterProvince = [];
let selectProvince;
let selectCity;
const dialogTimeEntranceHour = new mdDateTimePicker.default({
    type: 'time',
    trigger: document.querySelector("#inputEntranceHour")

});
const dialogTimeExitHour = new mdDateTimePicker.default({
    type: 'time',
    trigger: document.querySelector("#inputExitHour")

});
function capitalize(word) {
    return word[0].toUpperCase() + word.slice(1);
}
//------------------- functions -----------------------------------|
function SetValuesAllCities(values) {
    for (const key of values) {
        AllCities.push(new CitysProvince(key.region, key.c_digo_dane_del_departamento, key.departamento.normalize("NFD").replace(/[\u0300-\u036f]/g, ""),
            key.c_digo_dane_del_municipio, key.municipio.normalize("NFD").replace(/[\u0300-\u036f]/g, "")))
    }
}
function SetFilterProvince() {
    filterProvince = Array.from(new Set(AllCities.map(idProv => idProv.CodiDepartment)))
        .map(id => new CitysProvince("", id, AllCities.find(val => id === val.CodiDepartment).DepartmentName, "", ""))
}
function SetDataSelectProvince() {
    let insertSelect = document.querySelector("#select-province");
    for (const province of filterProvince) {
        const option = document.createElement("option");
        option.text = province.DepartmentName;
        option.value = province.CodiDepartment;
        insertSelect.add(option);
    }
}
function ClearDataCities() {
    const removeOption = document.querySelectorAll("#select-city option");
    removeOption.forEach(o => o.remove());
}
function SetDataSelectCity(id) {
    const insertOptionsSelectCitys = AllCities.filter(key => key.CodiDepartment === id).map(key => {
        return { text: key.City, id: key.CodiCity }
    });
    insertOptionsSelectCitys.splice(0,0, { text: "Seleccionar...", id: "" })
    const insertSelect = document.querySelector("#select-city");
    for (const selectCityElement of insertOptionsSelectCitys) {
        const option = document.createElement("option");
        option.text = selectCityElement.text;
        option.value = selectCityElement.id;
        insertSelect.add(option);
    }
}
window.addEventListener('load',function () {
    // assing year
    const year = document.querySelector("#year");
    year.innerHTML = new Date().getFullYear().toString();
    // init optionts
    selectProvince = new SlimSelect({
        select: '#select-province'
    });
    // init select
    selectCity = new SlimSelect({
        select: '#select-city'
    });
    // axios
    axios.get('https://www.datos.gov.co/resource/xdk5-pm3f.json')
        .then(response => {
            SetValuesAllCities(response.data);
            SetFilterProvince();
            SetDataSelectProvince();
        })
});
function HandleOptionSelectProvince() {
    const selected = selectProvince.selected();
    if (selected !== "") {
        ClearDataCities();
        SetDataSelectCity(selected)
    }
}
document.querySelector("#dialogOpenEntrance").addEventListener('click', function () {
    dialogTimeEntranceHour.toggle();
});
document.querySelector("#inputEntranceHour").addEventListener('onOk', function () {
    this.value = dialogTimeEntranceHour.time.set({second:0}).format('HH:mm:ss');
});
document.querySelector("#dialogOpenExit").addEventListener('click', function () {
    dialogTimeExitHour.toggle();
});
document.querySelector("#inputExitHour").addEventListener('onOk', function () {
    this.value = dialogTimeExitHour.time.set({second:0}).format('HH:mm:ss');
});
document.querySelector("#saveBranchOffice").addEventListener('click', function () {
    const newBranchOffice = {
        city: selectCity.selected(),
        province: selectProvince.selected(),
        address: document.querySelector("#inputAddress").value,
        checkInTime: document.querySelector("#inputEntranceHour").value,
        exitTime: document.querySelector("#inputExitHour").value
    }
    // remove info errors
    Object.keys(newBranchOffice).forEach(value => {
        document.querySelector(`#error${capitalize(value)}`).textContent = null;
    });
    // show Loadin
    document.querySelector("#loading").classList.remove("uk-hidden");
    axios({
        method: "POST",
        url: '/api/branch_office/',
        headers: {
            'Content-Type': 'application/json',
        },
        data: newBranchOffice
    })
        .then(() => location.reload())
        .catch(error => {
            const infoError = error.response;
            if (infoError.status === 400) {
                const errorData = infoError.data;
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
            // quit loading
            document.querySelector("#loading").classList.add("uk-hidden");
        })
})