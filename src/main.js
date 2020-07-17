function hideforms() {
    var forms = document.getElementsByClassName("autonewsform");
    for (var i = 0; i < forms.length; i++) {
        forms[i].style.display = "none";
    }
    var btns = document.getElementsByClassName("selbtn");
    for (var i = 0; i < btns.length; i++) {
        btns[i].classList.remove("active")
    }
}

function button(timeslotid) {
    hideforms();
    document.getElementById("form" + timeslotid).style.display = "block";
    document.getElementById("btn" + timeslotid).classList.add("active");
}

hideforms();
setTimeout(function() {
    console.log("Possible Error Here: Don't Worry :)");
    document.getElementById("success-alert").style.display = "none";
}, 3000);

setTimeout(function() { window.location.reload(); }, 5000);

function completeABox(boxNo) {
    document.getElementById("spinner" + boxNo).style.display = "none";
    document.getElementById("check" + boxNo).style.display = "block";
    document.getElementById("li" + boxNo).classList.remove("list-group-item-dark");
    document.getElementById("li" + boxNo).classList.add("list-group-item-success");
}

var listElements = document.getElementsByClassName("selector-component");
for (i = 0; i < listElements.length; i++) {
    listElements[i].style.display = "inline";
}

document.getElementById("autoselector").style.display = "block";

var t = new Date()
document.getElementById("text1").innerText = t.getHours() + ":59:45 - Switch to AutoNews";
document.getElementById("text2").innerText = (t.getHours() + 1) + ":00:00 - No Action";
document.getElementById("text3").innerText = (t.getHours() + 1) + ":02:02 - Switch to Jukebox";