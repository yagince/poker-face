var app = angular.module('poker', [])

app.controller('PokerController', ['$scope', function($scope) {
    var me = this;
    this.canSelect = true;
    $scope.selectedCards = [];
    this.cards = [0,1,2,3,5,8,16,20,40,100,'?'];

    var ws = new WebSocket("ws://localhost:8080/entry");

    this.click = function(card) {
        console.log(card);
        this.canSelect = false;
        var message = JSON.stringify({author: "hoge", card: card.toString()});
        console.log(message);
        ws.send(message);
    };

    this.reset = function() {
        var message = JSON.stringify({author: "hoge", reset: true});
        console.log(message);
        ws.send(message);
    };

    this.open = function() {
        if ($scope.selectedCards.length) {
            var message = JSON.stringify({author: "hoge", open: true});
            ws.send(message);
        }
    };

    ws.onmessage = function(e) {
        console.log(e.data);
        var card = JSON.parse(e.data);
        $scope.$apply(function() {
            if (card.reset) {
                $scope.selectedCards.length = 0;
                me.canSelect = true;
                me.openCards = false;
            } else if (card.open) {
                me.openCards = true;
            } else {
                $scope.selectedCards.push(JSON.parse(e.data));
            }
        });
    };

}]);
