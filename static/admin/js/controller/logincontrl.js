/**
 * Created by RoryGao on 15/12/25.
 */
(function () {
    'use strict';
    angular.module('myblog.loginform',[]).controller('LoginController', ['$scope', '$http', AdminLogin]);

    function AdminLogin($scope, $http) {
        $scope.loginform = {};
        var loginUrl = "/admin/login";
        var login = function() {
           $http({
               method: 'Get',
               url: loginUrl,
               data:{
                   'username':$scope.username,
                   'password':$scope.password
               }
           }).success(function() {
               $scope.loginform.result = '登录成功';
           }).error(function() {
               $scope.loginform.result = '用户名或密码错误';
               $scope.loginform.isErr = true;
           });
        };
        $scope.login = login;
    }
})();