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
        canvas.width = 800;
        canvas.height = 600;
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
            if (sprites[i].position && sprites[i].velocity)  {
                sprites[i].position.x += (sprites[i].velocity.x * dt) / 1000;
                sprites[i].position.y += (sprites[i].velocity.y * dt) / 1000;
            }
        }
    }

    exports.draw = function() {
        clearCanvas();
        for(var i = 0; i < sprites.length; i++){
            if (sprites[i].radius) {
                drawCircle(sprites[i].position, sprites[i].radius, [255,0,255,255]);
            } else if (sprites[i].width && sprites[i].height) {
                drawRectangle(sprites[i].position, sprites[i].width, sprites[i].height, [255,255,255,255]);
            } else {
                drawCircle(sprites[i].position, 5, [255,255,255,255]);
            }
            if (sprites[i].lookingAt) {
                drawCircle(sprites[i].lookingAt, 0.3, [255,0,0,128]);
                drawLine(sprites[i].position, sprites[i].lookingAt, [255,0,0,128]);
            }
        }
    }

    var clearCanvas = function() {
        canvas.width = canvas.width;
    };

    var drawRectangle = function(position, width, height, color) {
        var x1 = (position.x-width/2) * 10;
        var y1 = (position.y-height/2) * 10;
        var x2 = width * 10;
        var y2 = height * 10;
        ctx.fillStyle = toCss(color);
        ctx.rect(x1,y1,x2,y2);
        ctx.stroke();
    }

    var drawCircle = function(position, size, color) {
        var x = position.x * 10;
        var y = position.y * 10;
        var r = size * 10;
        ctx.fillStyle = toCss(color);
        ctx.beginPath();
        ctx.arc(x, y, r, 0, Math.PI * 2, true);
        ctx.closePath();
        ctx.fill();
    };

    var drawLine = function(positiona, positionb, color) {
        var x1 = positiona.x * 10;
        var y1 = positiona.y * 10;
        var x2 = positionb.x * 10;
        var y2 = positionb.y * 10;
        ctx.strokeStyle = toCss(color);
        ctx.beginPath();
        ctx.moveTo(x1,y1);
        ctx.lineTo(x2,y2);
        ctx.stroke();
    }

});