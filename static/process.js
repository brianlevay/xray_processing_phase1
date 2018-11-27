/* global getSelected */

let settings = {
    Low: 0,
    Mid: (2**14-1)/2,
    High: (2**14-1),
    SrcHeight: 65.0,
    CoreHeight: 0.5,
    CoreDiameter: 7.2,
    Motion: 12.5,
    CoreType: 'WR',
    Contrast: 'skewScale',
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
   settings.Low = parseFloat(document.getElementById('leftBounds').value);
   settings.Mid = parseFloat(document.getElementById('center').value);
   settings.High = parseFloat(document.getElementById('rightBounds').value);
   settings.SrcHeight = parseFloat(document.getElementById('srcHeight').value);
   settings.CoreHeight = parseFloat(document.getElementById('coreHeight').value);
   settings.CoreDiameter = parseFloat(document.getElementById('coreDiameter').value);
   settings.Motion = parseFloat(document.getElementById('motion').value);
   if (document.getElementById('halfRound').checked) {
       settings.CoreType = 'HR';
   } else {
       settings.CoreType = 'WR';
   }
   if (document.getElementById('clipScale').checked) {
       settings.Contrast = 'clipScale';
   } else {
       settings.Contrast = 'skewScale';
   }
   settings.FolderName = document.getElementById('folderName').value;
   settings.FileAppend = document.getElementById('fileAppend').value;
   return;
}