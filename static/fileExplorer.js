
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
    let dirNames = response['Dirs'];
    let fileNames = response['Files'];
    
    let dirEls = document.getElementsByClassName('dir');
    let fileEls = document.getElementsByClassName('file');
    while(fileEls[0]) {
        fileEls[0].parentNode.removeChild(fileEls[0]);
    }
    while(dirEls[0]) {
        dirEls[0].parentNode.removeChild(dirEls[0]);
    }
    
    let fileSection = document.getElementById('fileExplorer');
    for (var i = 0; i < dirNames.length; i++) {
        let dirEl = document.createElement('div');
        let dirName = document.createTextNode(dirNames[i]); 
        dirEl.className = "dir";
        dirEl.appendChild(dirName);
        fileSection.appendChild(dirEl);
    }
    for (var i = 0; i < fileNames.length; i++) {
        let fileEl = document.createElement('div');
        let fileName = document.createTextNode(fileNames[i]); 
        fileEl.className = "file";
        fileEl.appendChild(fileName);
        fileSection.appendChild(fileEl);
    }
}

//// Initial calls on page load ////

document.addEventListener('DOMContentLoaded', function(){ 
    fileExplorerAPI(".");
}, false);