/**
 * Created by RoryGao on 15/12/28.
 */
var blogMain = angular.module('myblog-main', ['ngRoute','ngFileUpload']);

blogMain.config(['$routeProvider', '$locationProvider', function ($routeProvider, $locationProvider) {
    $routeProvider
        .when('/admin/article', {
            templateUrl: '/view/article/article_main.html'
        })
        .when('/admin/article/add', {
            templateUrl: '/view/article/article_info.html'
        })
        .when('/admin/category', {
            templateUrl: '/view/category/category_main.html'
        })
        .when('/admin/blogowner', {
            templateUrl: '/view/blogowner/blogonwer_info.html',
            controller:'BlogUploadController'
        })
        .otherwise({
            redirectTo: '/admin/article'
        });
    $locationProvider.html5Mode(true);
}]);
