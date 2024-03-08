let encoder = new TextEncoder();

let state = {
	map_width: 0,
	map_height: 0,
	mine_count: 0,
	tiles: [],
	turn: false,
	code: ""
}

let socket = new WebSocket("ws://" + document.location.host + "/ws");
socket.onmessage = function(e) {
	let packet = JSON.parse(e.data);
	console.log(packet)
	switch (packet.id) {
		case "lobby_created":
			e_lobby_created(packet);
			break;
		case "map_changed":
			e_map_changed(packet);
			break;
		case "turn_changed":
			e_turn_changed(packet);
			break;
		case "game_end":
			e_game_end(packet);
			break;
		default:
			break;
	}
}

load_main_menu();