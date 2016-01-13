/**
 * Created by RoryGao on 16/1/12.
 */
define(['./module'], function(module) {
   'use strict';
    module.factory('ArticleRestService', ['$http', '$q', '$httpParamSerializer', function ($http, $q, $httpParamSerializer) {

        var addOrUpdateArticle = function (params, isUpdate) {
            var deferred = $q.defer();
            var _url = '/api/admin/article/add';
            if (isUpdate) {
                _url = '/api/admin/article/update/' + params.id;
            }
            $http({
                method: 'POST',
                url: _url,
                headers: {
                    'Content-type' : 'application/x-www-form-urlencoded'
                },
                data: $httpParamSerializer({
                    title: params.title,
                    content: params.content,
                    category_id: params.category_id
                })
            }).then(function (data) {
                deferred.resolve(data);
            }, function (reason) {
                deferred.reject(reason);
            });
            return deferred.promise;
        };

        var delArticle = function (params) {
            var deferred = $q.defer();
            $http({
                method: 'Get',
                url: '/api/admin/article/del/' + params.id
            }).then(function (data) {
                deferred.resolve(data);
            }, function (reason) {
                deferred.reject(reason);
            });
            return deferred.promise;
        };


        var getArticles = function (pageNum) {
            var deferred = $q.defer();
            $http({
                method: 'GET',
                url: '/api/admin/articles/' + pageNum
            }).then(function (data) {
                deferred.resolve(data);
            }, function (reason) {
                deferred.reject(reason);
            });
            return deferred.promise;
        };

        var getArticleInfo = function(id) {
            var deferred = $q.defer();
            $http({
                method: 'GET',
                url: '/api/admin/article/' + id
            }).then(function (data) {
                deferred.resolve(data);
            }, function (reason) {
                deferred.reject(reason);
            });
            return deferred.promise;
        };

        return {
            addOrUpdateArticle: addOrUpdateArticle,
            delArticle: delArticle,
            getArticles: getArticles,
            getArticleInfo: getArticleInfo
        };
    }]);
});