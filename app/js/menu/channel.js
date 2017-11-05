import $ from "/static/js/minilib.js";

class ChannelMenu {
	get html(){
		return $.get("menu.channel")
	}
	show() {
		console.debug("[channel:menu:show]")
		this.html.classList.add("show");
	}
	hide() {
		this.html.classList.remove("show");
	}
	toggle () {
		this.html.classList.toggle("show");
	}
}
export default ChannelMenu;
