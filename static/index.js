//// API Calls ////

function fileExplorerAPI(newRoot) {
    let xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            updateFileList(this);
        }
        return;
    };
    xhttp.open('POST', '/filesystem', true);
    xhttp.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
    xhttp.send('Root=' + newRoot);
    return;
}

function histogramAPI() {
    let xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            updateHistogram(this);
            toggleButton('getHistogram');
            statusUpdate('histogramIndicator', false);
        }
        return;
    };
    toggleButton('getHistogram');
    statusUpdate('histogramIndicator', true);
    xhttp.open('POST', '/histogram', true);
    xhttp.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
    let selected = getSelected();
    xhttp.send('Selected=' + JSON.stringify(selected));
    return;
}

function processAPI() {
    let xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            alert("Finished Processing!");
            toggleButton('processFiles');
            statusUpdate('processIndicator', false);
        }
        return;
    };
    toggleButton('processFiles');
    statusUpdate('processIndicator', true);
    xhttp.open("POST", "/processing", true);
    xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    let selected = getSelected();
    let settings = getSettings();
    xhttp.send("Selected=" + JSON.stringify(selected) + "&Settings=" + JSON.stringify(settings));
    return;
}

//// File Explorer Functions ////

function updateFileList(xhttp) {
    let response = JSON.parse(xhttp.response);
    let selectAll = document.getElementById('selectAll');
    let rootName = response['Root'];
    let dirNames = response['Dirs'];
    let fileNames = response['Files'];
    selectAll.checked = false;
    removeClassElements('root');
    removeClassElements('dir');
    removeClassElements('file');
    addRootElement('fileList', rootName);
    addDirElements('fileList', dirNames);
    addFileElements('fileList', fileNames);
    return;
}

let dirClickHandler = function(arg) {
    return function() {
        fileExplorerAPI(arg);
    };
};

function removeClassElements(className) {
    let classEls = document.getElementsByClassName(className);
    while(classEls[0]) {
        classEls[0].parentNode.removeChild(classEls[0]);
    }
    return;
}

function addRootElement(sectionName, rootName) {
    let section = document.getElementById(sectionName);
    let container = document.createElement('div');
    let symb = document.createElement('span');
    let text = document.createTextNode(rootName); 
    container.className = 'root';
    symb.innerHTML = '&#8617';
    container.appendChild(symb);
    container.appendChild(text);
    container.onclick = function() {
        fileExplorerAPI('..');
    };
    section.appendChild(container);
    return;
}

function addDirElements(sectionName, dirNames) {
    let section = document.getElementById(sectionName);
    let container, symb, text;
    for (let i = 0; i < dirNames.length; i++) {
        container = document.createElement('div');
        symb = document.createElement('span');
        text = document.createTextNode(dirNames[i]); 
        container.className = 'dir';
        symb.innerHTML = '&#128193;';
        container.appendChild(symb);
        container.appendChild(text);
        container.onclick = dirClickHandler(dirNames[i]);
        section.appendChild(container);
    }
    return;
}

function addFileElements(sectionName, fileNames) {
    let section = document.getElementById(sectionName);
    let container, check, symb, text;
    for (let i = 0; i < fileNames.length; i++) {
        container = document.createElement('div');
        check = document.createElement('input');
        symb = document.createElement('span');
        text = document.createTextNode(fileNames[i]); 
        container.className = 'file';
        check.setAttribute('type', 'checkbox');
        check.setAttribute('class', 'fileCheckbox');
        check.setAttribute('value', fileNames[i]);
        symb.innerHTML = '&#128462;';
        container.appendChild(check);
        container.appendChild(symb);
        container.appendChild(text);
        section.appendChild(container);
    }
    return;
}

function toggleAll(source) {
    let checkboxes = document.getElementsByClassName('fileCheckbox');
    for (let i = 0; i < checkboxes.length; i++) {
        checkboxes[i].checked = source.checked;
    }
    return;
}

function getSelected() {
    let checkboxes = document.getElementsByClassName('fileCheckbox');
    let selected = [];
    for (let i = 0; i < checkboxes.length; i++) {
        if (checkboxes[i].checked == true) {
            selected.push(checkboxes[i].value);
        }
    }
    return selected;
}

//// Histogram Functions ////

function updateHistogram(xhttp) {
    let histogramImg = document.getElementById('histogramImg');
    let IlowFrac = document.getElementById('IlowFrac');
    let IpeakFrac = document.getElementById('IpeakFrac');
    let IhighFrac = document.getElementById('IhighFrac');
    histogramImg.onload = function() {
        IlowFrac.value = 0.0;
        IpeakFrac.value = 0.5;
        IhighFrac.value = 1.0;
        IlowFrac.style.width = histogramImg.width + 'px';
        IpeakFrac.style.width = histogramImg.width + 'px';
        IhighFrac.style.width = histogramImg.width + 'px';
        histogramImg.style.border = "1px solid black";
    };
    histogramImg.src = 'data:image/png;base64,' + xhttp.response;
    return;
}

function setBoundsListeners() {
    let step = 0.01;
    let IlowFrac = document.getElementById('IlowFrac');
    let IpeakFrac = document.getElementById('IpeakFrac');
    let IhighFrac = document.getElementById('IhighFrac');
    IlowFrac.addEventListener('input', function() {
        if (+IlowFrac.value > +IpeakFrac.value) {
            IlowFrac.value = +IpeakFrac.value - step;
        }
    }, false);
    IpeakFrac.addEventListener('change', function() {
        if (+IpeakFrac.value < +IlowFrac.value) {
            IpeakFrac.value = +IlowFrac.value + step;
        } else if (+IpeakFrac.value > +IhighFrac.value) {
            IpeakFrac.value = +IhighFrac.value - step;
        }
    }, false);
    IhighFrac.addEventListener('change', function() {
        if (+IhighFrac.value < +IpeakFrac.value) {
            IhighFrac.value = +IpeakFrac.value + step;
        }
    }, false);
    return;
}

//// Processing Functions ////

function getSettings() {
    var settings = {};
    if (document.getElementById('halfRound').checked) {
       settings["CoreType"] = "HR";
   } else {
       settings["CoreType"] = "WR";
   }
   settings["CoreDiameter"] = parseFloat(document.getElementById('CoreDiameter').value);
   if (document.getElementById('autoDetect').checked) {
       settings["AxisMethod"] = "autoDetect";
   } else {
       settings["AxisMethod"] = "setAxis";
   }
   settings["AxisAngle"] = parseFloat(document.getElementById('AxisAngle').value);
   settings["AxisOffset"] = parseFloat(document.getElementById('AxisOffset').value);
   settings["IlowFrac"] = parseFloat(document.getElementById('IlowFrac').value);
   settings["IpeakFrac"] = parseFloat(document.getElementById('IpeakFrac').value);
   settings["IhighFrac"] = parseFloat(document.getElementById('IhighFrac').value);
   settings["FolderName"] = document.getElementById('FolderName').value;
   settings["FileAppend"] = document.getElementById('FileAppend').value;
   return settings;
}

//// Button locking and process indicators ////

 function toggleButton(id) {
     var button = document.getElementById(id);
     if (button.disabled == false) {
         button.disabled = true;
     } else {
         button.disabled = false;
     }
     return;
 }
 
function statusUpdate(id, show) {
    var span = document.getElementById(id);
    if (show == true) {
        span.innerHTML = "Processing...";
    } else {
        span.innerHTML = "";
    }
    return;
}

//// Initial calls on page load ////

setBoundsListeners();

document.addEventListener('DOMContentLoaded', function(){ 
    fileExplorerAPI('.');
    let selectAll = document.getElementById('selectAll');
    selectAll.onclick = function() {
        toggleAll(this);
    };
}, false);

