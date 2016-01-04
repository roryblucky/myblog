/**
 * Created by RoryGao on 15/12/29.
 */
BlogOwnerUpdate.$inject=['$scope','Upload'];
blogMain.controller('BlogUploadController', BlogOwnerUpdate);

function BlogOwnerUpdate($scope, Upload) {
    $scope.blogowner = {};

    var update = function() {
        var _data = {};
        if ($scope.blogowner.icon != undefined) {
            _data.icon = $scope.blogowner.icon;
        }

        if ($scope.blogowner.intro != undefined) {
            _data.introduction = $scope.blogowner.intro;
        }

        Upload.upload({
            url: '/api/admin/blogowner/update',
            data:_data
        }).success(function() {
            $scope.blogowner.result = '更新成功';
        }).error(function(error) {
            console.log(error);
            $scope.blogowner.result = '更新失败';
            $scope.blogowner.isErr = true;
        });
    };
    $scope.update = update;
}
