const websocket = new WebSocket("ws://localhost:8080/ws");

websocket.onopen = function(event) {
    console.log("Successfully connected to websocket server");
};

websocket.onerror = function(error) {
    console.log("Error connecting to websocket server");
    console.log(error);
};

websocket.onmessage = function(event) {
    orders = JSON.parse(event.data);
    // console.log(orders);
    document.getElementById('asks').innerHTML = orders.SumAsks
    document.getElementById('bids').innerHTML = orders.SumBids
};