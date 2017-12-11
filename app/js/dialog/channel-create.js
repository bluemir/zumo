import $ from "/static/js/minilib.js";
import Dialog from "/static/js/dialog/dialog.js";

class ChannelCreateDialog extends Dialog{
	constructor(){
		super();
		console.debug("[dialog:channel-create:init]");

		$.get(this.html, "button.ok").on("click", this.submit.bind(this));
		$.get(this.html, "button.cancel").on("click", this.hide.bind(this));
	}
	get html() {
		return $.get(".dialog.channel-create")
	}
	clear() {
		$.get(this.html, "input[type=text]").value = "";
	}
	async submit(e) {
		var name = $.get(this.html, "input[type=text]").value;
		e.preventDefault();

		try {
			var channel = await $.request("POST", "/api/v1/channels", {
				 body: {
					 Name: name
				 }
			 });
			console.log("create channel", name)
			this.hide();
			this.clear();
		} catch(e) {
			// TODO show error message
			console.warn("error on create channel", name)
			this.hide();
		}
	}
}

export default ChannelCreateDialog;
