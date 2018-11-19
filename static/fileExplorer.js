
function fileExplorerAPI(newRoot) {
    let xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            updateFileList(this);
        }
    };
    xhttp.open("POST", "/filesystem", true);
    xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhttp.send("Root=" + newRoot);
    return;
}


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
    addRootElement('fileExplorer', rootName);
    addDirElements('fileExplorer', dirNames);
    addFileElements('fileExplorer', fileNames);
    return;
}


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
    container.className = "root";
    symb.innerHTML = '&#8617';
    container.appendChild(symb);
    container.appendChild(text);
    container.onclick = function() {
        fileExplorerAPI("..");
    };
    section.appendChild(container);
    return;
}


function addDirElements(sectionName, dirNames) {
    let section = document.getElementById(sectionName);
    for (let i = 0; i < dirNames.length; i++) {
        let container = document.createElement('div');
        let symb = document.createElement('span');
        let text = document.createTextNode(dirNames[i]); 
        container.className = "dir";
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
    for (let i = 0; i < fileNames.length; i++) {
        let container = document.createElement('div');
        let check = document.createElement('input');
        let symb = document.createElement('span');
        let text = document.createTextNode(fileNames[i]); 
        container.className = "file";
        check.setAttribute("type", "checkbox");
        check.setAttribute("class", "fileCheckbox");
        check.setAttribute("value", fileNames[i]);
        symb.innerHTML = '&#128462;';
        container.appendChild(check);
        container.appendChild(symb);
        container.appendChild(text);
        section.appendChild(container);
    }
    return;
}


let dirClickHandler = function(arg) {
    return function() {
        fileExplorerAPI(arg);
    };
};


function toggleAll(source) {
    let checkboxes = document.querySelectorAll('input[type="checkbox"]');
    for (let i = 0; i < checkboxes.length; i++) {
        if (checkboxes[i] != source) {
            checkboxes[i].checked = source.checked;
        }
    }
    return;
}


function getSelected() {
    let checkboxes = document.querySelectorAll('input[type="checkbox"]');
    let selectAll = document.getElementById('selectAll');
    var selected = [];
    for (let i = 0; i < checkboxes.length; i++) {
        if ((checkboxes[i].checked == true) && (checkboxes[i] != selectAll)) {
            selected.push(checkboxes[i].value);
        }
    }
    return selected;
}


//// Initial calls on page load ////

document.addEventListener('DOMContentLoaded', function(){ 
    fileExplorerAPI(".");
    let selectAll = document.getElementById('selectAll');
    selectAll.onclick = function() {
        toggleAll(this);
    };
}, false);