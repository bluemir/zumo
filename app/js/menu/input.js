import $ from "/static/js/minilib.js";

class InputMenu{
	constructor(inputbox) {
		this._inputbox = inputbox;
		$.get(this.html, ".input.menu>button").on("click", this.toggle.bind(this));

		$.get(this.html, "button.location").on("click", this._putLocationToMessage.bind(this));
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
	async _putLocationToMessage(){
		// TODO maybe move this method to input box this menu only trigger or show menus...
		console.debug("[menu:input:_putLocationToMessage]")
		this.hide();
		this._inputbox.addLocation();
	}
	_putLinkToMessage() {
		this._inputbox.addDetail("zumo.message.link", {"data": ""});
	}
}
export default InputMenu;
