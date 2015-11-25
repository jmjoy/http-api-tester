http-api-tester:
	go build && ./http-api-tester

create-view:
	file2string -pkg text -o text/text.go -var Text \
		static/bower_components/bootstrap/dist/**/*.min.* \
		static/bower_components/bootstrap/dist/fonts/* \
		static/bower_components/bootstrap-select/dist/**/*.min.* \
		static/bower_components/bootstrap-select/dist/**/*.map \
		static/bower_components/bootstrap-switch/dist/**/*.min.* \
		static/bower_components/bootstrap-switch/dist/css/bootstrap3/*.min.* \
		static/bower_components/handlebars/handlebars.min.js \
		static/bower_components/jquery/dist/jquery.min.js \
		static/bower_components/jquery/dist/jquery.min.map \
		static/*.* \
		view/* \
		favicon.ico
