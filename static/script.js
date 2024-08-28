 // Function to change the counter value
 function changeCounter(action) {
    // Send a request to the server to update the counter
    fetch('/' + action)
        .then(response => response.text())
        .then(value => {
            // Update the displayed value
            document.getElementById('counter').textContent = value;
        });
}

// Initialize counter value when the page loads
fetch('/value')
    .then(response => response.text())
    .then(value => {
        // Set the initial counter value
        document.getElementById('counter').textContent = value;
    });

// Reset counter when page loads
fetch('/reset')
.then(response => response.text())
.then(value => {
    document.getElementById('counter').textContent = value;
});
