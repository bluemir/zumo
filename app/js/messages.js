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
			$html: `${message.Sender} - ${message.Text}`
		});

		elem.data = message

		if (message.Detail && message.Detail["zumo.message.detail.html"]) {
			// TODO add html element for custom element
			elem.appendChild($.create("div", {
				$html: message.Detail["zumo.message.detail.html"],
			}));
		}

		return elem
	}
}

export default Messages;
