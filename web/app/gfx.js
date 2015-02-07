define(['exports'], function (exports) {

    exports.init = function(canvasId){
        this.canvas = document.getElementById(canvasId);
        this.ctx = this.canvas.getContext('2d');
        this.canvas.width = 800;
        this.canvas.height = 600;
        this.clear();
    };

    exports.clear = function() {
        this.canvas.width = this.canvas.width;
    };

    exports.drawCircle = function(location, size, color) {
        this.ctx.fillStyle = toCss(color);
        this.ctx.beginPath();
        this.ctx.arc(location.x*10, location.y*10, size, 0, Math.PI * 2, true);
        this.ctx.closePath();
        this.ctx.fill();
    };

    exports.drawLine = function(locationa, locationb, color) {
        this.ctx.strokeStyle = toCss(color);
        this.ctx.beginPath();
        this.ctx.moveTo(locationa[0],locationa[1]);
        this.ctx.lineTo(locationb[0],locationb[1]);
        this.ctx.stroke();
    };

    var toCss = function(A) {
        for(var i=0;i<3;i++){
            A[i]=Math.round(A[i]);
        }
        A[3]/=255;
        return "rgba("+ A.join()+")";
    };
});