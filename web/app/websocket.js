define(['exports'], function (exports) {
    var messageHandlers = {};
    var ws = {};

    var callHandler = function(type, msg){
        if (messageHandlers[type]) {
            for(var i = 0; i < messageHandlers[type].length;i++){
                messageHandlers[type][i](msg);
            }
        } 
    };

    exports.on = function(type, handler) {
        if (!messageHandlers[type]) {
            messageHandlers[type] = [];
        }
        messageHandlers[type].push(handler);
    }

    exports.connect = function() {
        ws = new WebSocket("ws://"+window.location.host+window.location.pathname+"ws/");
        ws.onopen = function(e){
            callHandler("connect", e);
        };
        
        ws.onmessage = function(msg){
            message = JSON.parse(msg.data);
            callHandler(message.type, message.msg);
        }
    }

    exports.send = function(type, msg) {
        ws.send(JSON.stringify({type:type, msg:msg}));
    }
});