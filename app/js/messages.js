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

		var elem = $.create("li", {
			$html: `${message.Sender} - ${message.Text}`
		})


		elem.data = message

		this.html.appendChild(elem)
		this.scrollToEnd()
	}

	async _loadMessage() {
		var res = await $.request("GET", `/api/v1/channels/${KV.channelID}/messages`);

		this.html.clear();
		res.json.reverse().forEach(function(message){
			var elem = $.create("li", {
				$html: `${message.Sender} - ${message.Text}`
			});

			elem.data = message

			this.html.appendChild(elem);
			this.scrollToEnd();
		}, this)
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

		if (message.Labels["zumo.detail.html"]) {
			// TODO add html element for custom element
		}



		return elem
	}
}

export default Messages;
