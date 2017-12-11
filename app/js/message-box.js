import $ from "/static/js/minilib.js";

class MessageBox {
	constructor() {

	}
	get html() {
		return $.get("article.message-box");
	}
	info(text) {
		console.debug("called info")
		this.makeBox("info", text);
	}
	warn(text) {
		this.makeBox("warn", text);
	}
	makeBox(level, text) {
		var box = $.create("div", {
			class: level,
			$text: text,
		});

		box.on("animationstart", function() {

		})
		box.on("animationend", function(){
			box.remove();
		})

		this.html.appendChild(box);
		// TODO fade out
	}
}

export default MessageBox;
