/**
 * Created by RoryGao on 16/1/12.
 */
define(['./module'], function (module) {
    'use strict';
    module.controller('ArticleController', ArticleController);

    ArticleController.$inject = ['$scope', '$window', '$uibModal', '$stateParams','ArticleRestService', 'AlertService'];

    function ArticleController($scope, $window, $uibModal, $stateParams, ArticleRestService, AlertService) {
        var num = $stateParams.num;
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
                    $window.location.href = '/admin/article';
                }, function () {
                    AlertService.showAlert();
                });
            });
        };
    }

    module.controller("UpdateArticleController", UpdateArticleController);

    UpdateArticleController.$inject = ['$scope', '$window', '$stateParams', 'ArticleRestService', 'CategoryRestService'];

    function UpdateArticleController($scope, $window, $stateParams, ArticleRestService, CategoryRestService) {

        $scope.$on('$viewContentLoaded', function() {
            $("textarea[data-provide='markdown']").markdown({
                language: 'zh',
                fullscreen: {
                    enable: true
                },
                resize: 'vertical'
            });
        });
        CategoryRestService.getCategories().then(
            function (result) {
                $scope.categories = result.data.data;
            }
        );
        var isUpdate = false;
        $scope.article = {};
        if ($stateParams.op === 'update') {
            ArticleRestService.getArticleInfo($stateParams.id).then(
                function (result) {
                    var article = result.data.data;
                    $scope.article.title = article.title;
                    $scope.article.category_id = article.category.id;
                    $scope.article.content = article.content;
                }
            );
            $scope.article.id = $stateParams.id;
            isUpdate = true;
        }

        $scope.SaveOrUpdateArticle = function () {
            ArticleRestService.addOrUpdateArticle($scope.article, isUpdate).then(
                function () {
                    $window.location.href = '/admin/article';
                }
            )
        }
    }
});