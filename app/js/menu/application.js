import $ from "/static/js/minilib.js";
import context from "/static/js/context.js";

class ApplicationMenu {
	constructor() {
		// open button
		$.get(".application.menu>button").on("click", this.toggle.bind(this))

		$.get(this.html, "button.bot-create").on("click", this._showBotCreateDialog.bind(this))
		$.get(this.html, "button.my-profile").on("click", this._showMyProfile.bind(this))
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
		//botCreateDialog.show();
		$.get("zumo-dialog.create-bot").show();
		this.hide();
	}
	_showMessageBox() {
		context.log.warn("test");
		this.hide();
	}
	async _showMyProfile() {
		// some ajax and some data
		var res = await $.request("GET", "/api/v1/users/me");
		console.log(res);
		$.get("zumo-dialog.my-profile [slot=body]").innerHTML = "<pre>" +JSON.stringify(res.json, null, 4) + "</pre>";
		$.get("zumo-dialog.my-profile").show();
		this.hide();
	}
}


$.get("zumo-dialog.my-profile").on("cancel", function(){
	this.hide();
});
$.get("zumo-dialog.my-profile").on("ok", function(){
	this.hide();
});

export default ApplicationMenu;
