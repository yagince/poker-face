var websocketServices = angular.module('websocketServices', []);

websocketServices.factory('Poker', ['$location', function($location) {
    var ws;
    var entryUrl = (function() {
        var protocol = 'ws';
        if ($location.protocol() == 'https') {
            protocol = 'wss';
        }
        var port = '';
        if ($location.port() != 80 && $location.port() != 443) {
            port = ':' + $location.port();
        }
        return protocol + '://' + $location.host() + port + '/entry';
    })();

    return {
        listen: function(onMessage) {
            ws = new WebSocket(entryUrl);
            ws.onmessage = onMessage;
        },
        send: function(message) {
            ws.send(JSON.stringify(message));
        }
    };
}]);

