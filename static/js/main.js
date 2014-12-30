'use strict';

require.config({
	paths: {
		'jquery': 'vendor/jquery',
		'angular': 'vendor/angular',
		'angular-route': 'vendor/angular-route',
		'full-page': 'vendor/jquery.fullPage',
		'semantic': 'vendor/semantic',
		'slimscroll': 'vendor/jquery.slimscroll.min'
	},
	shim: {
		'angular': {
			deps: ['jquery'],
			exports: 'angular'
		},
		'angular-route': {
			deps: [ 'angular' ]
		},
		'semantic': {
			deps: ['jquery'],
			exports: 'semantic'
		},
		'slimscroll': {
			deps: ['jquery']
		},
		'full-page': {
			deps: ['slimscroll'],
			exports: 'full-page'
		}
	}
});

window.name = "NG_DEFER_BOOTSTRAP!";

require(['angular', 'app', 'routes', 'full-page'], function(angular, app, routes){
    var $html = angular.element(document.getElementsByTagName('html')[0]);

	angular.element().ready(function() {
		angular.resumeBootstrap([app['name']]);
	});

	var transform = function (data) {
	    return $.param(data);
	};

	var config = {
	    headers: { 'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'},
	    transformRequest: transform
	};

	$('#fullpage').fullpage({
		'css3': true,
		'easingcss3': 'cubic-bezier(0.175, 0.885, 0.320, 1.275)'
	});	

});	