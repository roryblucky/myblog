/**
 * Created by RoryGao on 16/1/12.
 */
define(['./module'], function(module) {
   module.factory('AlertService', ['$rootScope', '$timeout', function ($rootScope, $timeout) {

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
});