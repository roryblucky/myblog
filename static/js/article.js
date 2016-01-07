/**
 * Created by RoryGao on 15/12/16.
 */
function convertArticle2HTML(id, data) {
    $('#article-' + id).html(subContent(markdown.toHTML(data), 200))
}

function subContent(data, leng) {
    if (data.length == leng || data.length < leng) {
        return data;
    }
    var newData = data.substring(0, leng);
    console.log(newData);
    newData = newData + ".....";
    return newData
}