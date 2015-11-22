http-api-tester:
	go build && ./http-api-tester

create-view:
	file2string bower_components/bootstrap/dist/**/*.min.* bower_components/bootstrap-select/dist/**/*.min.* bower_components/bootstrap-switch/dist/**/*.min.* bower_components/handlebars/handlebars.min.js bower_components/jquery/dist/jquery.min.js  static/* view/*
