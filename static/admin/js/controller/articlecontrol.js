/**
 * Created by RoryGao on 15/12/31.
 */

blogMain.controller('ArticleController', ArticleController);

ArticleController.$inject = ['$scope', '$uibModal', '$routeParams','ArticleRestService', 'AlertService'];

function ArticleController($scope, $uibModal, $routeParams, ArticleRestService, AlertService) {
    var num = $routeParams.num;
    if (num == undefined) {
        num = 1;
    }
    ArticleRestService.getArticles(num).then(function (result) {
        $scope.articles = result.data.data;
        $scope.totalPages = result.data.totalPages;
        $scope.hasPrev = result.data.hasPrePage;
        $scope.hasNext = result.data.hasNextPage;
        $scope.currentPage = parseInt(num);
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