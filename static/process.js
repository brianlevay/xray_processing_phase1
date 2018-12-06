/* global getSelected */

let settings = {
    Height: 1550,
    Width: 1032,
    Bits: 14,
    CoreType: 'WR',
    CoreDiameter: 7.2,
    SrcHeight: 65.0,
    CoreHeight: 0.5,
    Motion: 12.5,
    AxisMethod: 'autoDetect',
    AxisAngle: 0.0,
    AxisOffset: 0.0,
    Ilow: 0,
    Ipeak: (2**14-1)/2,
    Ihigh: (2**14-1),
    FolderName: 'processed',
    FileAppend: '_processed'
};


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
   settings.SrcHeight = parseFloat(document.getElementById('srcHeight').value);
   settings.CoreHeight = parseFloat(document.getElementById('coreHeight').value);
   settings.Motion = parseFloat(document.getElementById('motion').value);
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