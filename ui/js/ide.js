let editor;

window.onload = function() {
    editor = ace.edit("editor");
    editor.setTheme("ace/theme/monokai");

}

function changeCompiler() {
}

function compileCode() {
    let compiler = $("#compilers").val();
    let code = editor.getSession().getValue();

    $.ajax({
        type: 'post',
        url: "http://localhost:8085/app/compile",
        crossDomain: true,
        data: JSON.stringify({'compiler': compiler, 'code': code}),
        success: function(response) {
            console.log(response);
            $(".output").text(response)
        },
        error: function(response) {
            console.log(response);
            console.log("error");
            $(".output").text(response)
        }
    })
}