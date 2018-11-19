/* global getSelected */

var hist = {
    Bits: 14,
    Nbins: 400,
    Width: 800,
    Height: 600
};

let leftBounds = document.getElementById('leftBounds');
let center = document.getElementById('center');
let rightBounds = document.getElementById('rightBounds');
let leftBoundsVal = document.getElementById('leftBoundsVal');
let centerVal = document.getElementById('centerVal');
let rightBoundsVal = document.getElementById('rightBoundsVal');


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
    let variableStr = "";
    for (var key in hist) {
        if (hist.hasOwnProperty(key)) {
            variableStr += "&" + key + "=" + hist[key];
        }
    }
    xhttp.send("Selected=" + JSON.stringify(selected) + variableStr);
    return;
}


function updateHistogram(xhttp) {
    let histogramImg = document.getElementById('histogramImg');
    histogramImg.src = "data:image/png;base64," + xhttp.response;
    leftBounds.max = (2**hist['Bits']) - 1;
    leftBounds.value = 0;
    leftBounds.style.width = hist["Width"] + "px";
    center.max = leftBounds.max;
    center.value = 0.5*center.max;
    center.style.width = hist["Width"] + "px";
    rightBounds.max = leftBounds.max;
    rightBounds.value = rightBounds.max;
    rightBounds.style.width = hist["Width"] + "px";
    updateSliderText();
    return;
}


function updateSliderText() {
    leftBoundsVal.innerHTML = leftBounds.value;
    centerVal.innerHTML = center.value;
    rightBoundsVal.innerHTML = rightBounds.value;
    return;
}


leftBounds.addEventListener("input", function() {
    if (+leftBounds.value > +center.value) {
        leftBounds.value = +center.value - 1;
    }
    updateSliderText();
}, false);


center.addEventListener("change", function() {
    if (+center.value < +leftBounds.value) {
        center.value = +leftBounds.value + 1;
    } else if (+center.value > +rightBounds.value) {
        center.value = +rightBounds.value - 1;
    }
    updateSliderText();
}, false);


rightBounds.addEventListener("change", function() {
    if (+rightBounds.value < +center.value) {
        rightBounds.value = +center.value + 1;
    }
    updateSliderText();
}, false);

