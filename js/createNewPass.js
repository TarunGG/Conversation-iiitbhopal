let alert = document.getElementById("alert");
setTimeout(function () {
  alert.style.display = "none";
}, 5000);

let pass = document.getElementById("pass");
let cpass = document.getElementById("cpass");
let cpassl = document.getElementById("cpassl");
let error = document.getElementById("error");
let container = document.getElementById("mainContainer");
let thankyou = document.getElementById("thankyouDiv");
let form = document.getElementById("form");
function validate() {
  if (pass.value != cpass.value) {
    error.style.visibility = "visible";
    cpassl.style.color = "red";
    cpass.style["border-bottom"] = "1px solid red";
    setTimeout(() => {
      error.style.visibility = "hidden";
      cpassl.style.color = "black";
      cpass.style["border-bottom"] = "1px solid black";
    }, 3000);
    return false;
  } else {
    // container.style.display = "none";
    // thankyou.innerHTML = `
// <span
//   >Your password has been set successfully. <br />
//   <hr
//     style="
//       margin-top: calc(0.5rem + 0.1vh);
//       margin-bottom: calc(0.5rem + 0.1vh);
//     "
//   />
//   <p style="font-size: calc(0.6rem + 0.5vw); text-align: center">
//     Not redirected to login! <a href="/login">click here</a>
//   </p>
// </span>`;
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