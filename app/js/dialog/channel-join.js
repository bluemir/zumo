import $ from "/static/js/minilib.js";

class ChannelJoinDialog {
	constructor(){
		$.get(this.html, "button.ok").on("click", this._submit.bind(this));
		$.get(this.html, "button.cancel").on("click", this._cancel.bind(this));
	}
	get html() {
		return $.get(".dialog.channel-join")
	}
	show(){
		this.html.classList.add("show")
		this._load();
	}
	hide() {
		this.html.classList.remove("show")
	}
	async _load() {
		var res = await $.request("GET", "/api/v1/channels")
		this.clear();
		res.json.map(function(e) {
			return $.create("option", {
				$text: `${e.Name}`,
				"value": e.ID
			});
		}).forEach(function(e){
			$.get(this.html, "select").appendChild(e);
		}, this);
	}
	async _submit(evt) {
		evt.preventDefault();
		var channelID = $.get(this.html, "select").value

		await $.request("PUT", `/api/v1/channels/${channelID}/join`, {})

		this.hide();
	}
	_cancel(evt) {
		evt.preventDefault();
		this.hide();
	}
	clear() {
		$.get(this.html, "select").clear();
	}
}

export default ChannelJoinDialog;
