import $ from "/static/js/minilib.js";

class Dialog {
	constructor(){
		//$.get(this.html, "button.ok").on("click", this._submit.bind(this));
		//$.get(this.html, "button.cancel").on("click", this._cancel.bind(this));
	}
	show() {
		this.html.classList.add("show");
	}
	hide() {
		this.html.classList.remove("show");
	}
	toggle() {
		this.html.classList.toggle("show");
	}
	_submit(evt){
		console.debug("Dialog:_submit");
		evt.preventDefault();
	}
	_cancel(evt){
		console.debug("Dialog:_cancel");
		evt.preventDefault();
	}
}
export default Dialog
