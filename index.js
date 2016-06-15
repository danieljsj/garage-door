var firebase = require("firebase");
try {
	var gpio = require("rpi-gpio");
} catch (err) {
	console.error("FAILED TO REQUIRE 'rpi-gpio' PACKAGE; err:",err,"Are you sure you're on a raspberry pi???");
}

gpio.setup(11, gpio.DIR_OUT);


firebase.initializeApp({
  serviceAccount: "auth-config/pi-service-account-credentials-prod.json",
  databaseURL: "https://garage-67a27.firebaseio.com"
});

var triggerTimeoutId;
firebase.database().ref('triggerings').on('child_added', function(snapshot){
	
	//console.log('snapshot.val()',snapshot.val());
	
	clearTimeout(triggerTimeoutId);
	triggerTimeoutId = setTimeout(function(){
		trigger();
	},0);

});

var shutoffTimeoutId;
function trigger(){
	//console.log("TRIGGERING THE THING!");
	
	gpio.write(11,false,function(err){
		if (err) { console.log("\n\n\nDID YOU FORGET TO RUN AS SUDO???\n\n\n"); throw err; }
		//console.log("Turned 11 low, light on");
		clearTimeout(shutoffTimeoutId);
		shutoffTimeoutId = setTimeout(function(){
			shutoff();
		},1000);
	});
}

function shutoff(){
	//console.log("SHUTTING OFF THE LIGHT");
	gpio.write(11,true,function(err){
		if (err) throw err;
		//console.log("light should be off now");
	});
}



function closePins() {
    gpio.destroy(function() {
        console.log('All pins unexported');
    });
}




/*
process.stdin.resume(); //so the program will not close instantly

function exitHandler(options, err) {
	closePins();
    if (options.cleanup) console.log('clean');
    if (err) console.log(err.stack);
    if (options.exit) process.exit();
}

//do something when app is closing
process.on('exit', exitHandler.bind(null,{cleanup:true}));

//catches ctrl+c event
process.on('SIGINT', exitHandler.bind(null, {exit:true}));

//catches uncaught exceptions
process.on('uncaughtException', exitHandler.bind(null, {exit:true}));
*/
