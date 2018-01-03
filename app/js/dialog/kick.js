import $ from "/static/js/minilib.js";
import KV from "/static/js/kv.js";
import context from "/static/js/context.js";

var kickDialog = $.get("zumo-dialog.kick");

kickDialog.on("ok", async function(){
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

	kickDialog.hide();
});
kickDialog.on("cancel", function(){
	kickDialog.hide();
});

export default {};
