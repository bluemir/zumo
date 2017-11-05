import $ from "/static/js/minilib.js";
import Dialog from "/static/js/dialog/dialog.js";
import KV from "/static/js/kv.js";

class InviteDialog extends Dialog {
	constructor(){
		super();
	}
	get html(){
		return $.get(".dialog.invite");
	}
	async _submit(evt){
		console.debug("[InviteDialog:_submit]");
		evt.preventDefault();
		if (!KV.channelID) {
			console.warn("[InviteDialog:_submit] cannot find target");
		}

		var name = $.get(this.html, "input[name=username]").value;

		await $.request("PUT", "/api/v1/channels/:channelID/invite/:username", {
			params: {
				channelID: KV.channelID,
				username: name,
			}
		})
	}
	_cancel(evt){
		console.debug("[InviteDialog:_cancel]");
		evt.preventDefault();
		this.hide();
	}
}

export default InviteDialog;
