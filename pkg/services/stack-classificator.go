package services

import (
	"github.com/demas/cowl-go/pkg/quzx-crawler"
	"fmt"
)

var stop_tags = []string{"php", "django", "cuda", "ionic2", "ionic-framework", "cordova", "wordpress", "laravel", "winforms", "mysql", "forms", "unity3d",
	"google-app-engine", "kendo-grid", "angularfire2", "firebase", "electron", "facebook", "pascal", "office365", "salesforce", "crm",
	"knockout.js", "lodash", "extjs", "swing", "javafx", "groovy", "neo4j", "google-cloud-storage", "pygame", "scene2d", "tkinter", "wildfly-9",
	"twilio", "tfs", "tfs2017", "jquery", "struts2", "kendo-ui", "vue.js", "vuejs2", "sharepoint", "vbscript", "dynamics-crm",
	"sap", "highcharts", "three.js", "d3.js", "ember.js", "vba", "sharepoint-2010", "vb.net", "amazon-web-services", "excel", "devexpress",
	"datagridview", "crystal-reports", "google-apps-script", "charts.js", "signalr", "carousel", "firebase-cloud-messaging", "word",
	"cloudkit", "facebook-graph-api", "amazon-web-services", "gruntjs", "regex", "handlebars.js", "google-apps-script",
	"laravel-5.1", "openwrt", "identityserver3", "wildfly", "nhibernate", "apache-cayenne", "apache-spark", "opencv", "alexa", "jq",
	"playframework", "codenvy", "botframework", "mobx", "eclipse", "transpose", "phantomjs", "ruby-on-rails", "flask", "ibm-mq",
	"docplexcloud", "underscore.js", "youtube-iframe-api", "actionscript-3", "gis", "openlayers", "devextreme", "google-maps", "js-pdf",
	"google-maps-api-3", "parse.com", "react-native", "woocommerce", "amazon-s3", "jhipster", "chart.js", "openerp",
	"primeng", "nativescript", "ckeditor4.x", "pentaho", "postman", "mondrian", "instagram", "jspdf", "socialshare", "markojs", "jrecorder",
	"slack", "google-analytics", "automapper-5", "unity2d", "polly", "excel-interop", "unity5", "sharppcap", "tibco-ems", "sonarqube",
	"protocol-buffers", "heroku", "sails.js", "waterline", "certbot", "snapchat", "feathersjs", "orientdb", "ogg", "ibm-mq", "pdfkit",
	"ffmpeg", "dreamfactory", "ms-access", "wso2esb", "ms-access-2013", "cakephp", "cygwin", "amazon-redshift", "aurelia", "picturefill",
	"ruby-on-rails-4", "slider", "plotly", "owl-carousel", "vis.js", "tizen", "geoserver", "mapbox", "froala", "navision",
	"dropbox-api", "mono", "hangfire", "vsix", "exchange-server", "sharepoint-2013", "flux", "webforms", "imagemagick", "google-sheets-api",
	"flash", "dailymotion-api", "openstreetmap", "rhino", "google-trends", "angular-dart", "angular2-dart", "payara", "polymer", "cloudfoundry",
	"c3.js", "joomla", "c++", "selenium", "onedrive-api", "visual-foxpro", "outlook", "selenium-webdriver", "mirth", "video.js",
	"imagemagick-convert", "laravel-5", "aforge", "square-connect", "arduino", "alloy-ui", "mojolicious", "mithril.js", "opentok",
	"google-spreadsheet", "iframe", "google-maps-markers", "pouchdb", "leaflet", "soundcloud", "jtable", "netbeans", "kotlin", "affdex-sdk",
	"mono.cecil", "powerpoint", "exchangewebservices", "amcharts", "meteor", "jasper-reports", "sqlalchemy", "solaris",
	"ghprb", "bitbucket", "google-search-console", "openshift", "activemq", "solr", "voicemail", "pug",
	"ejs", "grpc", "intel-xdk", "aws-lambda", "okta", "loopbackjs", "flowtype", "visualforce", "jsfeat", "tampermonkey",
	"symfony", "tweenmax", "backbone.js", "nunjucks", "blueimp", "google-recaptcha", "discord", "google-visualization", "sweetalert",
	"bacon.js", "thymeleaf", "zebble", "gatt", "docusign", "signalr-hub", "aleagpu", "google-api", "paypal-ipn", "pdfsharp", "yii2", "qweb",
	"dojo", "ms-word", "tampermonkey", "jira", "sapui5", "jszip", "google-form", "jqgrid", "deployd", "sinatra", "mediaelement.js", "jwplayer",
	"ckeditor", "accordion", "masonry", "shopify", "google-drive-sdk", "whatsapp", "slick.js", "tinymce", "skulpt", "ormlite-servicestack",
	"postgis", "pgadmin-4", "pgadmin", "eclipselink", "oozie", "ibm-bluemix" }

type flr struct {
	Site string
	Include string
	Result string
}

var firstLevelRules = []flr	{
		{"superuser", "*", "superuser"},
		{"serverfault", "*", "administration"},
		{"unix", "*", "unix"},
		{"ru.stackoverflow", "*", "russian"},
		{"security", "*", "information security"},
		{ "codereview", "*", "code review" },
		{ "softwareengineering" ,"unit-testing" ,"unit-testing" },
		{ "softwareengineering" ,"design" ,"software-desing" },
		{ "stackoverflow" ,"elasticsearch" ,"bigdata" },
		{ "stackoverflow" ,"cassandra" ,"bigdata" },
		{ "stackoverflow" ,"apache-kafka" ,"bigdata" },
		{ "stackoverflow" ,"rabbitmq" ,"bigdata" },
		{ "stackoverflow" ,"iis" ,"devops" },
		{ "stackoverflow" ,"nginx" ,"devops" },
		{ "stackoverflow" ,"unit-testing" ,"unit-testing" },
		{ "stackoverflow" ,"machine-learning" ,"machine-learning" },
		{ "stackoverflow" ,"go" ,"go" },
		{ "stackoverflow" ,"azure" ,"azure" },
		{ "stackoverflow" ,"f#" ,"fsharp" },
		{ "stackoverflow" ,"postgresql" ,"postgresql" },
		{ "stackoverflow" ,"mongodb" ,"mongodb" },
		{ "stackoverflow" ,"clojure" ,"clojure" },
		{ "stackoverflow" ,"angular" ,"angular2" },
		{ "stackoverflow" ,"angular2" ,"angular2" },
		{ "stackoverflow" ,"angular4" ,"angular2" },
		{ "stackoverflow" ,"angular5" ,"angular2" },
		{ "stackoverflow" ,"git" ,"git" },
		{ "stackoverflow" ,"docker" ,"docker" },
		{ "stackoverflow" ,"rust" ,"rust" },
		{ "stackoverflow" ,"scala" ,"scala" },
		{ "stackoverflow" ,"sql-server" ,"ms sql server" },
		{ "stackoverflow" ,"oracle" ,"oracle" },
		{ "stackoverflow" ,"powershell" ,"powershell" },
		{ "stackoverflow" ,"xamarin" ,"xamarin" },
		{ "stackoverflow" ,"xamarin.forms" ,"xamarin" },
		{ "stackoverflow" ,"xamarin.ios" ,"xamarin" },
		{ "stackoverflow" ,"c#" ,".net" },
		{ "stackoverflow" ,".net" ,".net" },
		{ "stackoverflow" ,"svg" ,"web" },
		{ "stackoverflow" ,"twitter-bootstrap" ,"web" },
		{ "stackoverflow" ,"android" ,"android" },
		{ "stackoverflow" ,"gulp" ,"js" },
		{ "stackoverflow" ,"mocha" ,"js" },
		{ "stackoverflow" ,"chai" ,"js" },
		{ "stackoverflow" ,"webpack" ,"js" },
		{ "stackoverflow" ,"eslint" ,"js" },
		{ "stackoverflow" ,"react" ,"js" },
		{ "stackoverflow" ,"redux" ,"js" },
		{ "stackoverflow" ,"google-chrome-extension" ,"js" },
		{ "stackoverflow" ,"typescript" ,"typescript" },
		{ "stackoverflow" ,"node.js" ,"node.js" },
		{ "stackoverflow" ,"javascript" ,"js" },
		{ "stackoverflow" ,"bash" ,"bash" },
		{ "stackoverflow" ,"r" ,"r lang" },
	    { "stackoverflow" ,"spring" ,"java" },
		{ "stackoverflow" ,"java" ,"java" },
		{ "stackoverflow" ,"python" ,"python" },
		{ "stackoverflow" ,"haskell" ,"haskell" },
		{ "stackoverflow" ,"ios" ,"ios" },
		{ "stackoverflow" ,"swift" ,"ios" },
	}

type slr struct {
	Site string
	First string
	Include string
	Result string
}

var secondLevelRules = []slr{
	{"superuser", "superuser", "linux", "linux"},
	{"superuser", "superuser", "windows-10", "windows-10"},
	{"superuser", "superuser", "networking", "networking"},
	{"serverfault", "administration", "linux", "linux"},
	{"serverfault", "administration", "nginx", "nginx"},
	{"serverfault", "administration", "iis", "iis"},
	{"serverfault", "administration", "powershell", "powershell"},
	{"unix", "unix", "linux", "linux"},
	{"ru.stackoverflow", "russian", "java", "java"},
	{"ru.stackoverflow", "russian", "c#", "csharp"},
	{"ru.stackoverflow", "russian", "sql-server", "sql-server"},
	{"ru.stackoverflow", "russian", "git", "git"},
	{"stackoverflow", "azure", "azure-functions", "azure-functions"},
	{"stackoverflow", "azure", "machine-learning", "machine-learning"},
	{"stackoverflow", "azure", "azure-web-sites", "web-sites"},
	{"stackoverflow", "azure", "azure-logic-apps", "logic-apps"},
	{"stackoverflow", "azure", "sql-server", "sql-server"},
	{"stackoverflow", "azure", "azure-active-directory", "active directory"},
	{"stackoverflow", "azure", "active-directory", "active directory"},
	{"stackoverflow", "azure", "azure-documentdb", "azure-documentdb"},
	{"stackoverflow", "azure", "windows-azure-storage", "storage"},
	{"stackoverflow", "azure", "azure-media-services", "media-services"},
	{"stackoverflow", "azure", "azure-service-fabric", "fabric"},
	{"stackoverflow", "azure", "iot", "iot"},
	{"stackoverflow", "azure", "hdinsight", "hdinsight"},
	{"stackoverflow", "azure", "azure-cdn", "cdn"},
	{"stackoverflow", "azure", "asp.net-mvc", "asp.net"},
	{"stackoverflow", "azure", "servicebus", "servicebus"},
	{"stackoverflow", "azure", "azure-servicebus-queues", "servicebus"},
	{"stackoverflow", "azure", "azure-security", "security"},
	{"stackoverflow", "azure", "azure-virtual-machine", "vms"},
	{"stackoverflow", "ms sql server", "reporting-services", "ssrs"},
	{"stackoverflow", "ms sql server", "ssis", "ssis"},
	{"stackoverflow", "ms sql server", "database-design", "database-design"},
	{"stackoverflow", "ms sql server", "merge-replication", "replication"},
	{"stackoverflow", "ms sql server", "transactional-replication", "replication"},
	{"stackoverflow", "ms sql server", "maintenance", "maintenance"},
	{"stackoverflow", "ms sql server", "tsql", "tsql"},
	{"stackoverflow", "ms sql server", "xml", "xml"},
	{"stackoverflow", "ms sql server", "pivot", "powerbi"},
	{"stackoverflow", "ms sql server", "powerbi", "powerbi"},
	{"stackoverflow", "ms sql server", "stored-procedures", "stored-procedures"},
	{"stackoverflow", "postgresql", "performance", "optimization"},
	{"stackoverflow", "postgresql", "optimization", "optimization"},
	{"stackoverflow", "postgresql", "backup", "backup"},
	{"stackoverflow", "postgresql", "restore", "backup"},
	{"stackoverflow", "postgresql", "psql", "psql"},
	{"stackoverflow", "angular2", "angular-cli", "angular-cli"},
	{"stackoverflow", "angular2", "angular2-cli", "angular-cli"},
	{"stackoverflow", "angular2", "rxjs", "rxjs"},
	{"stackoverflow", "angular2", "rxjs5", "rxjs"},
	{"stackoverflow", "angular2", "ngrx", "ngrx"},
	{"stackoverflow", "angular2", "ngrx-effects", "ngrx"},
	{"stackoverflow", "angular2", "bootstrap-4", "bootstrap"},
	{"stackoverflow", "angular2", "material", "material"},
	{"stackoverflow", "angular2", "webpack", "webpack"},
	{"stackoverflow", "angular2", "webpack-2", "webpack"},
	{"stackoverflow", "angular2", "karma-jasmine", "testing"},
	{"stackoverflow", "angular2", "angular-material2", "material"},
	{"stackoverflow", "angular2", "angular-material", "material"},
	{"stackoverflow", "angular2", "angular2-routing", "routing"},
	{"stackoverflow", "angular2", "angular2-router", "routing"},
	{"stackoverflow", "angular2", "angular-ui-router", "routing"},
	{"stackoverflow", "angular2", "routing", "routing"},
	{"stackoverflow", "angular2", "promise", "promise"},
	{"stackoverflow", "angular2", "angular2-forms", "forms"},
	{"stackoverflow", "angular2", "angular2-template", "templates"},
	{"stackoverflow", "angular2", "animation", "animation"},
	{"stackoverflow", "angular2", "zone.js", "zone.js"},
	{"stackoverflow", "angular2", "zonejs", "zone.js"},
	{"stackoverflow", "angular2", "cors", "cors"},
	{"stackoverflow", "angular2", "angular4", "angular4"},
	{"stackoverflow", "go", "goroutine", "concurrency"},
	{"stackoverflow", "js", "ecmascript-6", "es6"},
	{"stackoverflow", "js", "google-chrome-extension", "google-chrome-extension"},
	{"stackoverflow", "js", "gulp", "gulp"},
	{"stackoverflow", "js", "eslint", "eslint"},
	{"stackoverflow", "js", "webpack", "webpack"},
	{"stackoverflow", "js", "meteor", "meteor"},
	{"stackoverflow", "js", "mocha", "mocha"},
	{"stackoverflow", "js", "chai", "chai"},
	{"stackoverflow", "js", "reactjs", "reactjs"},
	{"stackoverflow", "js", "redux", "reactjs"},
	{"stackoverflow", "js", "d3.js", "d3.js"},
	{"stackoverflow", "js", "momentjs", "momentjs"},
	{"stackoverflow", "js", "underscore.js", "underscore.js"},
	{"stackoverflow", "js", "webpack", "webpack"},
	{"stackoverflow", "js", "jquery", "jquery"},
	{"stackoverflow", "js", "json", "json"},
	{"stackoverflow", "web", "svg", "svg"},
	{"stackoverflow", "web", "twitter-bootstrap", "bootstrap"},
	{"stackoverflow", "node.js", "express", "express"},
	{"stackoverflow", "node.js", "socket.io", "socket.io"},
	{"stackoverflow", "node.js", "npm", "npm"},
	{"stackoverflow", "node.js", "callback", "callback"},
	{"stackoverflow", "node.js", "promise", "promise"},
	{"stackoverflow", "node.js", "stream", "stream"},
	{"stackoverflow", "node.js", "mongoose", "mongoose"},
	{"stackoverflow", ".net", ".net-core", "net.core"},
	{"stackoverflow", ".net", "asp.net-core", "net.core"},
	{"stackoverflow", ".net", "entity-framework-core", "net.core"},
	{"stackoverflow", ".net", "asp.net-web-api", "web.api"},
	{"stackoverflow", ".net", "asp.net-web-api2", "web.api"},
	{"stackoverflow", ".net", "wpf", "wpf"},
	{"stackoverflow", ".net", "wcf", "wcf"},
	{"stackoverflow", ".net", "linq", "linq"},
	{"stackoverflow", ".net", "entity-framework", "entity-framework"},
	{"stackoverflow", ".net", "entity-framework-6", "entity-framework"},
	{"stackoverflow", ".net", "ado.net", "ado.net"},
	{"stackoverflow", ".net", "uwp", "uwp"},
	{"stackoverflow", ".net", "asp.net", "asp.net"},
	{"stackoverflow", ".net", "asp.net-mvc", "asp.net"},
	{"stackoverflow", ".net", "asp.net-mvc-4", "asp.net"},
	{"stackoverflow", ".net", "signalr", "signalr"},
	{"stackoverflow", ".net", "odata", "odata"},
	{"stackoverflow", ".net", "async-await", "async"},
	{"stackoverflow", ".net", "asynchronous", "async"},
	{"stackoverflow", ".net", "multithreading", "async"},
	{"stackoverflow", ".net", "generics", "csharp-core"},
	{"stackoverflow", ".net", "extension-methods", "csharp-core"},
	{"stackoverflow", ".net", "reflection", "csharp-core"},
	{"stackoverflow", ".net", "idisposable", "csharp-core"},
	{"stackoverflow", ".net", "dictionary", "csharp-core"},
	{"stackoverflow", ".net", "enums", "csharp-core"},
	{"stackoverflow", ".net", "garbage-collection", "csharp-core"},
	{"stackoverflow", ".net", "collections", "csharp-core"},
	{"stackoverflow", ".net", "lambda", "csharp-core"},
	{"stackoverflow", ".net", "serialization", "csharp-core"},
	{"stackoverflow", ".net", "json", "json"},
	{"stackoverflow", ".net", "xml", "xml"},
	{"stackoverflow", ".net", "roslyn", "roslyn"},
	{"stackoverflow", ".net", "dependency-injection", "design"},
	{"stackoverflow", ".net", "design-patterns", "design"},
	{"stackoverflow", ".net", "oop", "design"},
	{"stackoverflow", ".net", "visual-studio", "vs"},
	{"stackoverflow", ".net", "visual-studio-2015", "vs"},
	{"stackoverflow", ".net", "visual-studio-2017", "vs"},
	{"stackoverflow", "java", "maven", "maven"},
	{"stackoverflow", "java", "gradle", "gradle"},
	{"stackoverflow", "java", "tomcat", "tomcat"},
	{"stackoverflow", "java", "spring", "spring"},
	{"stackoverflow", "java", "spring-boot", "spring"},
	{"stackoverflow", "java", "hibernate", "spring"},
	{"stackoverflow", "java", "reflection", "java-core"},
	{"stackoverflow", "java", "try-catch", "java-core"},
	{"stackoverflow", "java", "jsp", "jsp"},
	{"stackoverflow", "ios", "uitableview", "uitableview"},
	{"stackoverflow", "ios", "uicollectionview", "uicollectionview"},
	{"stackoverflow", "ios", "swift", "swift"},
	{"stackoverflow", "ios", "xcode", "xcode"},
	{"stackoverflow", "python", "pandas", "sci-fi"},
	{"stackoverflow", "python", "numpy", "sci-fi"},
	{"stackoverflow", "python", "matplotlib", "sci-fi"},
	{"stackoverflow", "bash", "awk", "awk"},
	{"stackoverflow", "bash", "sed", "sed"},
	{"stackoverflow", "bash", "cron", "cron"},
	{"stackoverflow", "bash", "curl", "curl"},
	{"stackoverflow", "bash", "scripting", "scripting"},
	{"stackoverflow", "git", "merge", "merge"},
	{"stackoverflow", "git", "git-merge", "merge"},
	{"stackoverflow", "git", "gitignore", "gitignore"},
	{"stackoverflow", "git", "ssh", "ssh"},
	{"stackoverflow", "git", "git-submodules", "submodules"},
	{"stackoverflow", "git", "rebase", "rebase"},
	{"stackoverflow", "git", "gitlab", "gitlab"},
	{"stackoverflow", "git", "github", "github"},
	{"stackoverflow", "docker", "docker-swarm", "swarm"},
	{"stackoverflow", "docker", "swarm", "swarm"},
	{"stackoverflow", "typescript", "types", "types"},
	{"stackoverflow", "bigdata", "elasticsearch", "elasticsearch"},
	{"stackoverflow", "bigdata", "apache-kafka", "kafka"},
	{"stackoverflow", "bigdata", "rabbitmq", "rabbitmq"},
	{"stackoverflow", "bigdata", "cassandra", "cassandra"},
	{"security", "information security", "certificates", "certificates"},
	{"security", "information security", "tls", "certificates"},
	{"security", "information security", "openssl", "certificates"},
	{"security", "information security", "public-key-infrastructure", "certificates"},
	{"security", "information security", "privacy", "privacy"},
	{"security", "information security", "torjj", "privacy"},
	{"security", "information security", "passwords", "passwords"},
	{"security", "information security", "android", "android"},
	{"security", "information security", "firewalls", "networking"},
	{"security", "information security", "network", "networking"},
	{"security", "information security", "openvpn", "networking"},
	{"security", "information security", "wireless", "networking"},
	{"security", "information security", "dmz", "networking"},
	{"security", "information security", "wpa2", "networking"},
	{"security", "information security", "dns", "networking"},
	{"security", "information security", "network-scanners", "networking"},
	{"security", "information security", "linux", "linux"},
	{"security", "information security", "windows", "windows"},
	{"codereview", "code review", "java", "java"},
	{"codereview", "code review", "python", "python"},
	{"codereview", "code review", "ruby", "ruby"},
	{"codereview", "code review", "go", "go"},
	{"codereview", "code review", "clojure", "clojure"},
	{"codereview", "code review", "c", "c"},
	{"codereview", "code review", "c#", "csharp"},
	{"codereview", "code review", "javascript", "js"},
}

// Check tag existence in the question
func containTag(q quzx_crawler.SOQuestion, tag string) (bool) {

	for _, question_tag := range q.Tags {
		if question_tag == tag {
			return true
		}
	}

	return false
}

// Check stop tag existance in the question
func containStopTag(q quzx_crawler.SOQuestion) (bool, string) {

	for _, stop_tag := range stop_tags {
		if containTag(q, stop_tag) {
			return true, stop_tag
		}
	}

	return false, ""
}

func firstLevelClassification(q quzx_crawler.SOQuestion, site string) (string) {

	for _, flr := range firstLevelRules {

		if site == flr.Site {

			if flr.Include == "*" {
				return flr.Result
			} else if containTag(q, flr.Include) {
				return flr.Result
			}
		}
	}

	return ""
}

func secondLevelClassification(q quzx_crawler.SOQuestion, site string, first string) (string) {

	for _, slr := range secondLevelRules {

		if (slr.Site == site)  && (slr.First == first) && containTag(q, slr.Include) {
			return slr.Result
		}
	}

	return "general"
}

func Classify(q quzx_crawler.SOQuestion, site string) (quzx_crawler.SOQuestion) {

	contain_stop_tag, stop_tag := containStopTag(q)
	if contain_stop_tag {
		fmt.Println("Classificator: stop tag - " + stop_tag)
		q.Classification = "remove"
		q.Details = "remove"
		return q
	}

	flr := firstLevelClassification(q, site)
	q.Classification = flr
	fmt.Println("Classificator: flr = " + flr)

	slr := secondLevelClassification(q, site, flr)
	q.Details = slr
	fmt.Println("Classificator: slr = " + slr)

	return q
}
