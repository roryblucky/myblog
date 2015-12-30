/**
 * Created by RoryGao on 15/12/29.
 */
angular.module('myblog.services',[]).factory('CategoryRestService', ['$http', '$q','$httpParamSerializerJQLike', function ($http, $q, $httpParamSerializerJQLike) {


    var addCategory = function (params) {
        var deferred = $q.defer();
        $http({
            method: 'POST',
            url: '/api/admin/category/add',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
            },
            data: $httpParamSerializerJQLike({
                name: params.name
            })
        }).then(function (data) {
            deferred.resolve(data);
        },function (reason) {
            deferred.reject(reason);
        });
        return deferred.promise;
    };

    var delCategory = function (params) {
        var deferred = $q.defer();
        $http({
            method: 'GET',
            url: '/api/admin/category/del/' + params.id
        }).then(function (data) {
            deferred.resolve(data);
        },function (reason) {
            deferred.reject(reason);
        });
        return deferred.promise;
    };

    var updateCategory = function (params) {
        var deferred = $q.defer();
        $http({
            method: 'POST',
            url: '/api/admin/category/update/' + params.id,
            headers: {
                'Content-Type':'application/x-www-form-urlencoded'
            },
            data: $httpParamSerializerJQLike({
                name: params.name
            })
        }).success(function (data) {
            deferred.resolve(data);
        }).error(function (reason) {
            deferred.reject(reason);
        });
        return deferred.promise;
    };

    var getCategories = function () {
        var deferred = $q.defer();
        $http({
            method: 'GET',
            url: '/api/admin/categories'
        }).success(function (data) {
            deferred.resolve(data);
        }).error(function (reason) {
            deferred.reject(reason);
        });
        return deferred.promise;
    };

    return {
        addCategory: addCategory,
        delCategory: delCategory,
        updateCategory: updateCategory,
        getCategories: getCategories
    };
}]).factory('ArticleRestService', ['$http', '$q', function ($http, $q) {

    var addArticle = function (params) {
        var deferred = $q.defer();
        $http({
            method: 'POST',
            url: '/api/admin/article/add',
            data: {
                title: params.title,
                content: params.content,
                category_id: params.category_id
            }
        }).success(function (data) {
            deferred.resolve(data);
        }).error(function (reason) {
            deferred.reject(reason);
        });
        return deferred.promise;
    };

    var delArticle = function (params) {
        var deferred = $q.defer();
        $http({
            method: 'Get',
            url: '/api/admin/article/del/' + params.id
        }).success(function (data) {
            deferred.resolve(data);
        }).error(function (reason) {
            deferred.reject(reason);
        });
        return deferred.promise;
    };

    var updateArticle = function (params) {
        var deferred = $q.defer();
        $http({
            method: 'POST',
            url: '/api/admin/article/update/' + params.id,
            data: {
                title: params.title,
                content: params.content,
                category_id: params.category_id
            }
        }).success(function (data) {
            deferred.resolve(data);
        }).error(function (reason) {
            deferred.reject(reason);
        });
        return deferred.promise;
    };

    var getArticles = function () {
        var deferred = $q.defer();
        $http({
            method: 'GET',
            url: '/api/admin/articles'
        }).success(function (data) {
            deferred.resolve(data);
        }).error(function (error) {
            deferred.reject(error);
        });
        return deferred.promise;
    };

    return {
        addArticle: addArticle,
        delArticle: delArticle,
        updateArticle: updateArticle,
        getArticles: getArticles
    };
}]);