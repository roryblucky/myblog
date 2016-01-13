/**
 * Created by RoryGao on 16/1/12.
 */
define([
	'require',
	'angular',
	'app',
	'router'
], function(require, angular) {
	require(['domReady!'], function(document) {
		angular.bootstrap(document, ['myblog']);
	});
});