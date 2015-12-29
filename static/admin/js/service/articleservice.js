/**
 * Created by RoryGao on 15/12/29.
 */
blogMain.factory('ArticleRestService', ['$http', '$q', function ($http, $q) {
    var deferred = $q.defer();

    var addArticle = function (params) {
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
        $http({
            method: 'Get',
            url: '/api/admin/article/del/' + params.articleId
        }).success(function (data) {
            deferred.resolve(data);
        }).error(function (reason) {
            deferred.reject(reason);
        });
        return deferred.promise;
    };

    var updateArticle = function (params) {
        $http({
            method: 'POST',
            url: '/api/admin/article/update/' + params.articleId,
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
