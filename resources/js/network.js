function e_lobby_created(data) {
	state.map_height = data.height;
	state.map_width = data.width;
	state.mine_count = data.mines;
	state.code = data.code;
	state.tiles = encoder.encode(atob(data.map));
	load_game_menu();
}

function e_map_changed(data) {
	let tiles = encoder.encode(atob(data.map));
	update_game_menu(tiles);
}

function e_turn_changed(data) {
	state.turn = data.turn;
	update_turn();
}

function e_game_end(data) {
	if(data.victory) {
		alert("you win");
	}else{
		alert("you lose");
	}
}

function create_lobby(width, height, mines) {
	let a = JSON.stringify({
		id: "create_lobby",
		data: {
			width: width,
			height: height,
			mines: mines
		}
	});
	console.log(a);
	socket.send(a);
}

function join_lobby(code) {
	socket.send(JSON.stringify({
		id: "join_lobby",
		data: {
			code: code 
		}
	}));
}

function tile_clicked(tile) {
	if(state.turn) {
		let tile_id = parseInt(tile.getAttribute("offset"));
		socket.send(JSON.stringify({
			id: "tile_clicked",
			data: {
				offset: tile_id
			}
		}));
	}
}

function request_map_reset() {
	socket.send(JSON.stringify({
		id: "reset",
		data:{
			reset: true
		}
	}));
}