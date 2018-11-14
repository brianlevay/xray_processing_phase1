
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
}


function updateFileList(xhttp) {
    let response = JSON.parse(xhttp.response);
    let rootName = response['Root'];
    let dirNames = response['Dirs'];
    let fileNames = response['Files'];
    
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
    let text = document.createTextNode(rootName); 
    container.className = "root";
    container.appendChild(text);
    container.onclick = function() {
        fileExplorerAPI("..");
    };
    section.appendChild(container);
    return;
}

function addDirElements(sectionName, dirNames) {
    let section = document.getElementById(sectionName);
    for (var i = 0; i < dirNames.length; i++) {
        let container = document.createElement('div');
        let text = document.createTextNode(dirNames[i]); 
        container.className = "dir";
        container.appendChild(text);
        container.onclick = dirClickHandler(dirNames[i]);
        section.appendChild(container);
    }
    return;
}

function addFileElements(sectionName, fileNames) {
    let section = document.getElementById(sectionName);
    for (var i = 0; i < fileNames.length; i++) {
        let container = document.createElement('div');
        let text = document.createTextNode(fileNames[i]); 
        container.className = "file";
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


//// Initial calls on page load ////

document.addEventListener('DOMContentLoaded', function(){ 
    fileExplorerAPI(".");
}, false);