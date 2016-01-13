/**
 * Created by RoryGao on 16/1/12.
 */
require.config({
    paths: {
        'angular': 'bower_components/angular/angular',
        'angularAnimate': 'bower_components/angular-animate/angular-animate',
        'angularBootstrap': 'bower_components/angular-bootstrap/ui-bootstrap-tpls',
        'angularUIRouter': 'bower_components/angular-ui-router/release/angular-ui-router',
        'ngFileUpload': 'bower_components/ng-file-upload/ng-file-upload',
        'domReady': 'bower_components/domReady/domReady'
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
    ]
});