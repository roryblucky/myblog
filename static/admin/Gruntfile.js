/**
 * Created by RoryGao on 15/12/25.
 */
module.exports = function (grunt) {
    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json'),

        jshint: {
            options: {
                curly: true,
                eqeqeq: true,
                eqnull: true,
                unused: true
            },
            all:['/js/**/*.js']
        },

        bower: {
            install: {
                options: {
                    targetDir: '/vendor',
                    cleanup: true,
                    layout: 'byComponent',
                    install: true
                }
            }
        },

        //connect: {
        //    server: {
        //        options: {
        //            port: 9000,
        //            base: 'dist',
        //            keepalive: true
        //        }
        //    },
        //    dev: {
        //        options: {
        //            port: 8000,
        //            base: 'src',
        //            keepalive: true
        //        }
        //    }
        //}
    });

    grunt.loadNpmTasks('grunt-contrib-jshint');
    //grunt.loadNpmTasks('grunt-contrib-connect');
    grunt.loadNpmTasks('grunt-bower-task');


    grunt.registerTask('dev',['jshint', 'bower', 'connect:dev']);
};