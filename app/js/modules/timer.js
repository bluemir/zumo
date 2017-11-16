import $ from "/static/js/minilib.js";

class Timer {
	constructor(cb, amount) {
		this.amount = amount;
		this._ID = null;
		this._cb = cb;
	}
	start() {
		if ((!this.amount) || (this.amount < 0 )) {
			return;
		}
		clearTimeout(this._ID);
		this._ID = setTimeout(this._onTime, this.amount);
	}
	clear() {
		clearTimeout(this._ID);
		this._ID = null;
	}
	_onTime() {
		this._ID = null;
		this._cb();
	}
}

export default Timer;
