let pass = document.getElementById("p");
let cpass = document.getElementById("cp");
let cpassl=document.getElementById('cpl')
let error = document.getElementById("error");
let user = document.getElementById("user");
let userl = document.getElementById("userl");

function validate() {
  if (pass.value != cpass.value) {
    cpassl.style.color='red'
    cpass.style['border-bottom']='1px solid red'
    error.style.visibility = "visible";
    setTimeout(() => {
      error.style.visibility = "hidden";
      cpassl.style.color='black'
      cpass.style['border-bottom']='1px solid black'
    }, 3000);
    return false;
  } else {
    return true;
  }
}

function visible(ele,mainEle) {
  let x=document.getElementById(mainEle)
  if(ele.type==='password'){
    x.innerHTML=`<i class="fa-regular fa-eye"></i>`
    x.setAttribute('title','Hide')
    ele.type='text'
  }
  else{
    x.innerHTML=`<i class="fa-regular fa-eye-slash"></i>`
    x.setAttribute('title','Show')
    ele.type='password'
  }
}

function checkExistUser(user, userl) {
        let userNameExist = "{{.Isusername}}";
        let userErrorMsg = document.getElementById("userErrorMsg");
        if (userNameExist == "true") {
          user.style["border-bottom"] = "2px solid red";
          userl.style.color = "red";
          userErrorMsg.style.display = "block";
          setTimeout(function () {
            userErrorMsg.style.display = "none";
            user.style["border-bottom"] = "1px solid black";
            userl.style.color = "black";
          }, 4000);
        }
      }

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
    checkExistUser(user,userl)
}