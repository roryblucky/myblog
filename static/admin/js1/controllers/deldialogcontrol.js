/**
 * Created by RoryGao on 16/1/12.
 */
define(['./module'], function(module) {
   'use strict';
    module.controller('DelDialogController', ['$scope', '$uibModalInstance', function ($scope, $uibModalInstance) {
        $scope.ok = function () {
            $uibModalInstance.close();
        };

        $scope.cancel = function () {
            $uibModalInstance.dismiss();
        };
    }]);
});