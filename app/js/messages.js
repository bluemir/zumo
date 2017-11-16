import $ from "/static/js/minilib.js";
import KV from "/static/js/kv.js";

class Messages {
	constructor () {
		KV.watch("channelID", this._loadMessage.bind(this));
	}
	get html() {
		return $.get(".messages")
	}
	appendMessage(message){
		if (!KV.channelID) {
			console.warn("[messages:appendMessage] cannot find channelID");
		}

		if (message.ChannelID != KV.channelID) {
			console.debug("[messages:appendMessage] skip not includes messages")
			return // skip
		}

		this.html.appendChild(this._createMessageElement(message));

		this.scrollToEnd();
	}

	async _loadMessage() {
		var res = await $.request("GET", `/api/v1/channels/${KV.channelID}/messages`);

		this.html.clear();
		res.json.reverse().forEach(function(message){
			this.html.appendChild(this._createMessageElement(message));
		}, this)
		this.scrollToEnd();
	}
	scrollToEnd(){
		// maybe work...
		this.html.scrollTop = this.html.scrollHeight -  this.html.clientHeight;
	}

	_createMessageElement(message) {
		var elem = $.create("li", {
			class: "message",
			$html: `<a class="author">${message.Sender}</a><p class="text">${message.Text}</p>`,
		});

		elem.data = message;

		if (!message.Detail) {
			return elem;
		}

		if (message.Detail["zumo.message.html"]) {
			// replace text to html
			console.log($.get(elem, ".text"))
			$.get(elem, ".text").innerHTML = message.Detail["zumo.message.html"];
		}

		return elem
	}
}

export default Messages;
