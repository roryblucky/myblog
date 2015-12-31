/**
 * Created by RoryGao on 15/12/30.
 */
angular.module('myblog.alert', ['ngAnimate']).factory('AlertService', ['$rootScope', '$timeout', function ($rootScope, $timeout) {

    $rootScope.closeAlert = function () {
        $rootScope.isShow = false;
    };

    return {
        showAlert: function () {
            $rootScope.isShow = true;
            $timeout(function () {
                $rootScope.closeAlert();
            }, 2000)
        }
    }
}]);