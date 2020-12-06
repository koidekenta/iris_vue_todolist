const gulp = require("gulp");

const sass = require("gulp-sass");

gulp.task("default", function(){
	return(
		gulp
		.src("asset/scss/style.scss")
		.pipe(sass({outputStyle: "expanded"}))
		.pipe(gulp.dest("css"))
	);
});