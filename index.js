var firebase = require("firebase");

firebase.initializeApp({
  serviceAccount: "auth-config/pi-service-account-credentials-prod.json",
  databaseURL: "https://garage-67a27.firebaseio.com"
});

firebase.database().ref().on('value', function(snapshot){
	console.log(snapshot.val());
});