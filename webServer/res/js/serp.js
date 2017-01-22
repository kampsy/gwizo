/*navigation line*/
const horiLine = document.querySelector("div.hori-line-cognifly");
const border = document.querySelector("div.material-wrapper-cognifly");
if (horiLine && border) {
	addEventListener("scroll", function() {
	  let max = document.body.scrollHeight - innerHeight;
	  let percent = (pageYOffset / max) * 100;
	  horiLine.style.width = percent + "%";
		if (pageYOffset >= 43) {
			border.style.cssText = "border-bottom: 1px solid #ebebeb;";
		}else if (pageYOffset < 43) {
			border.style.cssText = "border-bottom: 1px solid transparent;";
		}
	});
}

// mobile navigation drawer
const show = document.getElementById("func-mobile-nav-show");
const hide = document.getElementById("func-mobile-nav-hide")
const nav = document.querySelector("div.mobile-nav-bar-cognifly");
if (show && hide && nav) {
	show.addEventListener("click", function() {
    nav.style.cssText = "margin-left: 0px;";
		document.body.style.cssText = "overflow-y: hidden;";
  }, false);
  hide.addEventListener("click", function() {
    nav.style.cssText = "margin-left: -260px;";
		document.body.style.cssText = "overflow-y: auto;";
  }, false);
}

//tabs switch scopes
let tabs = document.querySelectorAll("div.tab-scopes-cognifly");
const tabWeather = document.querySelector("section.weather-scope-cogni-ux-5");
const tabWiki = document.querySelector("section.wikipedia-scope-cogni-ux-5");
const tabRecipes = document.querySelector("section.recipes-scope-cogni-ux-5");
addEventListener("click", function(evt){
	if (evt.target.className == "tab-scopes-cognifly") {
		const tag = evt.target.id;
		for (let i = 0; i < tabs.length; i++) {
			tabs[i].style.cssText = "border-bottom: 2px solid #FFF;\
		  color: #848484;";
		}
		evt.target.style.cssText = "border-bottom: 2px solid #1E88E5;\
		color: #1E88E5;";
		if (tag == "func-weather-cognifly") {
			tabWeather.style.cssText = "display: block;";
			tabWiki.style.cssText = "display: none;";
			tabRecipes.style.cssText = "display: none;";
		}else if (tag == "func-wikipedia-cognifly") {
			tabWeather.style.cssText = "display: none;";
			tabWiki.style.cssText = "display: block;";
			tabRecipes.style.cssText = "display: none;";
		}else if (tag == "func-recipes-cognifly") {
			tabWeather.style.cssText = "display: none;";
			tabWiki.style.cssText = "display: none;";
			tabRecipes.style.cssText = "display: block;";
		}
	}
}, false);

//tabs switch between previews
let tabsPre = document.querySelectorAll("div.tab-btn-cognifly");
const tabIDE = document.querySelector("section.ide-cogni-ux-5");
const tabArch = document.querySelector("section.arch-cogni-ux-5");
addEventListener("click", function(evt){
	if (evt.target.className == "tab-btn-cognifly") {
		const tag = evt.target.id;
		for (let i = 0; i < tabsPre.length; i++) {
			tabsPre[i].style.cssText = "border-bottom: 2px solid #FFF;\
		  color: #848484;";
		}
		evt.target.style.cssText = "border-bottom: 2px solid #1E88E5;\
		color: #1E88E5;";
		if (tag == "func-ide-cognifly") {
			tabIDE.style.cssText = "display: block;";
			tabArch.style.cssText = "display: none;";
		}else if (tag == "func-arch-cognifly") {
			tabIDE.style.cssText = "display: none;";
			tabArch.style.cssText = "display: block;";
		}
	}
}, false);
