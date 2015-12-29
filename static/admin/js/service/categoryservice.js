/**
 * Created by RoryGao on 15/12/29.
 */
blogMain.factory('CategoryRestService', ['$http', '$q', function ($http, $q) {
    var deferred = $q.defer();

    var addCategory = function (params) {
        $http({
            method: 'POST',
            url: '/api/admin/category/add',
            data: {
                name: params.name
            }
        }).success(function (data) {
            deferred.resolve(data);
        }).error(function (reason) {
            deferred.reject(reason);
        });
        return deferred.promise;
    };

    var delCategory = function (params) {
        $http({
            method: 'GET',
            url: '/api/admin/category/del/' + params.categoryid
        }).success(function (data) {
            deferred.resolve(data);
        }).error(function (reason) {
            deferred.reject(reason);
        });
        return deferred.promise;
    };

    var updateCategory = function (params) {
        $http({
            method: 'POST',
            url: '/api/admin/category/update/' + params.categoryid,
            data: {
                name: params.name
            }
        }).success(function (data) {
            deferred.resolve(data);
        }).error(function (reason) {
            deferred.reject(reason);
        });
        return deferred.promise;
    };

    var getCategories = function () {
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
}]);