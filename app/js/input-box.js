import $ from "/static/js/minilib.js";
import KV from "/static/js/kv.js";

class InputBox {
	constructor(socket){
		this.html.on("submit", this._onSubmit.bind(this));
	}
	get html() {
		return $.get(".input-box");
	}
	get text() {
		return $.get(this.html, "input[type=text]").value;
	}
	set text(str) {
		$.get(this.html, "input[type=text]").value = str;
	}
	clear() {
		this.text = "";
	}
	_onSubmit(evt) {
		console.debug("[event:input-box:submit]");
		evt.preventDefault();

		if (!KV.channelID) {
			// TODO
			console.error("[InputBox:_onSubmit] cannot find channelID");
			return
		}
		if (this.text == "") {
			console.debug("[InputBox:_onSubmit] skip blank");
			return
		}

		$.request("POST", `/api/v1/channels/${KV.channelID}/messages`, {
			body: {
				Text: this.text,
			}
		})

		this.clear();
	}
}

export default InputBox;
