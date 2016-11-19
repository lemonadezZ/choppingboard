var gulp = require('gulp');
var webserver = require('gulp-webserver');

gulp.task('default', function() {
  // place code for your default task here
});
//编译
gulp.task('build', function() {
  // place code for your default task here
});

gulp.task('serve', function() {
  gulp.src('.')
    .pipe(webserver({
      livereload: true,
      directoryListing: true,
      open: true,
      fallback: 'index.html'
    }));
});