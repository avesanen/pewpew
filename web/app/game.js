define(function(require){
		var fps = 60;
		var particles = require("app/particles");
		var sprites = require("app/sprites");
		var websocket = require("app/websocket");
		var keyboard = require("app/keyboard");
		var gore = require("app/gfx");
		var terminal = require("app/console");

		var maskdiv = document.getElementById("maskdiv");
		document.oncontextmenu = function(e){return false};

		particles.init('particleCanvas');
		sprites.init('spriteCanvas');
		gore.init('goreCanvas');
		websocket.connect();

		keyboard.bindKeyup(function(e){
			websocket.send('keyup', JSON.stringify({"key":e.keyCode}));
		});

		keyboard.bindKeydown(function(e){
			websocket.send('keydown', JSON.stringify({"key":e.keyCode}));
		});


		window.addEventListener('mousedown', function(e) {
			//e.preventDefault();
			var x = e.x||e.clientX;
  			var y = e.y||e.clientY;
  			websocket.send('mousedown', JSON.stringify({"x":x/10, "y":y/10}));
		});

		window.addEventListener('mousemove', function(e) {
			//e.preventDefault();
			var x = e.x||e.clientX;
  			var y = e.y||e.clientY;
  			websocket.send('mouseover', JSON.stringify({"x":x/10, "y":y/10}));
		});

		websocket.on('gamestate', function(msg) {
			gamestate = JSON.parse(msg);
			entities = [];
			if(gamestate.collisions) {
				console.log(gamestate.collisions);
				for(var i = 0; i < gamestate.collisions.length; i++){
					console.log(gamestate.collisions[i].position)
					gore.drawCircle(gamestate.collisions[i].position, 5, [255,255,255,255])
				}
			}
			if (gamestate.players) {
				entities = entities.concat(gamestate.players);
			}
			if (gamestate.bullets) {
				entities = entities.concat(gamestate.bullets);
			}
			if (gamestate.tiles) {
				entities = entities.concat(gamestate.tiles);
			}
			sprites.setSprites(entities);
		});

		var lastRefresh = new Date().getTime();
    	setInterval(function(){
    		// Calculate delta-time
    		var now = new Date().getTime();
        	var dt = now-lastRefresh;

        	// Draw sprites
        	sprites.refresh(dt);
        	sprites.draw();

        	// Set new lastRefresh
        	lastRefresh = now;
    	},1000/fps);
});