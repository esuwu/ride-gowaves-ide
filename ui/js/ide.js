let editor;

window.onload = function() {
    editor = ace.edit("editor");
    editor.setTheme("ace/theme/monokai");

}

function changeCompiler() {
    let compiler = $("#compiler").val();

    if (compiler == "GowavesCompiler") {
       // send request to gowaves compiler
    }

    if (compiler == "WavesCompiler") {
        // send request to waves compiler
    }

}

// function executeCode() {
//     $.sjax({
//         url: "/execute",
//     })
// }