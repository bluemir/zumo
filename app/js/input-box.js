import $ from "/static/js/minilib.js";
import KV from "/static/js/kv.js";

class InputBox {
	constructor(socket){
		this.html.on("submit", this._onSubmit.bind(this));
		// TODO make buffer class
		this._buffer = {}; // detail buffer
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
		this._buffer = {};
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

		// TODO use await?
		$.request("POST", `/api/v1/channels/${KV.channelID}/messages`, {
			body: {
				Text: this.text,
				Detail: this._buffer,
			}
		});

		this.clear();
	}
	async addDetail(name, obj) {
		this._buffer[name] = obj;
		// TODO show detail icon or values on screen
	}
	// TODO disable submit (when get location ...)
}


export default InputBox;
