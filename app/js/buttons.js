import $ from "/static/js/minilib.js";

import ChannelCreateDialog from "/static/js/dialog/channel-create.js";
import ChannelJoinDialog from "/static/js/dialog/channel-join.js";

function init() {
	var channelCreateDialog = new ChannelCreateDialog();
	var channelJoinDialog   = new ChannelJoinDialog();

	button("button.channel-create",    e => channelCreateDialog.show())
	button("button.channel-join",      e => channelJoinDialog.show())
	button("button.logout",            e => logout())
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
