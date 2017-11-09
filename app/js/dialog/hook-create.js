import $ from "/static/js/minilib.js";
import Dialog from "/static/js/dialog/dialog.js";
import KV from "/static/js/kv.js";

class HookCreateDialog extends Dialog {
	constructor() {
		super();
		$.get(this.html, "button.ok").on("click", this._submit.bind(this));
		$.get(this.html, "button.cancel").on("click", this._cancel.bind(this));
	}
	get html() {
		return $.get(".dialog.hook-create");
	}
	async _submit(evt) {
		evt.preventDefault();
		var username = $.get(this.html, "input[name=username]").value

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
	}

	_cancel(evt) {
		evt.preventDefault();
		this.hide();
	}
}

export default HookCreateDialog;
