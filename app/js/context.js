import MessageBox from "/static/js/message-box.js";


class Context {
	constructor() {
		this._logger = new Logger();
	}
	get log(){
		return this._logger;
	}
}

class Logger {
	constructor() {
		this._MessageBox = new MessageBox();
	}
	info(text) {
		this._MessageBox.info(text);
	}
	warn(text) {
		this._MessageBox.warn(text);
	}
	error(text) {
		this._MessageBox.error(text);
	}
}

class Dialog {
	constructor() {

	}
}

export default new Context();
