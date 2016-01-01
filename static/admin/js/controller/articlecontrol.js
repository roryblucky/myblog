/**
 * Created by RoryGao on 15/12/31.
 */

blogMain.controller('ArticleController', ArticleController);

ArticleController.$inject = ['$scope', '$uibModal', 'ArticleRestService', 'AlertService'];

function ArticleController($scope, $uibModal, ArticleRestService, AlertService) {
    ArticleRestService.getArticles(1).then(function (result) {
        $scope.articles = result.data.data;
    });

    $scope.openDelDialog = function ($index) {
        var modalInstance = $uibModal.open({
            templateUrl: '/static/admin/view/article/article_del_dialog.html',
            controller: 'DelDialogController'
        });

        modalInstance.result.then(function () {
            ArticleRestService.delArticle($scope.articles[$index]).then(function () {
                $scope.articles.splice($index, 1);
            }, function () {
                AlertService.showAlert();
            });
        });
    };
}