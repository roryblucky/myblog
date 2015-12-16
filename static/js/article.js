/**
 * Created by RoryGao on 15/12/16.
 */
function convertArticle2HTML(data) {
    $(".article").html(markdown.toHTML(data))
}