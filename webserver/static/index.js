const health = async () => {
    let result;
    try {
        const resp = await fetch('/health');
        const json = await resp.json();
        result = json;
    } catch(err) {
        result = err;
    }
    document.getElementById('healthdata').innerHTML = JSON.stringify(result);
}

const deleteLogEntry = async (id) => {
    console.log(`Delete id:${id}`);
    try {
        const resp = await fetch(`/log/${id}`, {method: 'DELETE'});
        if (resp.ok) {
            getLogs();
        }
    } catch(err) {
        alert(err);
    }
}

const updateLogTable = (logEntries) => {
    const logTable = document.getElementById('log_tbody');
    logTable.innerHTML = "";
    let cntr = 0;
    
    logEntries.forEach(logEntry => {
        let row = logTable.insertRow(cntr);
        let idCell = row.insertCell(0);
        let messageCell = row.insertCell(1);
        let createdCell = row.insertCell(2);
        let optionsCell = row.insertCell(3);
        idCell.innerHTML = logEntry.id;
        messageCell.innerHTML = logEntry.message;
        createdCell.innerHTML = logEntry.created;
         const deleteButton = document.createElement('button');
         deleteButton.innerHTML = 'X';
         deleteButton.onclick = function() {deleteLogEntry(logEntry.id);}
         optionsCell.appendChild(deleteButton);
        cntr++;
    });
}

const addLog = async () => {
    const logMessage = document.getElementById('logEntry').value;
    if (!logMessage) {
        alert('Please provide a message');
        return;
    } 
    try {
        const body = {message: logMessage};
        const result = await fetch('/log', {'method':'POST', body: JSON.stringify(body)});
        if (!result.ok) {
            alert(`Result:${result.status}`);
            return;
        }
        document.getElementById('logEntry').value = '';
        getLogs();
    } catch(err) {
        alert(`Error: ${err}`);
    }
}

const getLogs = async () => {
    try {
        const resp = await fetch('/log');
        if (resp.ok) {
            const json = await resp.json();
            updateLogTable(json);
        }
    } catch(err) {
        console.error(err);
    }
}

getLogs();