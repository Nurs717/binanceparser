let subscribers = {};
const websocket = new WebSocket("ws://localhost:8080/ws");

websocket.onopen = function(event) {
    console.log("Successfully connected to websocket server");
};

websocket.onerror = function(error) {
    console.log("Error connecting to websocket server");
    console.log(error);
};

websocket.onmessage = function(event) {
    subscribers = JSON.parse(event.data);
    console.log(subscribers);
    document.getElementById('asks').textContent = subscribers
};