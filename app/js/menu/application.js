import $ from "/static/js/minilib.js";

class ApplicationMenu {
	constructor() {

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
}

export default ApplicationMenu;
