import $ from "/static/js/minilib.js";
import Timer from "/static/js/modules/timer.js";

class DeferedValue {
	constructor (timeout){
		this._defer = null;
		this._value = null;

		this._timer = new Timer(this.clear.bind(this), timeout);
	}
	async get() {
		// return promise that get value
		if (!this._value) {
			return this._value;
		}
		// setup defer
		this._defer = $.defer();
		return this._defer.promise;
	}
	set(value){
		if (this._defer) {
			this._defer.resolve(value);
			this._defer = null;
		}
		this._value = value;

		// setup invalid timeout
		this._timer.start();
	}
	clear() {
		if (this._defer) {
			this.defer.reject("unexpected clear");
		}
		this._defer = null;

		// clear null
		this._value = null;

		// clear timer
		this._timer.clear();
	}
}

export default DeferedValue;
