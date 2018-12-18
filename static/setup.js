//// API Calls ////

function getSetupAPI() {
    let xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            updateSetupUI(this);
        }
        return;
    };
    xhttp.open('GET', '/setup', true);
    xhttp.send();
    return;
}

function sendSetupAPI() {
    let xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            alert("Setup updated!");
        }
        return;
    };
    xhttp.open('POST', '/setup', true);
    xhttp.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
    let setup = getSetupValues();
    xhttp.send('Setup=' + JSON.stringify(setup));
    return;
}

//// Setup Functions ////

function updateSetupUI(xhttp) {
    let response = JSON.parse(xhttp.response);
    let element;
    for (var key in response) {
        if(response.hasOwnProperty(key)) {
            element = document.getElementById(key);
            element.value = response[key];
        }
    }
    return;
}

function getSetupValues() {
    let setup = {};
    let inputs = document.getElementsByTagName('input');
    for (var i = 0; i < inputs.length; i++) {
        setup[inputs[i].id] = parseFloat(inputs[i].value);
    }
    return setup;
}

//// Initial calls on page load ////

document.addEventListener('DOMContentLoaded', function(){ 
    getSetupAPI();
}, false);

