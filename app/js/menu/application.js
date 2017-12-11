import $ from "/static/js/minilib.js";
import context from "/static/js/context.js";

import BotCreateDialog from "/static/js/dialog/bot-create.js";
var botCreateDialog = new BotCreateDialog();

class ApplicationMenu {
	constructor() {
		// open button
		$.get(".application.menu>button").on("click", this.toggle.bind(this))

		$.get(this.html, "button.bot-create").on("click", this._showBotCreateDialog.bind(this))
		$.get(this.html, "button.testbtn").on("click", this._showMessageBox.bind(this))
	}
	get html(){
		return $.get(".application.menu");
	}
	show(){
		$.get(this.html, "ul").classList.add("show");
	}
	hide(){
		$.get(this.html, "ul").classList.remove("show");
	}
	toggle(){
		$.get(this.html, "ul").classList.toggle("show");
	}
	_showBotCreateDialog() {
		botCreateDialog.show();
		this.hide();
	}
	_showMessageBox() {
		context.log.warn("test");
		this.hide();
	}
}

export default ApplicationMenu;
