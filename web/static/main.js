const websocket = new WebSocket("ws://localhost:8080/ws");
const askList = document.getElementById('ask');
const bidList = document.getElementById('bid');

websocket.onopen = function(event) {
    console.log("Successfully connected to websocket server");
};

websocket.onerror = function(error) {
    console.log("Error connecting to websocket server");
    console.log(error);
};

websocket.onmessage = function(event) {
    orders = JSON.parse(event.data);
    askList.innerHTML = ''
    bidList.innerHTML = ''
    orders.Asks.map(ask => {
        let line = document.createElement('p')
        line.innerText = "price: " + ask.join(' order: ')
        askList.appendChild(line)
    })
    orders.Bids.map(bid => {
        let line = document.createElement('p')
        line.innerText = "price: " + bid.join(' order: ')
        bidList.appendChild(line)
    })
    document.getElementById('asks').innerHTML = orders.SumAsks
    document.getElementById('bids').innerHTML = orders.SumBids
};