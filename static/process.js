/* global getSelected */

let settings = {
    Bits: 14,
    Low: 0,
    Mid: (2**14-1)/2,
    High: (2**14-1),
    Method: 'compensation',
    SrcHeight: 65.0,
    CoreHeight: 0.5,
    CoreDiameter: 7.2,
    CoreType: 'WR',
    Contrast: 'skewScale',
    ROISize: 12.5,
    IncludeScale: true,
    IncludeROI: true,
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
    xhttp.open('POST', '/filesystem', true);
    xhttp.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
    let selected = getSelected();
    updateSettings();
    xhttp.send('Selected=' + JSON.stringify(selected) + '&Settings=' + JSON.stringify(settings));
    return;
}


function updateSettings() {
   if (document.getElementById('radio16b').checked) {
       settings.Bits = 16;
   } else {
       settings.Bits = 14;
   }
   settings.Low = document.getElementById('leftBounds').value;
   settings.Mid = document.getElementById('center').value;
   settings.High = document.getElementById('rightBounds').value;
   if (document.getElementById('radioConvert16b').checked) {
       settings.Method = 'convert16b';
   } else if (document.getElementById('radioMurhot').checked) {
       settings.Method = 'murhot';
   } else {
       settings.Method = 'compensation';
   }
   settings.SrcHeight = document.getElementById('srcHeight').value;
   settings.CoreHeight = document.getElementById('coreHeight').value;
   settings.CoreDiameter = document.getElementById('coreDiameter').value;
   if (document.getElementById('halfRound').checked) {
       settings.CoreType = 'HR';
   } else {
       settings.CoreType = 'WR';
   }
   if (document.getElementById('defaultScale').checked) {
       settings.Contrast = 'defaultScale';
   } else if (document.getElementById('clipScale').checked) {
       settings.Contrast = 'clipScale';
   } else {
       settings.Contrast = 'skewScale';
   }
   settings.ROISize = document.getElementById('roiSize').value;
   if (document.getElementById('includeScale').checked) {
       settings.IncludeScale = true;
   } else {
       settings.IncludeScale = false;
   }
   if (document.getElementById('includeROI').checked) {
       settings.IncludeROI = true;
   } else {
       settings.IncludeROI = false;
   }
   settings.FolderName = document.getElementById('folderName').value;
   settings.FileAppend = document.getElementById('fileAppend').value;
   return;
}