use strict;
// var xhr = new XMLHttpRequest();

xhr.setRequestHeader("Authorization", localStorage.jwt)

localStorage.jwt = xhr.getResponseHeader("Authorization")

function getHeader(){
	$.ajax({
		type : 'HEAD',
	}).done(function(data, textStatus, xhr) {
		localStorage.jwt = xhr.getResponseHeader("Authorization");
    })
}
getHeader();

