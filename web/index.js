const dataContainer = document.getElementById('data-container');
let eventSource = null; // Variable to hold the EventSource object

function createEventSource() {
    if (eventSource !== null) {
        // EventSource connection already exists
        return;
    }

    // Create an EventSource object to receive server-sent events
    eventSource = new EventSource('http://docker.inferable.farm:8080/juaoose/sse-example:latest');

    // Event listener to handle incoming events
    eventSource.addEventListener('message', function(event) {
        // Append the received data to the data container
        const data = event.data;
        const dataElement = document.createElement('p');
        dataElement.textContent = data;
        dataContainer.appendChild(dataElement);
    });

    // Event listener to handle errors
    eventSource.onerror = function() {
        console.error('Error occurred while streaming data.');
        closeEventSource(); // Close the EventSource connection on error
    };

    // Event listener to handle SSE stream termination
    eventSource.addEventListener('close', function() {
        console.log('SSE connection closed.');
        closeEventSource(); // Close the EventSource connection on stream termination
    });
}

function closeEventSource() {
    if (eventSource !== null) {
        eventSource.close();
        eventSource = null;
    }
}

// Call the createEventSource function to establish the EventSource connection
createEventSource();