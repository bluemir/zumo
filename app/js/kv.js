
//this one is singleton
var store = {};
var listeners = {};

const specialKeys = [
	"channelID",
]

class KV {
	constructor(){
		console.debug("[kv:init]");
	}
	set(key, value) {
		console.debug(`[kv:set] ${key} - ${value}`);
		store[key] = value;
		if (!listeners[key]){
			console.debug("[kv:set] listeners not found")
			return;
		}
		listeners[key].forEach(function(cb) {
			cb(key, value);
		});
	}
	get(key) {
		console.debug(`[kv:get] ${key}`);
		return store[key];
	}
	remove(key) {
		console.debug(`[kv:remove] ${key}`);
		if (specialKeys.contains(key)) {
			throw Error("cannot delete special key");
		}
		delete store[key];
	}
	watch(key, cb) {
		if (!listeners[key]){
			listeners[key] = [];
		}
		listeners[key].push(cb);
	}

	set channelID(v){
		this.set("channelID", v);
	}
	get channelID(){
		return this.get("channelID");
	}
}

export default new KV();
