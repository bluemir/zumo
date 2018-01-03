import $ from "/static/js/minilib.js";

import HookCreateDialog from "/static/js/dialog/hook-create.js";
import "/static/js/dialog/invite.js";
import "/static/js/dialog/kick.js";

var hookCreateDialog = new HookCreateDialog();

class ChannelMenu {
	constructor() {
		$.get(".channel.menu>button").on("click", this.toggle.bind(this));

		$.get(this.html, "button.invite").on("click", this._invite.bind(this));
		$.get(this.html, "button.leave").on("click", this._leave.bind(this));
		$.get(this.html, "button.kick").on("click", this._kick.bind(this));
		$.get(this.html, "button.hook-create").on("click", this._hookCreate.bind(this));
	}
	get html(){
		return $.get(".menu.channel");
	}
	show() {
		console.debug("[channel:menu:show]")
		$.get(this.html, "ul").classList.add("show");
	}
	hide() {
		$.get(this.html, "ul").classList.remove("show");
	}
	toggle () {
		console.debug("[channel:menu:toggle]")
		$.get(this.html, "ul").classList.toggle("show");
	}
	_invite() {
		this.hide();
		$.get("zumo-dialog.invite").show();
	}
	_leave() {

	}
	_hookCreate() {
		this.hide();
		hookCreateDialog.show();
	}
	_kick() {
		this.hide();
		//console.log($.get("zumo-dialog"))
		$.get("zumo-dialog.kick").show();
	}
}
export default ChannelMenu;
