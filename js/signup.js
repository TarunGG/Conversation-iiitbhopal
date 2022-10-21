let pass = document.getElementById("p");
let cpass = document.getElementById("cp");
let cpassl=document.getElementById('cpl')
let error = document.getElementById("error");

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