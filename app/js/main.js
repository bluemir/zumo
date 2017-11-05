import $ from "/static/js/minilib.js";
import ChannelList from "/static/js/channels.js";
import InputBox from "/static/js/input-box.js";
import Messages from "/static/js/messages.js";
import BotCreateDialog from "/static/js/dialog/bot-create.js";
import InviteDialog from "/static/js/dialog/invite.js";
import ReconnetSocket from "/static/js/reconnect-socket.js";
import ChannelMenu from "/static/js/menu/channel.js"
import ChannelCreateDialog from "/static/js/dialog/channel-create.js";
import ChannelJoinDialog from "/static/js/dialog/channel-join.js";
import KV from "/static/js/kv.js";

var socket = new ReconnetSocket("ws://" + location.host + "/ws");
socket.on("open", function(evt){
	console.debug("[event:websocket:open]");
	//
});
socket.on("close", function(evt){
	console.debug("[event:websocket:close]");
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
var inviteDialog = new InviteDialog();

channelList.onTargetChange(function(channel){
	KV.channelID = channel.ID;
	$.get(".channel-header").innerHTML = channel.Name;
})

var botCreateDialog = new BotCreateDialog();

$.get("button.bot-create").on("click", function(){
	botCreateDialog.show();
});

var channelMenu = new ChannelMenu();
$.get("button.channel-menu").on("click", function() {
	channelMenu.toggle();
});

$.get("button.logout").on("click", async function(){
	try {
		await $.request("GET", "/api/v1/channels", {
			$auth:{
				user: "__",
				password: "__",
			}
		});
	} catch(e) {
		console.log("help")
	}
	document.location = "/"
})

$.get("button.invite").on("click", function(){
	inviteDialog.show();
})

var channelCreateDialog = ChannelCreateDialog;
$.get(".channels button.channel-create").on("click", function() {
	console.debug("[event:button.channel-create:click]");
	channelCreateDialog.show();
})

var channelJoinDialog = new ChannelJoinDialog();
$.get(".channels button.channel-join").on("click", function(){
	console.debug("[event:button.channel-join:click]");
	channelJoinDialog.show();
});
