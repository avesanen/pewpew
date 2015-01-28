requirejs.config({
	baseUrl: '/',
	paths: {
		lib: '/js/lib'
	}
});

require(["app/keyboard", "app/sprites", "app/particles", "app/sfx"],
	function (keyboard, sprites, particles, sfx) {
		var fps = 60;
		var socket = io();

		// Initialize mainCanvas
		particles.init('particleCanvas');
		sprites.init('spriteCanvas');
		sfx.init([{name:"firework", url:"/sounds/firework.mp3"}]);

		// Keyboard bindings (quick and dirty edition)
		keyboard.bindKeyup(87, function(){
			socket.emit('keyup', JSON.stringify({"key":"w"}));
		});
		keyboard.bindKeyup(65, function(){
			socket.emit('keyup', JSON.stringify({"key":"a"}));
		});
		keyboard.bindKeyup(83, function(){
			socket.emit('keyup', JSON.stringify({"key":"s"}));
		});
		keyboard.bindKeyup(68, function(){
			socket.emit('keyup', JSON.stringify({"key":"d"}));
		});

		keyboard.bindKeydown(87, function(){
			socket.emit('keydown', JSON.stringify({"key":"w"}));
		});
		keyboard.bindKeydown(65, function(){
			socket.emit('keydown', JSON.stringify({"key":"a"}));
		});
		keyboard.bindKeydown(83, function(){
			socket.emit('keydown', JSON.stringify({"key":"s"}));
		});
		keyboard.bindKeydown(68, function(){
			socket.emit('keydown', JSON.stringify({"key":"d"}));
		});

		window.addEventListener('mousedown', function(e) {
			var x = e.x||e.clientX;
  			var y = e.y||e.clientY;
  			socket.emit('mousedown', JSON.stringify({"x":x, "y":y}));
		});

		socket.on('gamestate', function(msg) {
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