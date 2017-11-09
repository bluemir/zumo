// TODO insert geolocation?
// TODO optional module?
// TODO detail?
import $ from "/static/js/minilib.js";

class InputMenu{
	constructor(inputbox) {
		this._inputbox = inputbox;

		//$.get(this.html, "put geo location", this._putGeoToMessage.bind(this))
	}
	get html(){
		return $.get(".input.menu")
	}
	show() {
		$.get(this.html, "ul").classList.add("show");
	}
	hide() {
		$.get(this.html, "ul").classList.remove("show");
	}
	toggle() {
		$.get(this.html, "ul").classList.toggle("show");
	}
	_putGeoToMessage(){
		this._inputbox.addDetail("zumo.message.geo", {"geo":"temp"}); // invalidate by inputbox
	}
	_putLinkToMessage() {
		this._inputbox.addDetail("zumo.message.link", {"data": ""});
	}
}
export default InputMenu;
