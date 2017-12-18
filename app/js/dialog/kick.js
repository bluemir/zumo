import $ from "/static/js/minilib.js";
import KV from "/static/js/kv.js";

var kickDialog = $.get("zumo-dialog.kick");

kickDialog.on("ok", async function(){
	if (!KV.channelID) {
		console.warn("[InviteDialog:_submit] cannot find target");
	}

	var username = $.get(kickDialog, "input[name=username]").value;

	await $.request("PUT", "/api/v1/channels/:channelID/kick/:username", {
		params: {
			channelID: KV.channelID,
			username: username,
		}
	});

	kickDialog.hide();
});
kickDialog.on("cancel", function(){
	kickDialog.hide();
});

export default {};
