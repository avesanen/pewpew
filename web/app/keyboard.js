define(['exports'], function (exports) {
    var keyStates = {};
    var keyUpBinds = {};
    var keyDownBinds = {};

    function onKeyboard(e) {
        if(e.type == "keydown") {
            if(!keyStates[e.keyCode]) {
                keyStates[e.keyCode] = true;
                if(keyDownBinds[e.keyCode]) {
                    keyDownBinds[e.keyCode]();
                }
            }
        }
        if(e.type == "keyup") {
            if(keyStates[e.keyCode]){
                keyStates[e.keyCode] = false;
                if(keyUpBinds[e.keyCode]) {
                    keyUpBinds[e.keyCode]();
                }
            }
        }
    }

    window.onkeydown = function(e){
        e.preventDefault();
        onKeyboard(e);
    };

    window.onkeyup = function(e){
        e.preventDefault();
        onKeyboard(e);
    };

    exports.keyDown = function(key) {
        if (keyStates[key]){
            return true;
        } else {
            return false;
        }
    };

    exports.releaseKey = function(key) {
        keyStates[key]=false;
    }

    exports.bindKeydown = function(key, func) {
        keyDownBinds[key] = func;
    }

    exports.bindKeyup = function(key, func) {
        keyUpBinds[key] = func;
    }
});