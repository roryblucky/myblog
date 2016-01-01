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
        }).then(function (data) {
            deferred.resolve(data);
        }, function (reason) {
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
}]).factory('ArticleRestService', ['$http', '$q', '$httpParamSerializer', function ($http, $q, $httpParamSerializer) {

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