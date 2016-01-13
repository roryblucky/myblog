/**
 * Created by RoryGao on 16/1/12.
 */
require.config({
    paths: {
        'angular': 'bower_components/angular/angular.min',
        'angularAnimate': 'bower_components/angular-animate/angular-animate.min',
        'angularBootstrap': 'bower_components/angular-bootstrap/ui-bootstrap-tpls.min',
        'angularUIRouter': 'bower_components/angular-ui-router/release/angular-ui-router.min',
        'ngFileUpload': 'bower_components/ng-file-upload/ng-file-upload.min',
        'domReady': 'bower_components/domReady/domReady.min'
    },
    shim: {
        'angular': {
            exports: 'angular'
        },
        'angularAnimate': {
            deps: ['angular']
        },
        'angularBootstrap': {
            deps: ['angular']
        },
        'angularUIRouter': {
            deps: ['angular']
        },
        'ngFileUpload': {
            deps: ['angular']
        }
    },
    urlArgs: 'bust=' + (new Date()).getTime(),
    deps: [
        './bootstrap'
    ],
    waitSeconds: 30
});