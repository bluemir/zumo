import $ from "/static/js/minilib.js";
import ChannelList from "/static/js/channels.js";
import InputBox from "/static/js/input-box.js";
import Messages from "/static/js/messages.js";
import context from "/static/js/context.js";


import ReconnetSocket from "/static/js/reconnect-socket.js";

import KV from "/static/js/kv.js";


var socket = new ReconnetSocket(ReconnetSocket.protocal()+ "//" + location.host + "/ws");
socket.on("open", function(evt){
	console.debug("[event:websocket:open]");
	context.log.info("conneted!")
});
socket.on("close", function(evt){
	console.debug("[event:websocket:close]");
	context.log.warn("disconnected!");
	// close
});
socket.on("message", function(evt){
	console.debug("[event:websocket:message]");
	var packet = JSON.parse(evt.data)
	switch (packet.Type) {
		case "event":
			channelList.refresh();
			break;
		case "message":
			// Apppend to Message Panel
			messages.appendMessage(packet)
			break;
		default:
			console.debug("[event:websocket:message] packet:", packet);
	}
	// append to message panel
	// messages[channel.id].append(message)
});
setInterval(function() {
	socket.send('{"ping":"ping"}');
}, 60 * 1000);

var inputbox = new InputBox();
var channelList = new ChannelList();
var messages = new Messages();

channelList.onTargetChange(function(channel){
	KV.channelID = channel.ID;
	$.get(".channel-header h2").innerHTML = channel.Name;
})

// init menus
import InputMenu from "/static/js/menu/input.js"
//import ApplicationMenu from "/static/js/menu/application.js";

var inputMenu           = new InputMenu(inputbox);
//var applicationMenu     = new ApplicationMenu();

$.get("zumo-menu.channel").on("menu", function(e) {
	switch(e.detail.name) {
		case "invite":
			$.get("zumo-dialog.invite").show();
			this.hide();
			break;
		case "kick":
			$.get("zumo-dialog.kick").show();
			this.hide();
			break;
		case "leave":
			break;
		case "hook-create":
			$.get("zumo-dialog.create-hook").show()
			this.hide();
			break;
	}
});
$.get("zumo-menu.application").on("menu", async function(e) {
	switch(e.detail.name) {
		case "my-profile":
			var res = await $.request("GET", "/api/v1/users/me");
			$.get("zumo-dialog.my-profile [slot=body]").innerHTML = "<pre>" +JSON.stringify(res.json, null, 4) + "</pre>";
			$.get("zumo-dialog.my-profile").show();
			this.hide();
			break;
		case "bot-create":
			$.get("zumo-dialog.create-bot").show();
			this.hide();
			break;
		case "register-token":
			this.hide();
			break;
	}
});
$.get("zumo-dialog.my-profile").on("ok", function(){
	this.hide();
}).on("cancel", function() {
	this.hide();
});


// create channel dialog
$.get("zumo-dialog.create-channel").on("ok", async function(evt){
	var name = $.get(this, "input[type=text]").value.trim();
	try {
		var channel = await $.request("POST", "/api/v1/channels", {
			body: {
				Name: name
			}
		});
		console.log("create channel", name)
		context.log.info(`${name} channel created!`);
		this.hide();
		this.clear();
	} catch(e) {
		// TODO show error message
		console.warn("error on create channel", name)
		context.log.error("error on create channel!");
		this.hide();
	}
}).on("cancel", function(){
	$.get(this, "input[type=text]").value = "";
	this.hide();
});

// create bot dialog
$.get("zumo-dialog.create-bot").on("ok", async function(evt) {
	var name = $.get(this, "input[name=name]").value.trim();
	var driver = $.get(this, "input[name=driver]").value.trim();


	if (name == "") {
		context.log.error("name is blank");
		console.error("[dialog:bot-create:_submit] name is blank")
		return
	}

	if (driver == "") {
		context.log.error("driver is blank");
		console.error("[dialog:bot-create:_submit] driver is blank")
		return
	}

	try {
		await $.request("POST", "/api/v1/bots", {
			body: {
				Name: name,
				Driver: driver,
			}
		});
		this.hide()
		context.log.info("bot created");
	} catch(e) {
		console.error("[dialog:bot-create:_submit]", e)
		this.hide()
	}
}).on("cancel", function(evt) {
	this.hide();
});

$.get("button.channel-join").on("click", async function() {
	var dialog = $.get("zumo-dialog.join-channel")

	var res = await $.request("GET", "/api/v1/channels");
	$.get(dialog, "select").clear();
	res.json.map(function(e) {
		return $.create("option", {
			$text: `${e.Name}`,
			"value": e.ID
		});
	}).forEach(function(e){
		$.get(dialog, "select").appendChild(e);
	}, this);

	$.get("zumo-dialog.join-channel").show()
});

$.get("button.channel-create").on("click", async function() {
	$.get("zumo-dialog.create-channel").show()
});

// join channel dialog
$.get("zumo-dialog.join-channel").on("ok", async function(evt) {
	evt.preventDefault();
	var channelID = $.get(this, "select").value

	await $.request("PUT", `/api/v1/channels/${channelID}/join`, {})

	this.hide();
}).on("cancel", function(){
	this.hide();
})

// kick dialog
$.get("zumo-dialog.kick").on("ok", async function(){
	if (!KV.channelID) {
		console.warn("[kickDialog:_submit] cannot find target");
	}

	var username = $.get(kickDialog, "input[name=username]").value;

	try {
		var res = await $.request("PUT", "/api/v1/channels/:channelID/kick/:username", {
			params: {
				channelID: KV.channelID,
				username: username,
			}
		});
	} catch(e) {
		console.warn(e)
		context.log.warn(e.text);
	}

	this.hide();
}).on("cancel", function(){
	this.hide();
});

// create hook dialog
$.get("zumo-dialog.create-hook").on("ok", async function(evt) {
	evt.preventDefault();
	var username = $.get(this, "input[name=username]").value

	if (!KV.channelID) {
		console.warn("[HookCreateDialog:_submit] cannot find target");
	}

	var res = await $.request("POST", "/api/v1/hooks", {
		body: {
			ChannelID: KV.channelID,
			Username:  username,
		}
	});

	// TODO show ID
	console.log(`Hook ID: '${res.json.ID}'`)
	this.hide();
}).on("cancel", function () {
	this.hide();
})
