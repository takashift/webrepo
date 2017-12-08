use strict;
// var xhr = new XMLHttpRequest();

xhr.setRequestHeader("Authorization", localStorage.jwt);

localStorage.jwt = xhr.getResponseHeader("Authorization");

function getHeader(){
	$.ajax({
		type : 'HEAD',
	}).done(function(data, textStatus, xhr) {
		localStorage.jwt = xhr.getResponseHeader("Authorization");
    })
}
getHeader();

onload = function() {
	var element = document.getElementById({{}});
	element.setAttribute("selected", "");	
}

document.getElementById("review_list").insertAdjacentHTML('afterbegin', document.createTextNode({{.Content}}));


var xhr;
onload = function () {
  xhr = new XMLHttpRequest();
  xhr.open("Get", "https://api.twitter.com/1.1/search/tweets.json?q=%23本日の評価ページお題&result_type=recent&count=1", false);
  xhr.setRequestHeader('Authorization', 'Bearer ' + 'AAAAAAAAAAAAAAAAAAAAADEp3gAAAAAA1utz8ysPQJLfahTh%2BIievBRebMw%3DJcaj7Z0fj0edHPVOZADriheIMugHxR1VsyZUod0y2eNQHHhbUS');
  xhr.onreadystatechange = callback;
  xhr.send(null);
}
function callback() {
  if (xhr.status == 200) {
	var myHeader = document.getElementById('daily_theme');
	myHeader.value = xhr.responseText;
	var tNode = document.createTextNode(xhr.responseText);
	myHeader.appendChild(tNode);
  }
}
