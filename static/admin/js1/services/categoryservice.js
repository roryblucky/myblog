/**
 * Created by RoryGao on 16/1/12.
 */
define(['./module'], function(module) {
    'use strict';
    module.factory('CategoryRestService', ['$http', '$q','$httpParamSerializerJQLike', function ($http, $q, $httpParamSerializerJQLike) {
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
    }]);
});