import $ from "/static/js/minilib.js";

class Dialog {
	constructor() {
		console.debug("[dialog:bot-create:constructor]")
		$.get(this.html, "form").on("submit", this._submit.bind(this));
		$.get(this.html, "button.cancel").on("click", this._cancel.bind(this));
	}
	get html() {
		return $.get(".dialog.bot-create")
	}
	show(){
		this.html.classList.add("show")
	}
	hide() {
		this.html.classList.remove("show")
	}
	async _submit(e) {
		e.preventDefault()
		var name = $.get(this.html, "input[name=name]").value.trim();
		var driver = $.get(this.html, "input[name=driver]").value.trim();


		if (name == "") {
			console.error("[dialog:bot-create:_submit] name is blank")
			return
		}

		if (driver == "") {
			console.error("[dialog:bot-create:_submit] driver is blank")
			return
		}

		try {
			await $.request("POST", "/api/v1/bots", {
				body: {
					Name: name,
					Driver: driver,
				}
			});
			this.hide()
		} catch(e) {
			console.error("[dialog:bot-create:_submit]", e)
			this.hide()
		}
	}
	_cancel() {
		this.hide();
	}
}

export default Dialog;
