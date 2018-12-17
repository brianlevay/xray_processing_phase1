/* global getSelected */

let settings = {
    CoreType: 'WR',
    CoreDiameter: 7.2,
    AxisMethod: 'autoDetect',
    AxisAngle: 0.0,
    AxisOffset: 0.0,
    IlowFrac: 0.0,
    IpeakFrac: 0.5,
    IhighFrac: 1.0,
    FolderName: 'processed',
    FileAppend: '_processed'
};

// Processing Functions //

function processAPI() {
    let xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            alert("Finished Processing!");
        }
        return;
    };
    xhttp.open('POST', '/processing', true);
    xhttp.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
    let selected = getSelected();
    updateSettings();
    xhttp.send('Selected=' + JSON.stringify(selected) + '&Settings=' + JSON.stringify(settings));
    return;
}


function updateSettings() {
    if (document.getElementById('halfRound').checked) {
       settings.CoreType = 'HR';
   } else {
       settings.CoreType = 'WR';
   }
   settings.CoreDiameter = parseFloat(document.getElementById('coreDiameter').value);
   if (document.getElementById('autoDetect').checked) {
       settings.AxisMethod = 'autoDetect';
   } else {
       settings.AxisMethod = 'setAxis';
   }
   settings.AxisAngle = parseFloat(document.getElementById('axisAngle').value);
   settings.AxisOffset = parseFloat(document.getElementById('axisOffset').value);
   settings.Ilow = parseFloat(document.getElementById('leftBounds').value);
   settings.Ipeak = parseFloat(document.getElementById('center').value);
   settings.Ihigh = parseFloat(document.getElementById('rightBounds').value);
   settings.FolderName = document.getElementById('folderName').value;
   settings.FileAppend = document.getElementById('fileAppend').value;
   return;
}

// Histogram Functions //

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
    xhttp.send('Selected=' + JSON.stringify(selected));
    return;
}


function updateHistogram(xhttp) {
    let histogramImg = document.getElementById('histogramImg');
    let leftBounds = document.getElementById('leftBounds');
    let center = document.getElementById('center');
    let rightBounds = document.getElementById('rightBounds');
    histogramImg.onload = function() {
        leftBounds.value = 0.0;
        center.value = 0.5;
        rightBounds.value = 1.0;
        leftBounds.style.width = histogramImg.width + 'px';
        center.style.width = histogramImg.width + 'px';
        rightBounds.style.width = histogramImg.width + 'px';
    };
    histogramImg.src = 'data:image/png;base64,' + xhttp.response;
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
    }, false);
    center.addEventListener('change', function() {
        if (+center.value < +leftBounds.value) {
            center.value = +leftBounds.value + 1;
        } else if (+center.value > +rightBounds.value) {
            center.value = +rightBounds.value - 1;
        }
    }, false);
    rightBounds.addEventListener('change', function() {
        if (+rightBounds.value < +center.value) {
            rightBounds.value = +center.value + 1;
        }
    }, false);
    return;
}


//// Initial calls on page load ////

setBoundsListeners();