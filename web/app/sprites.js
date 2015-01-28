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
        canvas.width = canvas.width;
        for(var i = 0; i < sprites.length; i++){
            size = 10
            if (sprites[i].type == "player") {
                ctx.fillStyle = toCss([255,0,255,255]);
                size = 10
            } else {
                ctx.fillStyle = toCss([255,255,255,255]);
                size = 3
            }
            ctx.beginPath();
            ctx.arc(sprites[i].location[0], sprites[i].location[1], size, 0, Math.PI * 2, true);
            ctx.closePath();
            ctx.fill();
        }
    }

});