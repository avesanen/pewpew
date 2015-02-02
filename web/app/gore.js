define(['exports'], function (exports) {
    var canvas = null;
    var ctx = null;

    exports.init = function(canvasId) {
        canvas = document.getElementById(canvasId);
        ctx = canvas.getContext('2d');
        ctx.canvas.width = window.innerWidth;
        ctx.canvas.height = window.innerHeight;
    };

    exports.splatter = function(location) {
        drawCircle(location, 3, [255,0,0,128]);
    };

    var toCss = function(A) {
        for(var i=0;i<3;i++){
            A[i]=Math.round(A[i]);
        }
        A[3]/=255;
        return "rgba("+ A.join()+")";
    };



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

    exports.drawCircle = drawCircle;
    exports.drawLine = drawLine;
    exports.clearCanvas = clearCanvas;

});d