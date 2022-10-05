let editor;

window.onload = () => {
    editor = ace.edit("editor");
    editor.setTheme("ace/theme/github");
}

const compileCode = () => {
    const compilerSelect = $('#compiler')
    const runBtn = $('#run-btn')
    const outputBlock = $('#output')

    runBtn.attr('disabled', true)

    const compiler = compilerSelect.val();
    const code = editor.getSession().getValue();

    $.ajax({
        type: 'post',
        url: "http://localhost:8085/app/compile",
        crossDomain: true,
        data: JSON.stringify({ 'compiler': compiler, 'code': code }),
        success: function (response) {
            console.log(response);
            outputBlock.text(response)
        },
        error: function (response) {
            console.log(response);
            console.error("error");
            outputBlock.text(JSON.stringify(response))
        },
        complete: function () {
            runBtn.attr('disabled', false)
        }
    })
}

const toggleMode = () => {
    const root = $(':root');
    const toggleBtn = $('#toggle-btn');

    const theme = root.attr('theme') ?? 'light';

    if (theme === 'light') {
        root.attr('theme', 'dark');
        toggleBtn.text('Dark');
        editor.setTheme("ace/theme/monokai");
    } else {
        root.attr('theme', 'light');
        toggleBtn.text('Light');
        editor.setTheme("ace/theme/github");
    }
}
