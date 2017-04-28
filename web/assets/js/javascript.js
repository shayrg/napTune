$(document).ready(function(){
    loadSongs();
});
function loadSongs(){
    $("#player").hide();
    $("#player").on("ended",nextSong);
    $.ajax({
        url: "/songs",
        type: "json",
        success: function(data){
            $(data).each(function(i){
                var $div = $('<div>');
                var $p = $('<p>');
                var $a = $('<button>');
                //Button
                $a.click(changeSource);
                $a.text("Play");
                $a.attr("href", this.location);
                //Paragraph
                $p.text(this.name + ", " + this.length + " ");
                $p.append($a);
                //Div
                $div.append($p);
                //Playlist
                $("#playlist").append($div);
                if(i == 0){
                    $a.click();
                }
            });
        }
    });
}
function changeSource() {
    $("#player").attr("src", $(this).attr("href"));
    $("#playing").attr("id","");
    $(this).attr("id","playing");
    //Play
    playSong();
    return false;
}
function playSong(){
    $("#player").each(function(){
        this.play();
    });
}
function pauseSong(){
    $("#player").each(function(){
        this.pause();
    });
}
function nextSong(){
    $('#playing').parent().parent().next().children().children().click();
}