define(['exports'], function (exports) {
    var canvas = null;
    var ctx = null;
    var sprites = [];

    exports.setSprites = function(spriteArray) {
        sprites = spriteArray;
    }

    exports.init = function(canvasId) {
        canvas = document.getElementById(canvasId);
        ctx = canvas.getContext('2d');
        ctx.canvas.width = window.innerWidth;
        ctx.canvas.height = window.innerHeight;
    };

    var toCss = function(A) {
        for(var i=0;i<3;i++){
            A[i]=Math.round(A[i]);
        }
        A[3]/=255;
        return "rgba("+ A.join()+")";
    };

    exports.refresh = function(dt) {
        for(var i = 0; i < sprites.length; i++){
            sprites[i].location[0] += sprites[i].velocity[0] * dt / 1000;
            sprites[i].location[1] += sprites[i].velocity[1] * dt / 1000;
        }
    }

    exports.draw = function() {
        clearCanvas();
        for(var i = 0; i < sprites.length; i++){
            if (sprites[i].type == "player") {
                drawCircle(sprites[i].location, 10, [255,0,255,255]);
                drawCircle(sprites[i].aiming, 2, [255,0,0,255]);
                drawLine(sprites[i].location, sprites[i].aiming, [255,0,0,128])
            } else {
                drawCircle(sprites[i].location, 3, [255,255,255,255]);
            }
        }
    }

    var clearCanvas = function() {
        canvas.width = canvas.width;
    };

    var drawCircle = function(location, size, color) {
        ctx.fillStyle = toCss(color);
        ctx.beginPath();
        ctx.arc(location[0], location[1], size, 0, Math.PI * 2, true);
        ctx.closePath();
        ctx.fill();
    };

    var drawLine = function(locationa, locationb, color) {
        ctx.strokeStyle = toCss(color);
        ctx.beginPath();
        ctx.moveTo(locationa[0],locationa[1]);
        ctx.lineTo(locationb[0],locationb[1]);
        ctx.stroke();
    }

});