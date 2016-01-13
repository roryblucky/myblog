/**
 * Created by RoryGao on 16/1/12.
 */
define([
    'angular',
    'angularAnimate',
    'angularUIRouter',
    './controllers/controllers',
    './services/myblogservice',
    './filter/filters'
], function (angular) {
    'use strict';
    return angular.module('myblog', ['ngAnimate', 'ui.router', 'myblog.controllers', 'myblog.services','myblog.filters']);
});