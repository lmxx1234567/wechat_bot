angular.module('loginApp', [])
    .controller('LoginController', function ($scope, $http) {
        $http({
            method:'GET',
            url:'/getQRcode'
        }).then(function successCallback(response){
            $scope.code = response.data.code
        },function (response){});
    });