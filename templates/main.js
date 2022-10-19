
table_of_departments = document.getElementById("table_of_departments");

const url = "http://127.0.0.1:8084";

// Example POST method implementation:
async function postData(url = '', data = {}) {
    // Default options are marked with *
    const response = await fetch(url, {
        method: 'POST', // *GET, POST, PUT, DELETE, etc.
        mode: 'cors', // no-cors, *cors, same-origin
        cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
        credentials: 'same-origin', // include, *same-origin, omit
        headers: {
            'Content-Type': 'application/json'
            // 'Content-Type': 'application/x-www-form-urlencoded',
        },
        redirect: 'follow', // manual, *follow, error
        referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
        body: JSON.stringify(data) // body data type must match "Content-Type" header
    });
    return response.json(); // parses JSON response into native JavaScript objects
}

async function putData(url = '', data = {}) {
    // Default options are marked with *
    const response = await fetch(url, {
        method: 'PUT', // *GET, POST, PUT, DELETE, etc.
        mode: 'cors', // no-cors, *cors, same-origin
        cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
        credentials: 'same-origin', // include, *same-origin, omit
        headers: {
            'Content-Type': 'application/json'
            // 'Content-Type': 'application/x-www-form-urlencoded',
        },
        redirect: 'follow', // manual, *follow, error
        referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
        body: JSON.stringify(data) // body data type must match "Content-Type" header
    });
    return response.json(); // parses JSON response into native JavaScript objects
}

async function deleteData(url = '', data = {}) {
    // Default options are marked with *
    const response = await fetch(url, {
        method: 'DELETE', // *GET, POST, PUT, DELETE, etc.
        mode: 'cors', // no-cors, *cors, same-origin
        cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
        credentials: 'same-origin', // include, *same-origin, omit
        headers: {
            'Content-Type': 'application/json'
            // 'Content-Type': 'application/x-www-form-urlencoded',
        },
        redirect: 'follow', // manual, *follow, error
        referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
        body: JSON.stringify(data) // body data type must match "Content-Type" header
    });
    return response.json(); // parses JSON response into native JavaScript objects
}

function initialLoad() {
    loadDepartments();

}

function loadDepartments() {
    table_of_departments = document.getElementById("table_of_departments");

    fetch(url+'/api/v1/crud/departments')
        .then((response) => response.json())
        .then((data) => {
            console.log(data);
            departments = data.departments;
            console.log(departments);
            htmlstring = "";
            for (var i = 0; i < departments.length; i++) {
                console.log(departments[i]);
                htmlstring += "<tr><th scope>" + (i + 1) + "</th><td>" + departments[i].department_name + "</td><td>" + 
                "<div onclick='deleteDepartment("+departments[i].department_id+")'>🗑️</div>"+
                "<div onclick='getOneDepartment("+departments[i].department_id+")'>🖉</div>"+
                "</td></tr>";
            }
            console.log(htmlstring);
            console.log(table_of_departments);
            table_of_departments.innerHTML = htmlstring;

        })
}

async function createDepartment() {
    department_name = document.getElementById("department_name");
    console.log(department_name.value);

    await postData(url+'/api/v1/crud/departments', { department_name: department_name.value })
        .then((data) => {
            console.log(data); // JSON data parsed by `data.json()` call
        });

    department_name.value = "";
    loadDepartments()
}

async function getOneDepartment(id = 0) {
    department_name = document.getElementById("department_name");
    boton_bipolar = document.getElementById("boton_bipolar");
    department_name.value = "";
    await fetch(url+'/api/v1/crud/departments/' + id)
        .then((response) => response.json())
        .then((data) => {
            console.log(data);
            department_name.value = data.department.department_name;
            boton_bipolar.innerHTML = getUpdateButtonString(data.department.department_id);

        })
}

async function deleteDepartment(id = 0) {
    await deleteData(url+'/api/v1/crud/departments/' + id)

    loadDepartments()
}

async function updateDepartment(id = 0) {
    department_name = document.getElementById("department_name");
    boton_bipolar = document.getElementById("boton_bipolar");
    console.log(department_name.value);

    await putData(url+'/api/v1/crud/departments/' + id, { department_name: department_name.value })

    department_name.value = "";
    loadDepartments()
    boton_bipolar.innerHTML = getCreateButtonString();
}

function getCreateButtonString(){
    return '<button type="button" class="btn btn-primary" onclick="createDepartment()">'+
 ' <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" fill="currentColor" class="bi bi-plus-circle-fill" viewBox="0 0 16 16">'+
   ' <path d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0zM8.5 4.5a.5.5 0 0 0-1 0v3h-3a.5.5 0 0 0 0 1h3v3a.5.5 0 0 0 1 0v-3h3a.5.5 0 0 0 0-1h-3v-3z"/>'+
  '/svg>'+
'/button>';
}

function getUpdateButtonString(id = 0){
    return '<button type="button" class="btn btn-success" onclick="updateDepartment('+id+')">'+
    ' <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" fill="currentColor" class="bi bi-save" viewBox="0 0 16 16">'+
    ' <path d="M2 1a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H9.5a1 1 0 0 0-1 1v7.293l2.646-2.647a.5.5 0 0 1 .708.708l-3.5 3.5a.5.5 0 0 1-.708 0l-3.5-3.5a.5.5 0 1 1 .708-.708L7.5 9.293V2a2 2 0 0 1 2-2H14a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h2.5a.5.5 0 0 1 0 1H2z"/>'+
   '</svg>'+
'</button>';
}