/**
 * Created by RoryGao on 15/12/30.
 */

blogMain.controller('CategoryController', CategoryController).controller('updateCategoryDialogController', updateCategoryDialogController);

CategoryController.$inject = ['$scope', '$uibModal', '$window', 'CategoryRestService', 'AlertService'];

function CategoryController($scope, $uibModal, $window, CategoryRestService, AlertService) {

    CategoryRestService.getCategories().then(
        function (result) {
            $scope.categories = result.data;
        }
    );

    var openDelDialog = function ($index) {
        var modalInstance = $uibModal.open({
            templateUrl: '/static/admin/view/category/category_del_dialog.html',
            controller: 'DelDialogController'
        });

        modalInstance.result.then(function () {
            CategoryRestService.delCategory($scope.categories[$index]).then(function () {
                $scope.categories.splice($index, 1);
            }, function (err) {
                AlertService.showAlert();
            });
        });
    };

    var openUpdateDialog = function ($index) {
        var modelInstance = $uibModal.open({
            templateUrl: '/static/admin/view/category/category_addupdate_dialog.html',
            controller: 'updateCategoryDialogController',
            resolve: {
                categoryName: function () {
                    if ($index != undefined) {
                        return $scope.categories[$index].name;
                    }
                }
            }
        });
        if ($index != undefined) {
            modelInstance.result.then(function (categoryName) {
                var oldCategory = $scope.categories[$index];
                oldCategory.name = categoryName;
                CategoryRestService.updateCategory(oldCategory).then(function () {
                    $window.location.href = '/admin/category';
                });
            });
        } else {
            modelInstance.result.then(function (categoryName) {
                var params = {name: categoryName};
                CategoryRestService.addCategory(params).then(function () {
                    $window.location.href = '/admin/category';
                });
            });
        }
    };

    $scope.openDelDialog = openDelDialog;
    $scope.openUpdateDialog = openUpdateDialog;
}

updateCategoryDialogController.$inject = ['$scope', '$uibModalInstance', 'categoryName'];

function updateCategoryDialogController($scope, $uibModalInstance, categoryName) {
    $scope.categoryName = categoryName;

    $scope.ok = function () {
        $uibModalInstance.close($scope.categoryName);
    };

    $scope.cancel = function () {
        $uibModalInstance.dismiss();
    };
}