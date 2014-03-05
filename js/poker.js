var app = angular.module('poker', ['ngAnimate','websocketServices'])

app.controller('PokerController', ['$scope', 'Poker', function($scope, Poker) {
    var me = this;
    this.canSelect = true;
    $scope.selectedCards = [];
    this.cards = [0,'1/2',1,2,3,5,8,16,20,40,100,'?'];

    this.click = function(card) {
        this.canSelect = false;
        Poker.send({author: this.author, card: card.toString()});
    };

    this.reset = function() {
        Poker.send({author: this.author, reset: true});
    };

    this.open = function() {
        if ($scope.selectedCards.length) {
            Poker.send({author: this.author, open: true});
        }
    };

    Poker.listen(function(e) {
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
    });

}]);
