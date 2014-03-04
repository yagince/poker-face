var websocketServices = angular.module('websocketServices', []);

websocketServices.factory('Poker', function() {
    var ws;
    return {
        listen: function(onMessage) {
            ws = new WebSocket("ws://localhost:8080/entry");
            ws.onmessage = onMessage;
        },
        send: function(message) {
            ws.send(JSON.stringify(message));
        }
    };
});

