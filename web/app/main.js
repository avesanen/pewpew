requirejs.config({
	baseUrl: '/',
	paths: {
		lib: '/js/lib'
	}
});

require(["app/keyboard", "app/sprites", "app/particles", "app/sfx", "app/websocket"],
	function (keyboard, sprites, particles, sfx, websocket) {
		var fps = 60;
		websocket.connect();
		
		// Initialize mainCanvas
		particles.init('particleCanvas');
		sprites.init('spriteCanvas');

		keyboard.bindKeyup(function(e){
			websocket.send('keyup', JSON.stringify({"key":e.keyCode}));
		});

		keyboard.bindKeydown(function(e){
			websocket.send('keydown', JSON.stringify({"key":e.keyCode}));
		});


		window.addEventListener('mousedown', function(e) {
			var x = e.x||e.clientX;
  			var y = e.y||e.clientY;
  			websocket.send('mousedown', JSON.stringify({"x":x, "y":y}));
		});

		window.addEventListener('mousemove', function(e) {
			var x = e.x||e.clientX;
  			var y = e.y||e.clientY;
  			websocket.send('mouseover', JSON.stringify({"x":x, "y":y}));
		});

		websocket.on('gamestate', function(msg) {
			gamestate = JSON.parse(msg);
			entities = gamestate.players;
			if (gamestate.bullets) {
				entities = entities.concat(gamestate.bullets);
			}
			sprites.setSprites(entities);
		});

		// Start graphics loop
		var lastRefresh = new Date().getTime();
    	setInterval(function(){
    		// Calculate delta-time
    		var now = new Date().getTime();
        	var dt = now-lastRefresh;

        	// Refresh canvas entitiy positions
        	particles.refresh(dt);
        	sprites.refresh(dt);

        	// Redraw canvases
        	particles.draw();
        	sprites.draw();

        	// Set lastRefresh to calc. deltatime next time.
        	lastRefresh = now;
    	},1000/fps);
	}
);