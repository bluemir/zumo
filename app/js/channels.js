import $ from "/static/js/minilib.js";


class ChannelList {
	constructor() {
		console.debug("[channel-list:init]")

		this.refresh();
		this.html.on("click", this._onChannelChange.bind(this));

		var that = this;

		this._listeners = [];
	}
	get html(){
		return $.get(".channels ul");
	}
	async refresh() {
		console.debug("[channel-list:refresh]");
		var res = await $.request("GET", "/api/v1/users/me/joinned-channel");
		var channels = res.json.sort(function(a, b){
			if (a.Name === b.Name) {
				return 0;
			}
			return a.Name > b.Name ? 1 : -1;
		});

		this.html.clear();

		var that = this;
		channels.map(function(e){
			var elem = $.create("li", {
				$html: `<span>${e.Name}</span><button class="delete">X</button>`
			});
			elem.data = e;

			return elem;
		}).forEach(function(e) {
			that.html.appendChild(e);
		});
	}
	_onChannelChange(evt) {
		console.debug("[event:channel-channge:click]");

		if (evt.target.tagName == "UL") {
			return; // skip ul
		}

		var target = evt.target.parentElement;
		if (evt.target.classList.contains("delete")) { // if delete button
			this._deleteChannel(target.data.ID);
			evt.stopPropagation();
			return
		}

		this.html.childNodes.forEach(e => e.classList.remove("select"))
		target.classList.add("select");

		this._listeners.forEach(l => l(target.data))
	}
	onTargetChange(l) {
		this._listeners.push(l);
	}
	async _deleteChannel(id){
		console.debug("[ChannelList:_deleteChannel]", id)
		await $.request("DELETE", "/api/v1/channels/:id",{
			params: {id: id}
		});
	}
}

export default ChannelList;
