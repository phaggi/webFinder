document.getElementById('trigger-script').addEventListener('click', function() {
    fetch('/trigger_script', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ script_name: 'example_script' })
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('results').innerText = 'Task ID: ' + data.task_id;
    });
});