console.log("working");

//Border shape
let card = document.getElementsByClassName("card");
let cardArr = Array.from(card);
// let noOfCard = Array.from(card).length;

// for (let i = 0; i < noOfCard; i++) {
//   threadId = `thread${i + 1}`;
//   console.log(threadId);
//   thread = document.getElementById(threadId);
//   if (i % 2 == 0) thread.style["border-bottom-left-radius"] = "0";
//   else thread.style["border-bottom-right-radius"] = "0";
// }

//Like button
// function like(ele) {
//   let likeEle = document.getElementById(ele);
//   let likeEleHtml = likeEle.innerHTML;
//   console.log(likeEleHtml);
//   let str = `<i class="fa-regular fa-heart"></i>`;
//   let str2 = `<i class="fa-solid fa-heart"></i>`;
//   if (likeEleHtml == str) {
//     likeEle.innerHTML = str2;
//   } else {
//     likeEle.innerHTML = str;
//   }
// }


//LogOut Confirmation
let logOut=document.getElementById('logOut')
logOut.addEventListener('click',function(){
  let conf=confirm('Are you sure you want to Logout ?')
  if(conf==true){
    logOut.setAttribute('href','/logout')
  }
})

// Searching
const searchInput = document.getElementById("searchInput");
const clearBtn = document.getElementById("clearBtn");
clearBtn.addEventListener("click", function () {
  searchInput.value=''
  cardArr.forEach((ele)=>{ele.style.display='block'})
  document.getElementById('noResult').style.display='none'
});

function search(seachEle) {
  let toSearch = seachEle.toLowerCase();
  let divArr = [];
  let arr = document.getElementsByTagName("p");
  Array.from(arr).forEach((ele) => {
    let noteTxt = ele.innerText.toLowerCase();
    if (noteTxt.includes(toSearch)) {
      // console.log(noteTxt);
      let parent = ele.parentElement.parentElement;
      // console.log(parent);
      divArr.push(parent);
    }
  });
  cardArr.forEach((ele) => {
    ele.style.display = "none";
  });
  divArr.forEach((ele) => {
    ele.style.display = "block";
  });
  if(divArr.length==0){
    document.getElementById('noResult').style.display='block'
  }
  else{
    document.getElementById('noResult').style.display='none'
  }
}

//back to top arrow
const goAtTopArrow=document.getElementById('goAtTopArrow')
window.onscroll=function(){
  if(document.documentElement.scrollTop>150){
    goAtTopArrow.style.visibility="visible"
  }
  else{
    goAtTopArrow.style.visibility="hidden"
  }
}
