import $ from "/static/js/minilib.js";
import ChannelList from "/static/js/channels.js";
import InputBox from "/static/js/input-box.js";
import Messages from "/static/js/messages.js";
import context from "/static/js/context.js";


import ReconnetSocket from "/static/js/reconnect-socket.js";

import KV from "/static/js/kv.js";
import Buttons from "/static/js/buttons.js";


var socket = new ReconnetSocket(ReconnetSocket.protocal()+ "//" + location.host + "/ws");
socket.on("open", function(evt){
	console.debug("[event:websocket:open]");
	context.log.info("conneted!")
	//
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
			// TODO Apppend to Message Panel
			messages.appendMessage(packet)
			break;
		default:
			console.debug("[event:websocket:message] packet:", packet);
	}
	// append to message panel
	// messages[channel.id].append(message)
});

var inputbox = new InputBox();
var channelList = new ChannelList();
var messages = new Messages();

channelList.onTargetChange(function(channel){
	KV.channelID = channel.ID;
	$.get(".channel-header h2").innerHTML = channel.Name;
})

Buttons(inputbox);

// init menus
import ChannelMenu from "/static/js/menu/channel.js"
import InputMenu from "/static/js/menu/input.js"
import ApplicationMenu from "/static/js/menu/application.js";

var channelMenu         = new ChannelMenu();
var inputMenu           = new InputMenu(inputbox);
var applicationMenu     = new ApplicationMenu();


$.get("zumo-menu").on("menu", function(e) {
	switch(e.detail.name) {
		case "invite":
			$.get("zumo-dialog.invite").show();
			break;
		case "kick":
			$.get("zumo-dialog.kick").show();
			this.toggle();
			break;
		case "leave":
			break;
		case "hook-create":
			// TODO popup hook create
			break;
	}
})

// create channel-dialog
$.get("zumo-dialog.create-channel").on("ok", async function(evt){
	var name = $.get(this, "input[type=text]").value;
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
})

