/**
 * Created by RoryGao on 16/1/12.
 */
define(['app'], function(app) {
	'use strict';
	app.config(['$stateProvider', '$urlRouterProvider', '$locationProvider', function($stateProvider, $urlRouterProvider, $locationProvider) {
		$locationProvider.html5Mode(true);

		$urlRouterProvider.otherwise('article');

		$stateProvider
			.state('article', {
				url: '/article',
				templateUrl: '/static/admin/view/article/article_main.html',
				controller: 'ArticleController'
			})
			.state('updateArticle', {
				url: '/article/:op?id',
				templateUrl: '/static/admin/view/article/article_info.html',
				controller: 'UpdateArticleController'
			})
			.state('articlepage', {
				url: '/article/page/:num',
				templateUrl: '/static/admin/view/article/article_main.html',
				controller: 'ArticleController'
			})
			.state('category', {
				url: '/category',
				templateUrl: '/static/admin/view/category/category_main.html',
				controller: 'CategoryController'
			})
			.state('blogowner', {
				url: '/blogowner',
				templateUrl: '/static/admin/view/blogowner/blogonwer_info.html',
				controller: 'BlogUploadController'
			});

	}]);
});