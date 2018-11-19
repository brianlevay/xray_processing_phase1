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
    let selected = getSelected();
    xhttp.send("Selected=" + JSON.stringify(selected) + "&Bits=14&Nbins=1000&Width=800&Height=600");
    return;
}


function updateHistogram(xhttp) {
    let hist_img = document.getElementById('histogram_img');
    hist_img.src = "data:image/png;base64," + xhttp.response;
    return;
}