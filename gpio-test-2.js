
var gpio = require('rpi-gpio');

/* gpio.on('export', function(channel) {
    console.log('Channel set: ' + channel);
}); */
  
gpio.setup(11, gpio.DIR_OUT, write);

function write() {
	
	gpio.write(11, false, function(err) { // false means low, which is what we need to create flow from our 3.3V to this pin. Not sure why we did that instead of connecting to ground and then going high to trigger the LED... but that's okay. Either way is fine.
		if (err) throw err;
		console.log("WROTE!");
		setTimeout(closePins, 2000);
	});

}
 
function closePins() {
    gpio.destroy(function() {
        console.log('All pins unexported');
    });
}

return;
