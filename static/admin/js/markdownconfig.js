/**
 * Created by RoryGao on 16/1/1.
 */
function configMarkdown() {
    //markdown editor
    $("textarea[data-provide='markdown']").markdown({
        language: 'zh',
        fullscreen: {
            enable: true
        },
        resize: 'vertical'
    });
}