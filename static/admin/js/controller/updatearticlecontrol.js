/**
 * Created by RoryGao on 15/12/31.
 */
blogMain.controller("UpdateArticleController", UpdateArticleController);

UpdateArticleController.$inject = ['$scope', '$window', '$routeParams', 'ArticleRestService', 'CategoryRestService'];

function UpdateArticleController($scope, $window, $routeParams, ArticleRestService, CategoryRestService) {

    $scope.$on('$viewContentLoaded', configMarkdown);
    CategoryRestService.getCategories().then(
        function (result) {
            $scope.categories = result.data.data;
        }
    );
    var isUpdate = false;
    $scope.article = {};
    if ($routeParams.op === 'update') {
        ArticleRestService.getArticleInfo($routeParams.id).then(
            function (result) {
                var article = result.data.data;
                $scope.article.title = article.title;
                $scope.article.category_id = article.category.id;
                $scope.article.content = article.content;
            }
        );
        $scope.article.id = $routeParams.id;
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