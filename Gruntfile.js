/*global module:false*/
module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    // Metadata.
    pkg: grunt.file.readJSON('package.json'),
    // Task configuration.
    uglify: {
      options: {
        mangle: true
      },
      dist: {
        files: {
          'static/js/myblog_front.min.js': ['static/js/article.js','static/js/markdown.js'],
          'static/js/myblog_backend.min.js': ['static/admin/js/**/*.js','!static/admin/js/**/logincontrol.js']
        }
      }
    },
    cssmin: {
      target: {
        files: {
          'static/css/myblog_backend.min.css':['static/admin/css/*.css', '!static/admin/css/form.css'],
          'static/css/main.min.css':['static/css/main.css']
        }
      }
    }
  });

  // These plugins provide necessary tasks.
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-cssmin');

  // Default task.
  grunt.registerTask('default', ['cssmin', 'uglify']);

};
