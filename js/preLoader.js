let mainCode=document.getElementById('mainCode')
mainCode.style.display='none'
let preLoader=document.getElementById('preLoader')
let str = `             <div class="spinner-box">
<div class="configure-border-1">
<div class="configure-core"></div>
</div>
<div class="configure-border-2">
<div class="configure-core"></div>
</div>
<h3 id="loadingTxt">Loading...</h3>
</div>`
preLoader.innerHTML=str

window.onload=function(){
    preLoader.style.display='none'
    mainCode.style.display='block'
}