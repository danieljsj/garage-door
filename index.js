var firebase = require("firebase");
try {
	var gpio = require("pi-gpio");
} catch (err) {
	console.error("FAILED TO REQUIRE 'pi-gpio' PACKAGE; err:",err,"Are you sure you're on a raspberry pi???");
}

firebase.initializeApp({
  serviceAccount: "auth-config/pi-service-account-credentials-prod.json",
  databaseURL: "https://garage-67a27.firebaseio.com"
});

var triggerTimeoutId;
firebase.database().ref('triggerings').on('child_added', function(snapshot){
	
	console.log('snapshot.val()',snapshot.val());
	
	clearTimeout(triggerTimeoutId);
	triggerTimeoutId = setTimeout(function(){
		trigger();
	},0);

});

function trigger(){
	console.log("TRIGGERING THE THING!");
}
