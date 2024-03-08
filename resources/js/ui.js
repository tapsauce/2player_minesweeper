function load_main_menu() {
	$("#body").load("pages/mainmenu.html", function () {
		let join_lobbybtn = document.getElementById("join_lobby");
		let lobby_codebox = document.getElementById("lobby_code");

		let create_lobbybtn = document.getElementById("create_lobby");
		let widthbox = document.getElementById("width");
		let heightbox = document.getElementById("height");
		let minesbox = document.getElementById("mines");

		lobby_codebox.addEventListener("keyup", function () {
			let code = lobby_codebox.value;
			join_lobbybtn.disabled = code.length < 1;
		});

		join_lobbybtn.addEventListener("click", function () {
			let code = lobby_codebox.value;
			if (code.length != 0) {
				join_lobby(code);
			} else {
				alert("code empty");
			}
		});

		const create_lobby_update = function() {
			let width = parseInt(widthbox.value);
			let height = parseInt(heightbox.value);
			let mines = parseInt(minesbox.value);
			let enabled = width !== NaN && width > 0 && height !== NaN && height > 0 && mines !== NaN && mines > 0;
			create_lobbybtn.disabled = !enabled;
		}

		widthbox.addEventListener("keyup", create_lobby_update);
		heightbox.addEventListener("keyup", create_lobby_update);
		minesbox.addEventListener("keyup", create_lobby_update);

		create_lobbybtn.addEventListener("click", function () {
			let width = parseInt(widthbox.value);
			let height = parseInt(heightbox.value);
			let mines = parseInt(minesbox.value);

			if (width !== NaN && height !== NaN && mines !== NaN) {
				if (width < 40 || height < 40) {
					create_lobby(width,height,mines);
				}
			}
		});
	});
}

function load_game_menu() {
	$("#body").load("pages/game.html", function() {
		let game = document.getElementById("arena");
		for (let offset = 0; offset< state.map_width * state.map_height; offset++) {
			if(offset%state.map_width == 0) {
				game.appendChild(document.createElement("br"));
			}
			const tile = document.createElement("button");
			tile.classList.add("tile");
			tile.classList.add("tile_hidden");
			tile.setAttribute("offset", offset);
			tile.addEventListener("click",() => {
				tile_clicked(tile);
			});
			if(state.tiles[offset] == 10) {
				tile.innerText = " ";
			}else{
				tile.innerText = state.tiles[offset];
			}
			game.appendChild(tile);
			
		}
		update_code();
	});
}

function update_game_menu(tiles) {
	let old = state.tiles;
	for (let i = 0; i < tiles.length; i++) {
		if(old[i] != tiles[i]) {
			let tile = document.querySelector(`[offset="${i}"]`);
			if(tiles[i] != 9) {
				tile.innerText = tiles[i];
			}else{
				tile.innerText = "★";
			}
			tile.classList.remove(`tile${old[i]}`);
			tile.classList.remove("tile_hidden");
			tile.classList.add(`tile${tiles[i]}`);
		}
	}
	state.tiles = tiles;
}

function update_turn() {
	document.getElementById("turn").innerText = `Your turn: ${state.turn}`;
}

function update_code() {
	document.getElementById("code").innerText = `Lobby code: ${state.code}`;
}