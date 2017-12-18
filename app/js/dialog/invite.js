import $ from "/static/js/minilib.js";
import KV from "/static/js/kv.js";

$.get("zumo-dialog.invite").on("ok", async function(){
	if (!KV.channelID) {
		console.warn("[InviteDialog:_submit] cannot find target");
	}

	var name = $.get($.get("zumo-dialog.invite"), "input[name=username]").value;

	await $.request("PUT", "/api/v1/channels/:channelID/invite/:username", {
		params: {
			channelID: KV.channelID,
			username: name,
		}
	})
	this.hide();
});
$.get("zumo-dialog.invite").on("cancel", function(){
	this.hide();
});

export default {};
