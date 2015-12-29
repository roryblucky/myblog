/**
 * Created by RoryGao on 15/12/25.
 */
(function () {
    'use strict';
    angular.module('myblog.loginform', []).controller('LoginController', ['$scope', '$http', '$window', AdminLogin]);

    function AdminLogin($scope, $http, $window) {
        $scope.loginform = {};
        var loginUrl = "/admin/login";
        var login = function () {
            $http({
                method: 'POST',
                url: loginUrl,
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
                },
                data: $.param($scope.loginform)
            }).success(function (data) {
                if (data.code != 200) {
                    $scope.loginform.result = '用户名或密码错误';
                    $scope.loginform.isErr = true;
                } else {
                    $scope.loginform.result = '登录成功';
                    $window.location.href = '/admin/main'
                }
            })
        };
        $scope.login = login;
    }
})();