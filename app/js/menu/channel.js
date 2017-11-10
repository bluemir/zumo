import $ from "/static/js/minilib.js";
import InviteDialog from "/static/js/dialog/invite.js";
import HookCreateDialog from "/static/js/dialog/hook-create.js";

var inviteDialog = new InviteDialog();
var hookCreateDialog = new HookCreateDialog();

class ChannelMenu {
	constructor() {
		$.get(".channel.menu>button").on("click", this.toggle.bind(this));

		$.get(this.html, "button.invite").on("click", this._invite.bind(this));
		$.get(this.html, "button.leave").on("click", this._leave.bind(this));
		$.get(this.html, "button.hook-create").on("click", this._hookCreate.bind(this));
	}
	get html(){
		return $.get(".menu.channel")
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
		inviteDialog.show();
	}
	_leave() {

	}
	_hookCreate() {
		this.hide();
		hookCreateDialog.show();
	}

}
export default ChannelMenu;
