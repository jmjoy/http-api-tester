http-api-tester:
	go build && ./http-api-tester

create-view:
	file2string \
		static/bower_components/bootstrap/dist/**/*.min.* \
		static/bower_components/bootstrap/fonts/* \
		static/bower_components/bootstrap-select/dist/**/*.min.* \
		static/bower_components/bootstrap-switch/dist/**/*.min.* \
		static/bower_components/bootstrap-switch/dist/css/bootstrap3/*.min.* \
		static/bower_components/handlebars/handlebars.min.js \
		static/bower_components/jquery/dist/jquery.min.js \
		static/*.* \
		view/* \
		favicon.b64
