import $ from "/static/js/minilib.js";

import BotCreateDialog from "/static/js/dialog/bot-create.js";
import ChannelMenu from "/static/js/menu/channel.js"
import InputMenu from "/static/js/menu/input.js"
import ApplicationMenu from "/static/js/menu/application.js";
import ChannelCreateDialog from "/static/js/dialog/channel-create.js";
import ChannelJoinDialog from "/static/js/dialog/channel-join.js";

function init() {
	var channelCreateDialog = ChannelCreateDialog;
	var channelJoinDialog   = new ChannelJoinDialog();
	var botCreateDialog     = new BotCreateDialog();
	var channelMenu         = new ChannelMenu();
	var inputMenu           = new InputMenu();
	var applicationMenu     = new ApplicationMenu();

	button("button.channel-create",    e => channelCreateDialog.show())
	button("button.channel-join",      e => channelJoinDialog.show())
	button("button.logout",            e => logout())
	button("button.bot-create",        e => botCreateDialog.show())
	button("button.channel-menu",      e => channelMenu.toggle())
	button(".input.menu button",       e => inputMenu.toggle())
	button(".application.menu button", e => applicationMenu.toggle())
}
function button(name, fn) {
	$.get(name).on("click", function(evt){
		console.debug(`[event:${name}:click]`);
		fn(evt);
	});
}

async function logout() {
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
}


export default init;
