/**
 * Created by RoryGao on 16/1/12.
 */
define(['./module'],function(filters) {
   'use strict';
    filters.filter('range', function() {
        return function(input, total) {
            total = parseInt(total);
            for (var i=0; i<total; i++) {
                input.push(i);
            }
            return input;
        }
    });
});