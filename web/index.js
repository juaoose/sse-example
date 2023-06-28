const dataContainer = document.getElementById("data-container");
let eventSource = null;
function createEventSource() {
  if (eventSource !== null) {
    return;
  }

  eventSource = new EventSource(
    "http://docker.inferable.farm:8080/juaoose/sse-example:main"
  );

  eventSource.addEventListener("message", function (event) {
    const data = event.data;
    const dataElement = document.createElement("p");
    dataElement.textContent = data;
    dataContainer.appendChild(dataElement);
  });

  eventSource.onerror = function () {
    console.error("Error occurred while streaming data.");
    closeEventSource();
  };

  eventSource.addEventListener("close", function () {
    console.log("SSE connection closed.");
    closeEventSource();
  });
}

function closeEventSource() {
  if (eventSource !== null) {
    eventSource.close();
    eventSource = null;
  }
}

createEventSource();
