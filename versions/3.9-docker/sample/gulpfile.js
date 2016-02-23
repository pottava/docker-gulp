"use strict";

var gulp = require('gulp');
var shell = require('gulp-shell');

gulp.task("restart", function() {
  gulp.src("/app/gulpfile.js") // any file
      .pipe(shell(['docker restart app']));
});

gulp.task("default", function() {
    gulp.watch("/go/src/github.com/pottava/docker-gulp/versions/3.9-docker/sample/app/**/*.go", ["restart"]);
});
