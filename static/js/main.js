require.config({
	baseUrl: 'static/js/vendor',
	paths: {
		jquery: 'jquery',
		angular: 'angular',
		angularRoute: 'angular-route',
		fullPage: 'jquery.fullPage',
		semantic: 'semantic',
		slimscroll: 'jquery.slimscroll.min'
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
			deps: ['jquery', 'slimscroll'],
			exports: 'fullPage'
		}
	}	
});

window.name = "NG_DEFER_BOOTSTRAP";

require(['angular', 'angularRoute', 'fullPage'], function(angular) {
	$('#fullpage').fullpage({
		anchors: ['firstPage', 'secondPage', '3rdPage'],
		sectionsColor: ['#C63D0F', '#1BBC9B', '#7E8F7C'],
		css3: true
	});

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
