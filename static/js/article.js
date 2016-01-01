/**
 * Created by RoryGao on 15/12/16.
 */
function convertArticle2HTML(id, data) {
    $('#article-' + id).html(markdown.toHTML(data))
}