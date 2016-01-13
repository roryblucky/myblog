/**
 * Created by RoryGao on 16/1/12.
 */
define([
    'angular',
    'angularBootstrap',
    'ngFileUpload',
    '../services/myblogservice'
], function(angular){
   'use strict';
    return angular.module('myblog.controllers', ['ui.bootstrap', 'ngFileUpload', 'myblog.services']);
});