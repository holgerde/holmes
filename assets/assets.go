// Code generated by include_assets.go. DO NOT EDIT.

package assets

var (
	Analyticsjs = `!function(e){function t(t){t.holmesId=e.localStorage.getItem("_holmesId");for(var n=0;n<o.length;n++)o[n](t);var a="__HOLMES_BASE_URL__/track?u="+(new Date).getTime()+"&t="+encodeURIComponent(JSON.stringify(t)),r=new XMLHttpRequest;r.open("GET",a),r.send()}var n,o=[];null===e.localStorage.getItem("_holmesId")&&e.localStorage.setItem("_holmesId",function(){for(var e=[],t=0;t<256;t++)e[t]=(t<16?"0":"")+t.toString(16);var n=4294967296*Math.random()>>>0,o=4294967296*Math.random()>>>0,a=4294967296*Math.random()>>>0,r=4294967296*Math.random()>>>0;return e[255&n]+e[n>>8&255]+e[n>>16&255]+e[n>>24&255]+"-"+e[255&o]+e[o>>8&255]+"-"+e[o>>16&15|64]+e[o>>24&255]+"-"+e[63&a|128]+e[a>>8&255]+"-"+e[a>>16&255]+e[a>>24&255]+e[255&r]+e[r>>8&255]+e[r>>16&255]+e[r>>24&255]}()),e.Holmes={pageView:function(e){e.type="PAGE_VIEW",t(e)},addTrackingEnricher:function(e){o.push(e)},track:t},(n=document.createEvent("Event")).initEvent("holmesloaded",!1,!1),e.dispatchEvent(n)}(window);`
	Bannertxt   = `
       ,_
     ,'  ` + "`" + `\,_       ██╗  ██╗ ██████╗ ██╗     ███╗   ███╗███████╗███████╗
     |_,-'_)        ██║  ██║██╔═══██╗██║     ████╗ ████║██╔════╝██╔════╝
     /##c '\  (     ███████║██║   ██║██║     ██╔████╔██║█████╗  ███████╗
    ' |'  -{.  )    ██╔══██║██║   ██║██║     ██║╚██╔╝██║██╔══╝  ╚════██║
      /\__-' \[]    ██║  ██║╚██████╔╝███████╗██║ ╚═╝ ██║███████╗███████║
     /` + "`" + `-_` + "`" + `\         ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝     ╚═╝╚══════╝╚══════╝
     '     \
`
)
