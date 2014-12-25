require.config({
	baseUrl: 'static/js/vendor',
	paths: {
		jquery: 'jquery',
		angular: 'angular',
		angularRoute: 'angular-route',
		fullPage: 'jquery.fullPage',
		semantic: 'semantic'
	},
	shim: {
		angular: {
			deps: ['jquery'],
			exports: 'angular'
		},
		angularRoute: {
			deps: ['angular']
		},
		semantic: {
			deps: ['jquery']
		},
		fullPage: {
			deps: ['jquery']
		}
	}	
});

window.name = "NG_DEFER_BOOTSTRAP";

require(['angular', 'angularRoute', 'semantic', 'fullPage'], function(angular) {
	var app = angular.module('sifan', ['ngRoute']);

	app.config(['$routeProvider', function ($routeProvider) {
	    $routeProvider.otherwise({redirectTo: '/home'});
	}]);

	var transform = function (data) {
	    return $.param(data);
	};

	var config = {
	    headers: { 'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'},
	    transformRequest: transform
	};
});
