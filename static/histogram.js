/* global getSelected */

function histogramAPI() {
    let xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            updateHistogram(this);
        }
    };
    xhttp.open("POST", "/histogram", true);
    xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    var selected = getSelected();
    xhttp.send("Selected=" + JSON.stringify(selected));
    return;
}


function updateHistogram(xhttp) {
    let response = JSON.parse(xhttp.response);
    console.log(response);
    return;
}