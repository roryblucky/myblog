/**
 * Created by RoryGao on 15/12/30.
 */
angular.module('myblog.alert', []).factory('AlertService', ['$rootScope', '$timeout', function ($rootScope, $timeout) {

    function closeAlert() {
        $rootScope.isShow = false;
    }

    return {
        showAlert: function () {
            $rootScope.isShow = true;
            $timeout(function () {
                closeAlert();
            }, 2000)
        }
    }
}]);