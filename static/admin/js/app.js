/**
 * Created by RoryGao on 15/12/28.
 */
var blogMain = angular.module('myblog-main', ['ngRoute', 'ngAnimate', 'ngFileUpload', 'ui.bootstrap', 'myblog.services', 'myblog.alert']);

blogMain.config(['$routeProvider', '$locationProvider', function ($routeProvider, $locationProvider) {
    $routeProvider
        .when('/admin/article', {
            templateUrl: '/static/admin/view/article/article_main.html',
            controller: 'ArticleController'
        })
        .when('/admin/article/:op', {
            templateUrl: '/static/admin/view/article/article_info.html',
            controller: 'UpdateArticleController'
        })
        .when('/admin/category', {
            templateUrl: '/static/admin/view/category/category_main.html',
            controller: 'CategoryController'
        })
        .when('/admin/blogowner', {
            templateUrl: '/static/admin/view/blogowner/blogonwer_info.html',
            controller:'BlogUploadController'
        })
        .when('/admin/article/page/:num', {
            templateUrl: '/static/admin/view/article/article_main.html',
            controller: 'ArticleController'
        })
        .otherwise({
            redirectTo: '/admin/article'
        });
    $locationProvider.html5Mode(true);
}]);

//del dialog controller
blogMain.controller('DelDialogController', ['$scope', '$uibModalInstance', function ($scope, $uibModalInstance) {
    $scope.ok = function () {
        $uibModalInstance.close();
    };

    $scope.cancel = function () {
        $uibModalInstance.dismiss();
    };
}]);

blogMain.filter('range', function() {
   return function(input, total) {
       total = parseInt(total);
       for (var i=0; i<total; i++) {
           input.push(i);
       }

       return input;
   }
});