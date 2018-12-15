/* global getSelected */

var hist = {
    Width: 800,
    Height: 600,
    R: 66,
    G: 134,
    B: 244
};


function histogramAPI() {
    let xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            updateHistogram(this);
        }
        return;
    };
    xhttp.open('POST', '/histogram', true);
    xhttp.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
    let selected = getSelected();
    xhttp.send('Selected=' + JSON.stringify(selected) + '&Style=' + JSON.stringify(hist));
    return;
}


function updateHistogram(xhttp) {
    let histogramImg = document.getElementById('histogramImg');
    let leftBounds = document.getElementById('leftBounds');
    let center = document.getElementById('center');
    let rightBounds = document.getElementById('rightBounds');
    histogramImg.src = 'data:image/png;base64,' + xhttp.response;
    leftBounds.max = (2**14) - 1;
    leftBounds.value = 0;
    leftBounds.style.width = hist['Width'] + 'px';
    center.max = leftBounds.max;
    center.value = 0.5*center.max;
    center.style.width = hist['Width'] + 'px';
    rightBounds.max = leftBounds.max;
    rightBounds.value = rightBounds.max;
    rightBounds.style.width = hist['Width'] + 'px';
    updateSliderText();
    return;
}


function updateSliderText() {
    let leftBounds = document.getElementById('leftBounds');
    let center = document.getElementById('center');
    let rightBounds = document.getElementById('rightBounds');
    let leftBoundsVal = document.getElementById('leftBoundsVal');
    let centerVal = document.getElementById('centerVal');
    let rightBoundsVal = document.getElementById('rightBoundsVal');
    leftBoundsVal.innerHTML = leftBounds.value;
    centerVal.innerHTML = center.value;
    rightBoundsVal.innerHTML = rightBounds.value;
    return;
}


function setBoundsListeners() {
    let leftBounds = document.getElementById('leftBounds');
    let center = document.getElementById('center');
    let rightBounds = document.getElementById('rightBounds');
    leftBounds.addEventListener('input', function() {
        if (+leftBounds.value > +center.value) {
            leftBounds.value = +center.value - 1;
        }
        updateSliderText();
    }, false);
    center.addEventListener('change', function() {
        if (+center.value < +leftBounds.value) {
            center.value = +leftBounds.value + 1;
        } else if (+center.value > +rightBounds.value) {
            center.value = +rightBounds.value - 1;
        }
        updateSliderText();
    }, false);
    rightBounds.addEventListener('change', function() {
        if (+rightBounds.value < +center.value) {
            rightBounds.value = +center.value + 1;
        }
        updateSliderText();
    }, false);
    return;
}


//// Initial calls on page load ////

setBoundsListeners();