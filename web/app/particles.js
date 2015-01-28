define(['exports'], function (exports) {
    var canvas = null;
    var ctx = null;

    exports.init = function(canvasId) {
        canvas = document.getElementById(canvasId);
        ctx = canvas.getContext('2d');
        document.body.onresize = function(){
  			ctx.canvas.width = window.innerWidth;
  			ctx.canvas.height = window.innerHeight;
		};
		document.body.onresize();
    };

    var toCss = function(A) {
        for(var i=0;i<3;i++){
            A[i]=Math.round(A[i]);
        }
        A[3]/=255;
        return "rgba("+ A.join()+")";
    };

    var lerp = function(p, a, b){
        return Number(a)+(b-a)*p;
    };

    var lerpA = function(p, A, B){
        var res=[];
        for(var i=0; i<A.length; i++){
            res[i]=lerp(p, A[i], B[i]);
        }
        return res;
    };

    var easeOut = function(p, a, b){
    	return (b-a) * Math.sin(p * (Math.PI/2)) + a;
    }

    var particles = [];

    var Particle = function(x,y,speed,endspeed,life,startcolor,endcolor,startsize,endsize,startangle){
        this.x = x;
        this.y = y;
        this.startspeed = speed;
        this.endspeed = endspeed;

        this.startangle = startangle;
        this.endangle = startangle;

        this.life = life;
        this.startlife = this.life;

        this.startsize = startsize;
        this.endsize = endsize;

        this.startcolor = startcolor;
        this.endcolor = endcolor;
    };

    Particle.prototype.refresh = function(dt) {
        this.life -= dt;
        var speed = easeOut(this.life/this.startlife, this.endspeed, this.startspeed);
        var angle = lerp(this.life/this.startlife, this.endangle, this.startangle);
        this.x += speed * Math.sin(angle * Math.PI / 180) * dt / 1000;
        this.y += -speed * Math.cos(angle * Math.PI / 180) * dt / 1000;
        this.size = lerp(this.life/this.startlife, this.endsize, this.startsize);
        this.color = lerpA(this.life/this.startlife, this.endcolor, this.startcolor);
    };

    Particle.prototype.draw = function() {
        ctx.fillStyle = toCss(this.color);
        ctx.beginPath();
        ctx.arc(this.x, this.y, this.size, 0, Math.PI * 2, true);
        ctx.closePath();
        ctx.fill();
    };

    exports.emitter = function(x,y,amount) {
        exports.emit(x,y,Math.random()*360,40,1000,[255,255,255,128],[1,1,1,1],1,10);
        setTimeout(function(){
            if (this.attached) {
                this.x = this.attached.x;
                this.y = this.attached.y;
            }
            if (amount > 0) {
                amount -= 1;
                exports.emitter(x,y,amount);
            }},1000/30);
    };

    exports.emitter.prototype.attachTo = function(spr) {
        this.attached = spr;
    };

    exports.emit = function(x,y,speed,endspeed,life,startcolor,endcolor,startsize,endsize,startangle,endangle) {
        particles.push(new Particle(x,y,speed,endspeed,life,startcolor,endcolor,startsize,endsize,startangle,endangle));
    };

    exports.refresh = function(dt) {
        for(var i = 0; i < particles.length;i++){
            if(particles[i].life > 0) {
                particles[i].refresh(dt);
            } else {
                particles.splice(i,1);
            }
        }
    };

    exports.draw = function() {
        canvas.width = canvas.width;
        for(var i = 0; i < particles.length;i++){
            if(particles[i].life > 0) {
                particles[i].draw();
            } else {
                // TODO: This is slow?
                particles.splice(i,1);
            }
        }
    };

});