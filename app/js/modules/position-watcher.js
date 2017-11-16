import $ from "/static/js/minilib.js";

const second = 1000;

class PositionWatcher {
	constructor(){
		this._ID = null; // empty
		this._defer = null; // Promise value
		this._timer = null;
	}
	// starting watch
	async start(){
		console.debug("[PositionWatcher:start]")
		this._ID = navigator.geolocation.watchPosition(this._onUpdate.bind(this));
		this._defer = $.defer();
	}
	// end and clear
	end() {
		console.debug("[PositionWatcher:end]")
		navigator.geolocation.clearWatch(this._ID);
		this._ID = null;
		this._defer.reject("unexpected end");
	}
	// reaturn promise when not exist value;
	async get() {
		console.debug("[PositionWatcher:get]", this)

		if (!this._ID) {
			throw new Error("MUST start watcher before get")
		}
		return this._defer.promise;
	}

	get isStarted (){
		return this._ID != null;
	}

	_onUpdate(pos) {
		console.debug("[PositionWatcher:_onUpdate]", pos)

		var data = {
			coords: {
				accuracy: pos.coords.accuracy,
				latitude: pos.coords.latitude,
				longitude: pos.coords.longitude,
			},
			timestamp: pos.timestamp,
		}

		if (this._defer == null) {
			this._defer = $.defer();
		}
		this._defer.resolve(data);

		if (this._timer) {
			clearTimeout(this._timer);
		}

		this._timer = setTimeout(this._onInvaild.bind(this), 30 * second);
	}
	_onInvaild() {
		console.debug("[PositionWatcher:_onInvaild]")
		this._timer = null;
		this._defer.reject("Timeout")
		this._defer = null;
		// TODO restart watch?
	}
}



export default PositionWatcher;
