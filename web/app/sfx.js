define(['exports'], function (exports) {
    // Init AudioContext
    window.AudioContext = window.AudioContext||window.webkitAudioContext;
    var context = new AudioContext();


    var gainNode = context.createGain();

    var sounds = {};
    function loadSound(name, url) {
        var request = new XMLHttpRequest();
        request.open('GET', url, true);
        request.responseType = 'arraybuffer';
        var audiobuffer;

        request.onload = function() {
            if (request.status == 200) {
                context.decodeAudioData(request.response, function(buffer) {
                    sounds[name] = buffer;
                });
            }
        };
        request.send();
    }

    exports.init = function(soundlist) {
        for(var i = 0; i < soundlist.length;i++){
            loadSound(soundlist[i].name,soundlist[i].url);
        }
    };

    function makeSource(buffer) {
        var source = context.createBufferSource();
        source.buffer = buffer;

        var gain = context.createGain();
        gain.gain.value = 0.01;
        
        source.connect(gain);
        gain.connect(context.destination);

        return source;
    }

    exports.play = function(name) {
        var source = makeSource(sounds[name]);
        //source.playbackRate.value = 1 + Math.random();
        source.start(0);
    };
});