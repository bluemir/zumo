import $ from "/static/js/minilib.js";

import KV from "/static/js/kv.js";
import PositionWatcher from "/static/js/modules/position-watcher.js";

import context from "/static/js/context.js";

const STATUS = "data-status";
const LOADING = "loading";
const READY = "ready";
const NONE = "";
const SECOND = 1000;

class InputBox {
	constructor(socket){
		this.html.on("submit", this._onSubmit.bind(this));
		// TODO make buffer class
		this._buffer = {}; // detail buffer
		this._location = new DetailLocation(this.html);
	}
	get html() {
		return $.get(".input-box");
	}
	get text() {
		return $.get(this.html, "input[type=text]").value;
	}
	set text(str) {
		$.get(this.html, "input[type=text]").value = str;
	}
	clear() {
		this.text = "";
		this._buffer = {};
		this._location.clear();
		$.get(this.html, ".status .location").attr(STATUS, NONE);
	}
	async _onSubmit(evt) {
		console.debug("[event:input-box:submit]");
		evt.preventDefault();

		if (!KV.channelID) {
			// TODO
			console.error("[InputBox:_onSubmit] cannot find channelID");
			context.log.error("must click channel first");
			return
		}
		if (this.text == "") {
			console.debug("[InputBox:_onSubmit] skip blank");
			return
		}

		this._buffer["zumo.message.location"] = await this._location.get();

		// TODO use await?
		console.log(this._buffer)
		$.request("POST", `/api/v1/channels/${KV.channelID}/messages`, {
			body: {
				Text: this.text,
				Detail: this._buffer,
			}
		});

		this.clear();
	}
	async addDetail(name, obj) {
		this._buffer[name] = obj;
		// TODO show detail icon or values on screen
	}
	async addLocation() {
		this._location.ready();
	}
}

class DetailLocation {
	constructor(parent) {
		this["--"] = {
			html: $.get(parent, ".status .location"),
		};
		this._ID = null;
		this._defer = null;
		this._timer = null;
	}
	get html(){
		return this["--"].html;
	}
	ready() {
		console.debug("[Inputbox:DetailLocation:ready]")
		if (!this._ID) {
			this._ID = navigator.geolocation.watchPosition(this._onData.bind(this));
		}
		this.html.attr(STATUS, LOADING);
	}
	async get() {
		console.debug("[Inputbox:DetailLocation:get]")
		if (!this._ID) {
			return; // no active;
		}
		if (this._value) {
			return this._value;
		}
		this._defer = $.defer();
		return this._defer.promise;
	}
	clear() {
		console.debug("[Inputbox:DetailLocation:clear]")
		if (this._ID) {
			navigator.geolocation.clearWatch(this._ID)
			this._ID = null;
		}

		this._value = null;

		if (this._defer) {
			this._defer.reject("unexpected clear")
		}
		this.html.attr(STATUS, NONE);

		clearTimeout(this._timer);
		this._timer = null;
	}

	_onData(pos) {
		console.debug("[Inputbox:DetailLocation:_onData]")
		this._value = {
			coords: {
				accuracy: pos.coords.accuracy,
				latitude: pos.coords.latitude,
				longitude: pos.coords.longitude,
			},
			timestamp: pos.timestamp,
		};
		if (this._defer) {
			this._defer.resolve(this._value);
			this._defer = null;
		}
		this.html.attr(STATUS, READY);

		// set invalid Timer
		clearTimeout(this._timer);
		this._timer = setTimeout(this._onInvaild.bind(this), 30 * SECOND);
	}

	_onInvaild() {
		console.debug("[Inputbox:DetailLocation:_onInvaild]")

		if (this._ID) {
			navigator.geolocation.clearWatch(this._ID)
			this._ID = null;
		}
		this._value = null;

		// TODO sprate invaild status
		this.html.attr(STATUS, NONE);

		this._timer = null;
	}
}

export default InputBox;
